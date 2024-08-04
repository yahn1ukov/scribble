package grpc

import (
	"context"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/config"
	pb "github.com/yahn1ukov/scribble/proto/notebook"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *Server) {
	listener, _ := net.Listen(cfg.GRPC.Server.Network, cfg.GRPC.Server.Address)

	svr := grpc.NewServer()

	pb.RegisterNotebookServiceServer(svr, server)

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
