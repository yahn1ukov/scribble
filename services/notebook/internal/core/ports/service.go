package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/services/notebook/internal/core/dto"
)

type Service interface {
	Exists(context.Context, uuid.UUID) error
	Create(context.Context, *dto.CreateInput) error
	GetAll(context.Context) ([]*dto.GetOutput, error)
	Update(context.Context, uuid.UUID, *dto.UpdateInput) error
	Delete(context.Context, uuid.UUID) error
}
