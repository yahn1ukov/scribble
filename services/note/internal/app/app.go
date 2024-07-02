package app

import (
	"github.com/yahn1ukov/scribble/services/note/internal/adapters"
	"github.com/yahn1ukov/scribble/services/note/internal/core/ports"
	"github.com/yahn1ukov/scribble/services/note/internal/core/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(NewConfig, NewPostgres, NewNotebookGRPC, NewFileGRPC, NewMux),
		fx.Provide(
			adapters.NewNotebookGRPCClient,
			adapters.NewFileGRPCClient,
			adapters.NewHTTPHandler,
			fx.Annotate(adapters.NewPostgresRepository, fx.As(new(ports.Repository))),
			fx.Annotate(services.NewService, fx.As(new(ports.Service))),
		),
		fx.Invoke(RunPostgres, RunHTTP),
	)
}
