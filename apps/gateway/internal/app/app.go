package app

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/directives"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/resolvers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			grpc.NewUser,
			grpc.NewNotebook,
			grpc.NewNote,
			grpc.NewFile,
			grpc.NewAuth,
			directives.NewDirective,
			resolvers.NewResolver,
			gql.New,
			http.NewMiddleware,
			http.NewHandler,
		),
		fx.Invoke(http.Run),
	)
}
