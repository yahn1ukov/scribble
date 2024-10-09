package app

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/directives"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/mappers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/resolvers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/handlers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/middlewares"
	"go.uber.org/fx"
)

func New(configPath string) *fx.App {
	return fx.New(
		fx.Provide(
			func() (*config.Config, error) {
				return config.New(configPath)
			},
		),

		fx.Provide(
			grpc.NewUser,
			grpc.NewNotebook,
			grpc.NewNote,
			grpc.NewFile,
			grpc.NewAuth,
		),

		fx.Provide(
			directives.New,
			mappers.New,
			resolvers.New,
			gql.New,
		),

		fx.Provide(
			middlewares.New,
			handlers.New,
		),

		fx.Invoke(http.Run),
	)
}
