package app

import (
	"github.com/yahn1ukov/scribble/apps/notebook/internal/config"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/database"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/services"
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
		fx.Invoke(database.Run, grpc.Run),
	)
}
