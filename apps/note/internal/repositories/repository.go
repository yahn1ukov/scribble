package repositories

import (
	"context"

	"github.com/yahn1ukov/scribble/apps/note/internal/model"
)

type Repository interface {
	Create(context.Context, string, *model.Note) (string, error)
	GetAll(context.Context, string) ([]*model.Note, error)
	Get(context.Context, string, string) (*model.Note, error)
	Update(context.Context, string, string, *model.Note) error
	Delete(context.Context, string, string) error
}
