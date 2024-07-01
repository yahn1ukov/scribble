package services

import (
	"context"

	"github.com/yahn1ukov/scribble/services/notebook/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/notebook/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/notebook/internal/core/ports"
)

type service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) ports.Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Exists(ctx context.Context, id string) error {
	return s.repository.Exists(ctx, id)
}

func (s *service) Create(ctx context.Context, in *dto.CreateInput) error {
	if in.Title == "" {
		return domain.ErrTitleRequired
	}

	notebook := &domain.Notebook{
		Title: in.Title,
	}

	return s.repository.Create(ctx, notebook)
}

func (s *service) GetAll(ctx context.Context) ([]*dto.GetOutput, error) {
	notebooks, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*dto.GetOutput
	for _, notebook := range notebooks {
		out = append(
			out,
			&dto.GetOutput{
				ID:        notebook.ID,
				Title:     notebook.Title,
				CreatedAt: notebook.CreatedAt,
			},
		)
	}

	return out, nil
}

func (s *service) Update(ctx context.Context, id string, in *dto.UpdateInput) error {
	notebook, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	if in.Title == "" {
		return domain.ErrTitleRequired
	}

	notebook.Title = in.Title

	return s.repository.Update(ctx, id, notebook)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
