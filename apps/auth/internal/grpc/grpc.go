package grpc

import (
	"context"
	"net"

	"github.com/yahn1ukov/scribble/apps/auth/internal/config"
	pb "github.com/yahn1ukov/scribble/proto/auth"
	userpb "github.com/yahn1ukov/scribble/proto/user"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUser(cfg *config.Config) (userpb.UserServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.User.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(connection)

	return client, nil
}

func Run(lc fx.Lifecycle, cfg *config.Config, server *Server) {
	listener, _ := net.Listen(cfg.GRPC.Server.Network, cfg.GRPC.Server.Address)

	svr := grpc.NewServer()

	pb.RegisterAuthServiceServer(svr, server)

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
