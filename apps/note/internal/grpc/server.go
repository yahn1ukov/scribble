package grpc

import (
	"context"
	"errors"

	"github.com/yahn1ukov/scribble/apps/note/internal/dto"
	"github.com/yahn1ukov/scribble/apps/note/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/note/internal/services"
	pb "github.com/yahn1ukov/scribble/proto/note"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedNoteServiceServer

	service services.Service
}

func New(service services.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	input := &dto.CreateInput{
		Title:   req.Title,
		Content: req.Content,
	}

	id, err := s.service.Create(ctx, req.NotebookId, input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateNoteResponse{
		Id: id,
	}, nil
}

func (s *Server) ListNotes(ctx context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	notes, err := s.service.GetAll(ctx, req.NotebookId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListNotesResponse{
		Notes: notes,
	}, nil
}

func (s *Server) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.Note, error) {
	note, err := s.service.GetByID(ctx, req.Id, req.NotebookId)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return note, nil
}

func (s *Server) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.service.Update(ctx, req.Id, req.NotebookId, input); err != nil {
		if errors.Is(err, services.ErrNoFieldsToUpdate) {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*emptypb.Empty, error) {
	if err := s.service.Delete(ctx, req.Id, req.NotebookId); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
