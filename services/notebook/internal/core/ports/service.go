package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/notebook/internal/core/dto"
)

type Service interface {
	Exists(context.Context, string) error
	Create(context.Context, *dto.CreateInput) error
	GetAll(context.Context) ([]*dto.GetOutput, error)
	Update(context.Context, string, *dto.UpdateInput) error
	Delete(context.Context, string) error
}
