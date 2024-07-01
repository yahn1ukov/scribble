package adapters

import (
	"context"
	"errors"

	pb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"github.com/yahn1ukov/scribble/services/notebook/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/notebook/internal/core/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	pb.UnimplementedNotebookServiceServer

	service ports.Service
}

func NewGRPCServer(service ports.Service) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) Exists(ctx context.Context, req *pb.ExistsNotebookRequest) (*emptypb.Empty, error) {
	if err := s.service.Exists(ctx, req.Id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
