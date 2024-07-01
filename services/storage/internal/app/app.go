package app

import (
	"github.com/yahn1ukov/scribble/services/storage/internal/adapters"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/ports"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(NewConfig, NewMinIO),
		fx.Provide(
			adapters.NewGRPCServer,
			fx.Annotate(adapters.NewMinIORepository, fx.As(new(ports.Repository))),
			fx.Annotate(services.NewService, fx.As(new(ports.Service))),
		),
		fx.Invoke(RunMinIO, RunGRPC),
	)
}
