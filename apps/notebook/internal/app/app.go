package app

import (
	"github.com/yahn1ukov/scribble/apps/notebook/internal/config"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/database"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/services"
	"go.uber.org/fx"
)

func New(configPath string) *fx.App {
	return fx.New(
		fx.Provide(
			func() (*config.Config, error) {
				return config.New(configPath)
			},
			database.New,
		),

		fx.Provide(
			fx.Annotate(repositories.New, fx.As(new(repositories.Repository))),
			fx.Annotate(services.New, fx.As(new(services.Service))),
		),

		fx.Provide(grpc.New),

		fx.Invoke(grpc.Run),
	)
}
