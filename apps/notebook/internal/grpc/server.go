package grpc

import (
	"context"
	"errors"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/dto"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/services"
	pb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedNotebookServiceServer

	service services.Service
}

func NewServer(service services.Service) *server {
	return &server{
		service: service,
	}
}

func (s *server) CreateNotebook(ctx context.Context, req *pb.CreateNotebookRequest) (*emptypb.Empty, error) {
	input := &dto.CreateInput{
		Title: req.Title,
	}

	if err := s.service.Create(ctx, input); err != nil {
		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetAllNotebooks(ctx context.Context, _ *emptypb.Empty) (*pb.Notebooks, error) {
	notebooks, err := s.service.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Notebooks{
		Notebooks: notebooks,
	}, nil
}

func (s *server) UpdateNotebook(ctx context.Context, req *pb.UpdateNotebookRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Title: req.Title,
	}

	if err := s.service.Update(ctx, req.Id, input); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteNotebook(ctx context.Context, req *pb.DeleteNotebookRequest) (*emptypb.Empty, error) {
	if err := s.service.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
