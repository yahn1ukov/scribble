package grpc

import (
	"context"
	"errors"
	"io"

	"github.com/yahn1ukov/scribble/apps/file/internal/dto"
	"github.com/yahn1ukov/scribble/apps/file/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/file/internal/services"
	pb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedFileServiceServer

	service services.Service
}

func NewServer(service services.Service) *server {
	return &server{
		service: service,
	}
}

func (s *server) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*emptypb.Empty, error) {
	input := &dto.UploadInput{
		Name:        req.Name,
		Size:        req.Size,
		ContentType: req.ContentType,
		Content:     req.Content,
	}

	if err := s.service.Upload(ctx, req.NoteId, input); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *server) UploadAllFiles(stream pb.FileService_UploadAllFilesServer) error {
	for {
		ctx := stream.Context()

		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		input := &dto.UploadInput{
			Name:        req.Name,
			Size:        req.Size,
			ContentType: req.ContentType,
			Content:     req.Content,
		}

		if err = s.service.Upload(ctx, req.NoteId, input); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}

func (s *server) GetAllFiles(ctx context.Context, req *pb.GetAllFilesRequest) (*pb.Files, error) {
	files, err := s.service.GetAll(ctx, req.NoteId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Files{
		Files: files,
	}, nil
}

func (s *server) DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	file, err := s.service.Get(ctx, req.Id, req.NoteId)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return file, nil
}

func (s *server) RemoveFile(ctx context.Context, req *pb.RemoveFileRequest) (*emptypb.Empty, error) {
	if err := s.service.Remove(ctx, req.Id, req.NoteId); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
