package grpc

import (
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
	notepb "github.com/yahn1ukov/scribble/libs/grpc/note"
	notebookpb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Params struct {
	fx.In

	Cfg *config.Config
}

func NewNotebook(p Params) (notebookpb.NotebookServiceClient, error) {
	connection, err := grpc.NewClient(
		p.Cfg.GRPC.Client.Notebook.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := notebookpb.NewNotebookServiceClient(connection)

	return client, nil
}

func NewNote(p Params) (notepb.NoteServiceClient, error) {
	connection, err := grpc.NewClient(
		p.Cfg.GRPC.Client.Note.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := notepb.NewNoteServiceClient(connection)

	return client, nil
}

func NewFile(p Params) (filepb.FileServiceClient, error) {
	connection, err := grpc.NewClient(
		p.Cfg.GRPC.Client.File.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := filepb.NewFileServiceClient(connection)

	return client, nil
}
