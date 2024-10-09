package resolvers

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/mappers"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/middlewares"
	authpb "github.com/yahn1ukov/scribble/proto/auth"
	filepb "github.com/yahn1ukov/scribble/proto/file"
	notepb "github.com/yahn1ukov/scribble/proto/note"
	notebookpb "github.com/yahn1ukov/scribble/proto/notebook"
	userpb "github.com/yahn1ukov/scribble/proto/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cfg            *config.Config
	middleware     *middlewares.Middleware
	mapper         *mappers.Mapper
	authClient     authpb.AuthServiceClient
	fileClient     filepb.FileServiceClient
	noteClient     notepb.NoteServiceClient
	notebookClient notebookpb.NotebookServiceClient
	userClient     userpb.UserServiceClient
}

func New(
	cfg *config.Config,
	middleware *middlewares.Middleware,
	mapper *mappers.Mapper,
	authClient authpb.AuthServiceClient,
	fileClient filepb.FileServiceClient,
	noteClient notepb.NoteServiceClient,
	notebookClient notebookpb.NotebookServiceClient,
	userClient userpb.UserServiceClient,
) *Resolver {
	return &Resolver{
		cfg:            cfg,
		middleware:     middleware,
		mapper:         mapper,
		authClient:     authClient,
		fileClient:     fileClient,
		noteClient:     noteClient,
		notebookClient: notebookClient,
		userClient:     userClient,
	}
}
