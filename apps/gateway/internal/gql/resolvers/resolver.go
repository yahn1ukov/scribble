package resolvers

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/http"
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
	middleware     *http.Middleware
	userClient     userpb.UserServiceClient
	notebookClient notebookpb.NotebookServiceClient
	noteClient     notepb.NoteServiceClient
	fileClient     filepb.FileServiceClient
	authClient     authpb.AuthServiceClient
}

func NewResolver(
	cfg *config.Config,
	middleware *http.Middleware,
	userClient userpb.UserServiceClient,
	notebookClient notebookpb.NotebookServiceClient,
	noteClient notepb.NoteServiceClient,
	fileClient filepb.FileServiceClient,
	authClient authpb.AuthServiceClient,
) *Resolver {
	return &Resolver{
		cfg:            cfg,
		middleware:     middleware,
		userClient:     userClient,
		notebookClient: notebookClient,
		noteClient:     noteClient,
		fileClient:     fileClient,
		authClient:     authClient,
	}
}
