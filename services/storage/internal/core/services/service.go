package services

import (
	"context"
	"fmt"

	"github.com/yahn1ukov/scribble/services/storage/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/ports"
)

type service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) ports.Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Upload(ctx context.Context, in *dto.UploadInput) error {
	file := &domain.File{
		URL:         fmt.Sprintf("%s/%s", in.NoteID, in.Name),
		Size:        in.Size,
		ContentType: in.ContentType,
		Content:     in.Content,
	}

	return s.repository.Upload(ctx, file)
}
