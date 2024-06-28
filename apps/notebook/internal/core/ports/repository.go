package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/domain"
)

type Repository interface {
	Exists(context.Context, uuid.UUID) error
	Create(context.Context, *domain.Notebook) error
	GetAll(context.Context) ([]*domain.Notebook, error)
	Get(context.Context, uuid.UUID) (*domain.Notebook, error)
	Update(context.Context, uuid.UUID, *domain.Notebook) error
	Delete(context.Context, uuid.UUID) error
}
