package clients

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc/models"
	"github.com/yahn1ukov/scribble/libs/grpc"
	notepb "github.com/yahn1ukov/scribble/libs/grpc/note"
	"google.golang.org/grpc/codes"
)

func (c *Client) CreateNote(ctx context.Context, notebookID uuid.UUID, input *models.CreateNoteInput) (uuid.UUID, error) {
	res, err := c.note.CreateNote(ctx,
		&notepb.CreateNoteRequest{
			Title:      input.Title,
			Body:       input.Body,
			NotebookId: notebookID.String(),
		},
	)
	if err != nil {
		return uuid.Nil, grpc.ParseError(err).Error()
	}

	return uuid.MustParse(res.Id), nil
}

func (c *Client) GetAllNotes(ctx context.Context, notebookID uuid.UUID) ([]*notepb.Note, error) {
	res, err := c.note.GetAllNotes(ctx,
		&notepb.GetAllNotesRequest{
			NotebookId: notebookID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return res.Notes, nil
}

func (c *Client) GetNote(ctx context.Context, id uuid.UUID, notebookID uuid.UUID) (*notepb.Note, error) {
	res, err := c.note.GetNote(ctx,
		&notepb.GetNoteRequest{
			Id:         id.String(),
			NotebookId: notebookID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UpdateNote(ctx context.Context, id uuid.UUID, notebookID uuid.UUID, input *models.UpdateNoteInput) error {
	if _, err := c.note.UpdateNote(ctx,
		&notepb.UpdateNoteRequest{
			Id:         id.String(),
			Title:      input.Title,
			Body:       input.Body,
			NotebookId: notebookID.String(),
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

func (c *Client) DeleteNote(ctx context.Context, id uuid.UUID, notebookID uuid.UUID) error {
	if _, err := c.note.DeleteNote(ctx,
		&notepb.DeleteNoteRequest{
			Id:         id.String(),
			NotebookId: notebookID.String(),
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
