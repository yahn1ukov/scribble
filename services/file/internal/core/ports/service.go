package ports

import (
	"context"

	pb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"github.com/yahn1ukov/scribble/services/file/internal/core/dto"
)

type Service interface {
	Upload(context.Context, *dto.UploadInput) error
	GetAll(context.Context, string) ([]*pb.File, error)
	Get(context.Context, string) (*dto.GetOutput, error)
	Delete(context.Context, string) error
}
