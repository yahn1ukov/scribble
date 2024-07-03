package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/file/internal/core/domain"
)

type Repository interface {
	Create(context.Context, *domain.File) error
	GetAll(context.Context, string) ([]*domain.File, error)
	Get(context.Context, string) (*domain.File, error)
	Delete(context.Context, string) error
}
