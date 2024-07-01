package ports

import (
	"context"

	"github.com/yahn1ukov/scribble/services/notebook/internal/core/domain"
)

type Repository interface {
	Exists(context.Context, string) error
	Create(context.Context, *domain.Notebook) error
	GetAll(context.Context) ([]*domain.Notebook, error)
	Get(context.Context, string) (*domain.Notebook, error)
	Update(context.Context, string, *domain.Notebook) error
	Delete(context.Context, string) error
}
