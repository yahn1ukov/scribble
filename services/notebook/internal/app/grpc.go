package app

import (
	"context"
	"fmt"
	"net"

	pb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"github.com/yahn1ukov/scribble/services/notebook/internal/adapters"
	"github.com/yahn1ukov/scribble/services/notebook/internal/config"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func RunGRPC(lc fx.Lifecycle, cfg *config.Config, grpcServer *adapters.GRPCServer) {
	listener, _ := net.Listen(
		cfg.GRPC.Network,
		fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port),
	)

	server := grpc.NewServer()

	pb.RegisterNotebookServiceServer(server, grpcServer)

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go server.Serve(listener)
			return nil
		},
		OnStop: func(_ context.Context) error {
			server.GracefulStop()
			return nil
		},
	})
}
