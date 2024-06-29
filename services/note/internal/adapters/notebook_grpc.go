package adapters

import (
	"context"

	"github.com/google/uuid"
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

func (c *NotebookGRPCClient) Exists(ctx context.Context, id uuid.UUID) *grpc.Error {
	if _, err := c.client.Exists(
		ctx,
		&notebookpb.ExistsNotebookRequest{
			Id: id.String(),
		},
	); err != nil {
		return grpc.ParseError(err)
	}

	return nil
}
