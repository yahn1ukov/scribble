package app

import (
	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	"github.com/yahn1ukov/scribble/apps/file/internal/database"
	"github.com/yahn1ukov/scribble/apps/file/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/file/internal/minio"
	"github.com/yahn1ukov/scribble/apps/file/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/file/internal/services"
	"go.uber.org/fx"
)

func New(configPath string) *fx.App {
	return fx.New(
		fx.Provide(
			func() (*config.Config, error) {
				return config.New(configPath)
			},
			database.New,
			minio.New,
		),

		fx.Provide(
			fx.Annotate(repositories.New, fx.As(new(repositories.Repository))),
			fx.Annotate(services.New, fx.As(new(services.Service))),
		),

		fx.Provide(grpc.New),

		fx.Invoke(
			minio.Run,
			grpc.Run,
		),
	)
}
