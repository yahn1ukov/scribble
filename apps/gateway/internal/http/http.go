package http

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"go.uber.org/fx"
	"net/http"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *handler.Server, handler *Handler, middleware *Middleware) {
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middleware.AuthMiddleware(server))

	mux.Handle("GET /notes/{noteId}/files/{fileId}", middleware.AuthMiddleware(http.HandlerFunc(handler.DownloadFile)))

	svr := &http.Server{
		Addr:    cfg.HTTP.Address,
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
