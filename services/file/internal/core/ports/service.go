package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/file/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/file/internal/core/dto"
)

type Service interface {
	Create(context.Context, *dto.CreateInput) error
	GetAll(context.Context, string) ([]*domain.File, error)
}
