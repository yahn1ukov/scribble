package app

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/resolvers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc/clients"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			grpc.NewNotebook,
			grpc.NewNote,
			grpc.NewFile,
			clients.NewClient,
			resolvers.NewResolver,
			gql.New,
			http.NewHandler,
		),
		fx.Invoke(http.Run),
	)
}
