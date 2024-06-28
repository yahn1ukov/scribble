package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/config"
	"go.uber.org/fx"
)

func RunHTTP(lc fx.Lifecycle, cfg *config.Config, mux *http.ServeMux) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
