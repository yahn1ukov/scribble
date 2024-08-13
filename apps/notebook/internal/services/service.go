package services

import (
	"context"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/dto"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/model"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/repositories"
	pb "github.com/yahn1ukov/scribble/proto/notebook"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Create(context.Context, string, *dto.CreateInput) error
	GetAll(context.Context, string) ([]*pb.Notebook, error)
	Update(context.Context, string, string, *dto.UpdateInput) error
	Delete(context.Context, string, string) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, userID string, input *dto.CreateInput) error {
	if input.Title == "" {
		return ErrTitleIsRequired
	}

	notebook := &model.Notebook{
		Title:       input.Title,
		Description: input.Description,
	}

	return s.repository.Create(ctx, userID, notebook)
}

func (s *service) GetAll(ctx context.Context, userID string) ([]*pb.Notebook, error) {
	notebooks, err := s.repository.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	var output []*pb.Notebook
	for _, notebook := range notebooks {
		output = append(
			output,
			&pb.Notebook{
				Id:          notebook.ID,
				Title:       notebook.Title,
				Description: notebook.Description,
				CreatedAt:   timestamppb.New(notebook.CreatedAt),
			},
		)
	}

	return output, nil
}

func (s *service) Update(ctx context.Context, id string, userID string, input *dto.UpdateInput) error {
	updatedFields := make(map[string]interface{})

	if input.Title != nil && *input.Title != "" {
		updatedFields["title"] = *input.Title
	}

	if input.Description != nil {
		updatedFields["description"] = *input.Description
	}

	if len(updatedFields) == 0 {
		return ErrNoFieldsToUpdate
	}

	return s.repository.Update(ctx, id, userID, updatedFields)
}

func (s *service) Delete(ctx context.Context, id string, userID string) error {
	return s.repository.Delete(ctx, id, userID)
}
