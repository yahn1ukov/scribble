package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *handler.Server, handler *Handler) {
	mux := http.NewServeMux()

	mux.Handle("/", server)
	mux.HandleFunc("GET /files/{fileId}/notes/{noteId}", handler.DownloadFile)

	svr := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go svr.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return svr.Shutdown(ctx)
		},
	})
}
