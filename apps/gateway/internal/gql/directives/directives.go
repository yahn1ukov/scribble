package directives

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/middlewares"
)

type Directive struct {
	middleware *middlewares.Middleware
}

func New(middleware *middlewares.Middleware) *Directive {
	return &Directive{
		middleware: middleware,
	}
}
