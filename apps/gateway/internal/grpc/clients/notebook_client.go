package clients

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc/models"
	"github.com/yahn1ukov/scribble/libs/grpc"
	notebookpb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) CreateNotebook(ctx context.Context, req *models.CreateNotebookRequest) error {
	if _, err := c.notebook.CreateNotebook(
		ctx,
		&notebookpb.CreateNotebookRequest{
			Title: req.Title,
		},
	); err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.AlreadyExists {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}

func (c *Client) GetAllNotebooks(ctx context.Context) ([]*notebookpb.Notebook, error) {
	res, err := c.notebook.GetAllNotebooks(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, grpc.ParseError(err).Error()
	}

	return res.Notebooks, nil
}

func (c *Client) UpdateNotebook(ctx context.Context, id uuid.UUID, req *models.UpdateNotebookRequest) error {
	if _, err := c.notebook.UpdateNotebook(
		ctx,
		&notebookpb.UpdateNotebookRequest{
			Id:    id.String(),
			Title: req.Title,
		},
	); err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.NotFound {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}

func (c *Client) DeleteNotebook(ctx context.Context, id uuid.UUID) error {
	if _, err := c.notebook.DeleteNotebook(
		ctx,
		&notebookpb.DeleteNotebookRequest{
			Id: id.String(),
		},
	); err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.NotFound {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}
