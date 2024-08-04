package grpc

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	authpb "github.com/yahn1ukov/scribble/proto/auth"
	filepb "github.com/yahn1ukov/scribble/proto/file"
	notepb "github.com/yahn1ukov/scribble/proto/note"
	notebookpb "github.com/yahn1ukov/scribble/proto/notebook"
	userpb "github.com/yahn1ukov/scribble/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUser(cfg *config.Config) (userpb.UserServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.User.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(connection)

	return client, nil
}

func NewNotebook(cfg *config.Config) (notebookpb.NotebookServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.Notebook.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := notebookpb.NewNotebookServiceClient(connection)

	return client, nil
}

func NewNote(cfg *config.Config) (notepb.NoteServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.Note.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := notepb.NewNoteServiceClient(connection)

	return client, nil
}

func NewFile(cfg *config.Config) (filepb.FileServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.File.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := filepb.NewFileServiceClient(connection)

	return client, nil
}

func NewAuth(cfg *config.Config) (authpb.AuthServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.Auth.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := authpb.NewAuthServiceClient(connection)

	return client, nil
}
