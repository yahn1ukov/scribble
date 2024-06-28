package app

import (
	"github.com/yahn1ukov/scribble/apps/notebook/internal/adapters"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/ports"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(NewConfig, NewPostgres, NewMux),
		fx.Provide(
			adapters.NewGRPCServer,
			adapters.NewHTTPHandler,
			fx.Annotate(adapters.NewPostgresRepository, fx.As(new(ports.Repository))),
			fx.Annotate(services.NewService, fx.As(new(ports.Service))),
		),
		fx.Invoke(RunPostgres, RunGRPC, RunHTTP),
	)
}
