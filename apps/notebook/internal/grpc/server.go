package grpc

import (
	"context"
	"errors"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/dto"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/services"
	pb "github.com/yahn1ukov/scribble/proto/notebook"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedNotebookServiceServer

	service services.Service
}

func NewServer(service services.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) CreateNotebook(ctx context.Context, req *pb.CreateNotebookRequest) (*emptypb.Empty, error) {
	input := &dto.CreateInput{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := s.service.Create(ctx, req.UserId, input); err != nil {
		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListNotebooks(ctx context.Context, req *pb.ListNotebooksRequest) (*pb.ListNotebooksResponse, error) {
	notebooks, err := s.service.GetAll(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListNotebooksResponse{
		Notebooks: notebooks,
	}, nil
}

func (s *Server) UpdateNotebook(ctx context.Context, req *pb.UpdateNotebookRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := s.service.Update(ctx, req.Id, req.UserId, input); err != nil {
		if errors.Is(err, services.ErrNoFieldsToUpdate) {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteNotebook(ctx context.Context, req *pb.DeleteNotebookRequest) (*emptypb.Empty, error) {
	if err := s.service.Delete(ctx, req.Id, req.UserId); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
