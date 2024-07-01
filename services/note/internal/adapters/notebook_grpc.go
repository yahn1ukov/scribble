package adapters

import (
	"context"

	"github.com/yahn1ukov/scribble/libs/grpc"
	notebookpb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
)

type NotebookGRPCClient struct {
	client notebookpb.NotebookServiceClient
}

func NewNotebookGRPCClient(client notebookpb.NotebookServiceClient) *NotebookGRPCClient {
	return &NotebookGRPCClient{
		client: client,
	}
}

func (c *NotebookGRPCClient) Exists(ctx context.Context, id string) *grpc.Error {
	if _, err := c.client.Exists(ctx,
		&notebookpb.ExistsNotebookRequest{
			Id: id,
		},
	); err != nil {
		return grpc.ParseError(err)
	}

	return nil
}
