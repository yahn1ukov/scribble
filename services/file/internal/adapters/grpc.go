package adapters

import (
	"context"

	pb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"github.com/yahn1ukov/scribble/services/file/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/file/internal/core/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
	pb.UnimplementedFileServiceServer

	service ports.Service
}

func NewGRPCServer(service ports.Service) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) Create(ctx context.Context, req *pb.CreateFileRequest) (*emptypb.Empty, error) {
	for _, file := range req.Files {
		if err := s.service.Create(ctx,
			&dto.CreateInput{
				Name:        file.Name,
				Size:        file.Size,
				ContentType: file.ContentType,
				NoteID:      req.NoteId,
			},
		); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) GetAll(ctx context.Context, req *pb.GetAllFileRequest) (*pb.Files, error) {
	files, err := s.service.GetAll(ctx, req.NoteId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var filesPb []*pb.File
	for _, file := range files {
		filesPb = append(
			filesPb,
			&pb.File{
				Id:          file.ID,
				Name:        file.Name,
				Size:        file.Size,
				ContentType: file.ContentType,
				Url:         file.URL,
				CreatedAt:   timestamppb.New(file.CreatedAt),
			},
		)
	}

	return &pb.Files{
		Files: filesPb,
	}, nil
}
