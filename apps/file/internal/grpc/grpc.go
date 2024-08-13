package grpc

import (
	"context"
	"net"

	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	pb "github.com/yahn1ukov/scribble/proto/file"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *Server) {
	listener, _ := net.Listen(cfg.GRPC.Server.Network, cfg.GRPC.Server.Address)

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
