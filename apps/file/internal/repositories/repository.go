package repositories

import (
	"context"

	"github.com/yahn1ukov/scribble/apps/file/internal/model"
)

type Repository interface {
	Create(context.Context, string, *model.File) error
	GetAll(context.Context, string) ([]*model.File, error)
	Get(context.Context, string, string) (*model.File, error)
	Delete(context.Context, string, string) error
}
