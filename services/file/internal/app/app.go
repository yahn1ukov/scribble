package app

import (
	"github.com/yahn1ukov/scribble/services/file/internal/adapters"
	"github.com/yahn1ukov/scribble/services/file/internal/core/ports"
	"github.com/yahn1ukov/scribble/services/file/internal/core/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(NewConfig, NewPostgres),
		fx.Provide(
			adapters.NewGRPCServer,
			fx.Annotate(adapters.NewPostgresRepository, fx.As(new(ports.Repository))),
			fx.Annotate(services.NewService, fx.As(new(ports.Service))),
		),
		fx.Invoke(RunPostgres, RunGRPC),
	)
}
