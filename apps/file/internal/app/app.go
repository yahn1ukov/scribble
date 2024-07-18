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

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			database.New,
			minio.New,
			fx.Annotate(repositories.NewPostgresRepository, fx.As(new(repositories.Repository))),
			fx.Annotate(services.NewService, fx.As(new(services.Service))),
			grpc.NewServer,
		),
		fx.Invoke(minio.Run, grpc.Run),
	)
}
