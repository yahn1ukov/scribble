package grpc

import (
	"context"
	"net"

	"github.com/yahn1ukov/scribble/apps/note/internal/config"
	pb "github.com/yahn1ukov/scribble/libs/grpc/note"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Params struct {
	fx.In

	Lc     fx.Lifecycle
	Cfg    *config.Config
	Server *server
}

func Run(p Params) {
	listener, _ := net.Listen(p.Cfg.GRPC.Server.Network, p.Cfg.GRPC.Server.Address)

	server := grpc.NewServer()

	pb.RegisterNoteServiceServer(server, p.Server)

	p.Lc.Append(fx.Hook{
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
