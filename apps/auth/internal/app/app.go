package app

import (
	"github.com/yahn1ukov/scribble/apps/auth/internal/config"
	"github.com/yahn1ukov/scribble/apps/auth/internal/grpc"
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
			grpc.New,
		),

		fx.Invoke(grpc.Run),
	)
}
