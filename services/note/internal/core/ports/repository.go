package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/services/note/internal/core/domain"
)

type Repository interface {
	Create(context.Context, uuid.UUID, *domain.Note) error
	GetAll(context.Context, uuid.UUID) ([]*domain.Note, error)
	Get(context.Context, uuid.UUID, uuid.UUID) (*domain.Note, error)
	Update(context.Context, uuid.UUID, uuid.UUID, *domain.Note) error
	Delete(context.Context, uuid.UUID, uuid.UUID) error
}
