package app

import (
	"github.com/yahn1ukov/scribble/apps/auth/internal/config"
	"github.com/yahn1ukov/scribble/apps/auth/internal/grpc"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			grpc.NewUser,
			grpc.NewServer,
		),
		fx.Invoke(grpc.Run),
	)
}
