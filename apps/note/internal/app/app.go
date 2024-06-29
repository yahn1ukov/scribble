package app

import (
	"github.com/yahn1ukov/scribble/apps/note/internal/adapters"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/ports"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(NewConfig, NewPostgres, NewNotebookGRPC, NewMux),
		fx.Provide(
			adapters.NewNotebookGRPCClient,
			adapters.NewHTTPHandler,
			fx.Annotate(adapters.NewPostgresRepository, fx.As(new(ports.Repository))),
			fx.Annotate(services.NewService, fx.As(new(ports.Service))),
		),
		fx.Invoke(RunPostgres, RunHTTP),
	)
}
