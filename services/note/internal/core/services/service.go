package services

import (
	"context"

	"github.com/yahn1ukov/scribble/services/note/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/note/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/note/internal/core/ports"
)

type service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) ports.Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, notebookId string, in *dto.CreateInput) (string, error) {
	if in.Title == "" {
		return "", domain.ErrTitleRequired
	}

	note := &domain.Note{
		Title: in.Title,
		Body:  in.Body,
	}

	return s.repository.Create(ctx, notebookId, note)
}

func (s *service) GetAll(ctx context.Context, notebookId string) ([]*dto.GetOutput, error) {
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

func (s *service) Get(ctx context.Context, id string, notebookId string) (*dto.GetOutput, error) {
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

func (s *service) Update(ctx context.Context, id string, notebookId string, in *dto.UpdateInput) error {
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

func (s *service) Delete(ctx context.Context, id string, notebookId string) error {
	return s.repository.Delete(ctx, id, notebookId)
}
