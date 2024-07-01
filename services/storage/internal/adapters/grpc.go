package adapters

import (
	"bytes"
	"io"

	pb "github.com/yahn1ukov/scribble/libs/grpc/storage"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	pb.UnimplementedStorageServiceServer

	service ports.Service
}

func NewGRPCServer(service ports.Service) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) Upload(stream pb.StorageService_UploadServer) error {
	for {
		ctx := stream.Context()

		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		content := bytes.NewReader(req.Content)

		in := &dto.UploadInput{
			Name:        req.Name,
			Size:        req.Size,
			ContentType: req.ContentType,
			NoteID:      req.NoteId,
			Content:     content,
		}

		if err = s.service.Upload(ctx, in); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}
