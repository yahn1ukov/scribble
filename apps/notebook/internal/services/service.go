package services

import (
	"context"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/dto"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/model"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/repositories"
	pb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Create(context.Context, *dto.CreateInput) error
	GetAll(context.Context) ([]*pb.Notebook, error)
	Update(context.Context, string, *dto.UpdateInput) error
	Delete(context.Context, string) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, input *dto.CreateInput) error {
	if input.Title == "" {
		return repositories.ErrTitleRequired
	}

	notebook := &model.Notebook{
		Title: input.Title,
	}

	return s.repository.Create(ctx, notebook)
}

func (s *service) GetAll(ctx context.Context) ([]*pb.Notebook, error) {
	notebooks, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*pb.Notebook
	for _, notebook := range notebooks {
		output = append(output,
			&pb.Notebook{
				Id:        notebook.ID,
				Title:     notebook.Title,
				CreatedAt: timestamppb.New(notebook.CreatedAt),
			},
		)
	}

	return output, nil
}

func (s *service) Update(ctx context.Context, id string, input *dto.UpdateInput) error {
	notebook, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	if input.Title == "" {
		return repositories.ErrTitleRequired
	}

	notebook.Title = input.Title

	return s.repository.Update(ctx, id, notebook)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
