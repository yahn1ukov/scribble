package repositories

import (
	"context"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/model"
)

type Repository interface {
	Create(context.Context, *model.Notebook) error
	GetAll(context.Context) ([]*model.Notebook, error)
	Get(context.Context, string) (*model.Notebook, error)
	Update(context.Context, string, *model.Notebook) error
	Delete(context.Context, string) error
}
