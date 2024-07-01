package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/storage/internal/core/dto"
)

type Service interface {
	Upload(context.Context, *dto.UploadInput) error
}
