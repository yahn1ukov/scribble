package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/domain"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/dto"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/ports"
)

type service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) ports.Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, notebookId uuid.UUID, in *dto.CreateInput) error {
	if in.Title == "" {
		return domain.ErrTitleRequired
	}

	note := &domain.Note{
		Title: in.Title,
		Body:  in.Body,
	}

	return s.repository.Create(ctx, notebookId, note)
}

func (s *service) GetAll(ctx context.Context, notebookId uuid.UUID) ([]*dto.GetOutput, error) {
	notes, err := s.repository.GetAll(ctx, notebookId)
	if err != nil {
		return nil, err
	}

	var out []*dto.GetOutput
	for _, note := range notes {
		out = append(
			out,
			&dto.GetOutput{
				ID:        note.ID,
				Title:     note.Title,
				Body:      note.Body,
				CreatedAt: note.CreatedAt,
			},
		)
	}

	return out, nil
}

func (s *service) Get(ctx context.Context, id uuid.UUID, notebookId uuid.UUID) (*dto.GetOutput, error) {
	note, err := s.repository.Get(ctx, id, notebookId)
	if err != nil {
		return nil, err
	}

	out := &dto.GetOutput{
		ID:        note.ID,
		Title:     note.Title,
		Body:      note.Body,
		CreatedAt: note.CreatedAt,
	}

	return out, nil
}

func (s *service) Update(ctx context.Context, id uuid.UUID, notebookId uuid.UUID, in *dto.UpdateInput) error {
	note, err := s.repository.Get(ctx, id, notebookId)
	if err != nil {
		return err
	}

	if in.Title == "" {
		return domain.ErrTitleRequired
	}

	note.Title = in.Title
	note.Body = in.Body

	return s.repository.Update(ctx, id, notebookId, note)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID, notebookId uuid.UUID) error {
	return s.repository.Delete(ctx, id, notebookId)
}
