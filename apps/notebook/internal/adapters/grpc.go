package adapters

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/domain"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/ports"
	pb "github.com/yahn1ukov/scribble/apps/notebook/proto"
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
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = s.service.Exists(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
