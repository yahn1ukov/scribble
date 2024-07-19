package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	pb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *server) {
	listener, _ := net.Listen(
		cfg.GRPC.Server.Network,
		fmt.Sprintf("%s:%d", cfg.GRPC.Server.Host, cfg.GRPC.Server.Port),
	)

	svr := grpc.NewServer()

	pb.RegisterFileServiceServer(svr, server)

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go svr.Serve(listener)
			return nil
		},
		OnStop: func(_ context.Context) error {
			svr.GracefulStop()
			return nil
		},
	})
}
