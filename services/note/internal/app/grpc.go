package app

import (
	"fmt"

	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
	notebookpb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	storagepb "github.com/yahn1ukov/scribble/libs/grpc/storage"
	"github.com/yahn1ukov/scribble/services/note/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewNotebookGRPC(cfg *config.Config) (notebookpb.NotebookServiceClient, error) {
	connection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.GRPC.Notebook.Host, cfg.GRPC.Notebook.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := notebookpb.NewNotebookServiceClient(connection)

	return client, nil
}

func NewStorageGRPC(cfg *config.Config) (storagepb.StorageServiceClient, error) {
	connection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.GRPC.Storage.Host, cfg.GRPC.Storage.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := storagepb.NewStorageServiceClient(connection)

	return client, nil
}

func NewFileGRPC(cfg *config.Config) (filepb.FileServiceClient, error) {
	connection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.GRPC.File.Host, cfg.GRPC.File.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := filepb.NewFileServiceClient(connection)

	return client, nil
}
