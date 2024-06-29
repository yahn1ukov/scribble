package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/services/note/internal/core/dto"
)

type Service interface {
	Create(context.Context, uuid.UUID, *dto.CreateInput) error
	GetAll(context.Context, uuid.UUID) ([]*dto.GetOutput, error)
	Get(context.Context, uuid.UUID, uuid.UUID) (*dto.GetOutput, error)
	Update(context.Context, uuid.UUID, uuid.UUID, *dto.UpdateInput) error
	Delete(context.Context, uuid.UUID, uuid.UUID) error
}
