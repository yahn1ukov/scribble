package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/note/internal/core/domain"
)

type Repository interface {
	Create(context.Context, string, *domain.Note) (string, error)
	GetAll(context.Context, string) ([]*domain.Note, error)
	Get(context.Context, string, string) (*domain.Note, error)
	Update(context.Context, string, string, *domain.Note) error
	Delete(context.Context, string, string) error
}
