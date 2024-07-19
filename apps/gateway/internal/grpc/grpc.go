package grpc

import (
	"fmt"

	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
	notepb "github.com/yahn1ukov/scribble/libs/grpc/note"
	notebookpb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewNotebook(cfg *config.Config) (notebookpb.NotebookServiceClient, error) {
	connection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.GRPC.Client.Notebook.Host, cfg.GRPC.Client.Notebook.Port),
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
		fmt.Sprintf("%s:%d", cfg.GRPC.Client.Note.Host, cfg.GRPC.Client.Note.Port),
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
		fmt.Sprintf("%s:%d", cfg.GRPC.Client.File.Host, cfg.GRPC.Client.File.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := filepb.NewFileServiceClient(connection)

	return client, nil
}
