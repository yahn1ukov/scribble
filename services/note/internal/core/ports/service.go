package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/note/internal/core/dto"
)

type Service interface {
	Create(context.Context, string, *dto.CreateInput) (string, error)
	GetAll(context.Context, string) ([]*dto.GetOutput, error)
	Get(context.Context, string, string) (*dto.GetOutput, error)
	Update(context.Context, string, string, *dto.UpdateInput) error
	Delete(context.Context, string, string) error
}
