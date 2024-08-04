package directives

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http"
)

type Directive struct {
	middleware *http.Middleware
}

func NewDirective(middleware *http.Middleware) *Directive {
	return &Directive{
		middleware: middleware,
	}
}

func (c *Directive) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	userID := c.middleware.GetUserIDFromCtx(ctx)
	if userID == "" {
		return nil, fmt.Errorf("access denied")
	}

	return next(ctx)
}
