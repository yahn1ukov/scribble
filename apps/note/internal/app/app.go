package app

import (
	"github.com/yahn1ukov/scribble/apps/note/internal/config"
	"github.com/yahn1ukov/scribble/apps/note/internal/database"
	"github.com/yahn1ukov/scribble/apps/note/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/note/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/note/internal/services"
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
