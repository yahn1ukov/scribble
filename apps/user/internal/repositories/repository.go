package repositories

import (
	"context"
	"github.com/yahn1ukov/scribble/apps/user/internal/model"
)

type Repository interface {
	Create(context.Context, *model.User) (string, error)
	FindByEmail(context.Context, string) (*model.User, error)
	GetByID(context.Context, string) (*model.User, error)
	Update(context.Context, string, map[string]interface{}) error
	UpdatePassword(context.Context, string, *model.User) error
	Delete(context.Context, string) error
}
