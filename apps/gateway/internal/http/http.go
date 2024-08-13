package http

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/handlers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/middlewares"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *handler.Server, handler *handlers.Handler, middleware *middlewares.Middleware) {
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middleware.Auth(server))

	mux.Handle("GET /notes/{noteId}/files/{fileId}", middleware.Auth(http.HandlerFunc(handler.DownloadFile)))

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
