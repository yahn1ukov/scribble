package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/graph"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/resolvers"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Resolver *resolvers.Resolver
}

func New(p Params) *handler.Server {
	cfg := graph.Config{
		Resolvers: p.Resolver,
	}

	schema := graph.NewExecutableSchema(cfg)

	server := handler.New(schema)

	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.MultipartForm{})

	server.Use(extension.Introspection{})

	return server
}
