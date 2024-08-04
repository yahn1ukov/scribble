package repositories

import (
	"context"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/model"
)

type Repository interface {
	Create(context.Context, string, *model.Notebook) error
	GetAll(context.Context, string) ([]*model.Notebook, error)
	Update(context.Context, string, string, map[string]interface{}) error
	Delete(context.Context, string, string) error
}
