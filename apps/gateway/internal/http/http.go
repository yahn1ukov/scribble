package http

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc     fx.Lifecycle
	Cfg    *config.Config
	Server *handler.Server
}

func Run(p Params) {
	mux := http.NewServeMux()

	mux.Handle("/", p.Server)

	server := &http.Server{
		Addr:    p.Cfg.HTTP.Address,
		Handler: mux,
	}

	p.Lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
