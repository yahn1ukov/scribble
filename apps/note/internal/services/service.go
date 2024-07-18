package services

import (
	"context"

	"github.com/yahn1ukov/scribble/apps/note/internal/dto"
	"github.com/yahn1ukov/scribble/apps/note/internal/model"
	"github.com/yahn1ukov/scribble/apps/note/internal/repositories"
	pb "github.com/yahn1ukov/scribble/libs/grpc/note"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Create(context.Context, string, *dto.CreateInput) (string, error)
	GetAll(context.Context, string) ([]*pb.Note, error)
	Get(context.Context, string, string) (*pb.Note, error)
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

func (s *service) Create(ctx context.Context, notebookID string, input *dto.CreateInput) (string, error) {
	if input.Title == "" {
		return "", repositories.ErrTitleRequired
	}

	return s.repository.Create(ctx, notebookID,
		&model.Note{
			Title: input.Title,
			Body:  input.Body,
		},
	)
}

func (s *service) GetAll(ctx context.Context, notebookID string) ([]*pb.Note, error) {
	notes, err := s.repository.GetAll(ctx, notebookID)
	if err != nil {
		return nil, err
	}

	var output []*pb.Note
	for _, note := range notes {
		output = append(output,
			&pb.Note{
				Id:        note.ID,
				Title:     note.Title,
				Body:      note.Body,
				CreatedAt: timestamppb.New(note.CreatedAt),
			},
		)
	}

	return output, nil
}

func (s *service) Get(ctx context.Context, id string, notebookID string) (*pb.Note, error) {
	note, err := s.repository.Get(ctx, id, notebookID)
	if err != nil {
		return nil, err
	}

	return &pb.Note{
		Id:        note.ID,
		Title:     note.Title,
		Body:      note.Body,
		CreatedAt: timestamppb.New(note.CreatedAt),
	}, nil
}

func (s *service) Update(ctx context.Context, id string, notebookID string, input *dto.UpdateInput) error {
	note, err := s.repository.Get(ctx, id, notebookID)
	if err != nil {
		return err
	}

	if input.Title == "" {
		return repositories.ErrTitleRequired
	}

	note.Title = input.Title
	note.Body = input.Body

	return s.repository.Update(ctx, id, notebookID, note)
}

func (s *service) Delete(ctx context.Context, id string, notebookID string) error {
	return s.repository.Delete(ctx, id, notebookID)
}
