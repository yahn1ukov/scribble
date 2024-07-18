package resolvers

import "github.com/yahn1ukov/scribble/apps/gateway/internal/grpc/clients"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	grpc *clients.Client
}

func NewResolver(grpc *clients.Client) *Resolver {
	return &Resolver{
		grpc: grpc,
	}
}
