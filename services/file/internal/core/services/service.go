package services

import (
	"context"
	"fmt"

	"github.com/yahn1ukov/scribble/services/file/internal/config"
	"github.com/yahn1ukov/scribble/services/file/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/file/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/file/internal/core/ports"
)

type service struct {
	cfg        *config.Config
	repository ports.Repository
}

func NewService(cfg *config.Config, repository ports.Repository) ports.Service {
	return &service{
		cfg:        cfg,
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, in *dto.CreateInput) error {
	file := &domain.File{
		Name:        in.Name,
		Size:        in.Size,
		ContentType: in.ContentType,
		URL:         fmt.Sprintf("%s/%s/%s", s.cfg.Storage.MinIO.Bucket, in.NoteID, in.Name),
		NoteID:      in.NoteID,
	}

	return s.repository.Create(ctx, file)
}

func (s *service) GetAll(ctx context.Context, noteId string) ([]*domain.File, error) {
	return s.repository.GetAll(ctx, noteId)
}
