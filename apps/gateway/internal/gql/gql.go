package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/directives"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/graph"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/resolvers"
)

func New(resolver *resolvers.Resolver, directive *directives.Directive) *handler.Server {
	cfg := graph.Config{
		Resolvers: resolver,
		Directives: graph.DirectiveRoot{
			IsAuthenticated: directive.IsAuthenticated,
		},
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
