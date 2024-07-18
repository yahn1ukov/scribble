package app

import (
	"github.com/yahn1ukov/scribble/apps/note/internal/config"
	"github.com/yahn1ukov/scribble/apps/note/internal/database"
	"github.com/yahn1ukov/scribble/apps/note/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/note/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/note/internal/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			database.New,
			fx.Annotate(repositories.NewPostgresRepository, fx.As(new(repositories.Repository))),
			fx.Annotate(services.NewService, fx.As(new(services.Service))),
			grpc.NewServer,
		),
		fx.Invoke(grpc.Run),
	)
}
