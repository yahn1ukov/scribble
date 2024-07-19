package grpc

import (
	"context"
	"errors"

	"github.com/yahn1ukov/scribble/apps/note/internal/dto"
	"github.com/yahn1ukov/scribble/apps/note/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/note/internal/services"
	pb "github.com/yahn1ukov/scribble/libs/grpc/note"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedNoteServiceServer

	service services.Service
}

func NewServer(service services.Service) *server {
	return &server{
		service: service,
	}
}

func (s *server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	input := &dto.CreateInput{
		Title: req.Title,
		Body:  req.Body,
	}

	id, err := s.service.Create(ctx, req.NotebookId, input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateNoteResponse{
		Id: id,
	}, nil
}

func (h *server) GetAllNotes(ctx context.Context, req *pb.GetAllNotesRequest) (*pb.Notes, error) {
	notes, err := h.service.GetAll(ctx, req.NotebookId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Notes{
		Notes: notes,
	}, nil
}

func (h *server) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.Note, error) {
	note, err := h.service.Get(ctx, req.Id, req.NotebookId)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return note, nil
}

func (h *server) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Title: req.Title,
		Body:  req.Body,
	}

	if err := h.service.Update(ctx, req.Id, req.NotebookId, input); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (h *server) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*emptypb.Empty, error) {
	if err := h.service.Delete(ctx, req.Id, req.NotebookId); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
