package app

import (
	"fmt"

	"github.com/yahn1ukov/scribble/apps/note/internal/config"
	notebookpb "github.com/yahn1ukov/scribble/apps/notebook/proto"
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
