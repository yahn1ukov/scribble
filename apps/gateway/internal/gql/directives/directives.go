package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/middlewares"
)

type Directive struct {
	middleware *middlewares.Middleware
}

func NewDirective(middleware *middlewares.Middleware) *Directive {
	return &Directive{
		middleware: middleware,
	}
}

func (c *Directive) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	userID := c.middleware.GetUserIDFromCtx(ctx)
	if userID == "" {
		return nil, middlewares.ErrUnauthorized
	}

	return next(ctx)
}
