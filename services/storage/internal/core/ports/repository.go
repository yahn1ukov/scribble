package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/storage/internal/core/domain"
)

type Repository interface {
	Upload(context.Context, *domain.File) error
}
