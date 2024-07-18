package clients

import (
	"context"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/libs/grpc"
	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"google.golang.org/grpc/codes"
)

func (c *Client) UploadFile(ctx context.Context, noteID uuid.UUID, file *graphql.Upload) error {
	content, err := io.ReadAll(file.File)
	if err != nil {
		return err
	}

	if _, err := c.file.UploadFile(ctx,
		&filepb.UploadFileRequest{
			Name:        file.Filename,
			Size:        file.Size,
			ContentType: file.ContentType,
			NoteId:      noteID.String(),
			Content:     content,
		},
	); err != nil {
		return grpc.ParseError(err).Error()
	}

	return nil
}

func (c *Client) UploadAllFiles(ctx context.Context, noteID uuid.UUID, files []*graphql.Upload) error {
	stream, err := c.file.UploadAllFiles(ctx)
	if err != nil {
		return grpc.ParseError(err).Error()
	}

	for _, file := range files {
		content, err := io.ReadAll(file.File)
		if err != nil {
			return err
		}

		if err = stream.Send(
			&filepb.UploadFileRequest{
				Name:        file.Filename,
				Size:        file.Size,
				ContentType: file.ContentType,
				NoteId:      noteID.String(),
				Content:     content,
			},
		); err != nil {
			return grpc.ParseError(err).Error()
		}
	}

	if err = stream.CloseSend(); err != nil {
		return err
	}

	if _, err = stream.CloseAndRecv(); err != nil {
		return grpc.ParseError(err).Error()
	}

	return nil
}

func (c *Client) GetAllFiles(ctx context.Context, noteID uuid.UUID) ([]*filepb.File, error) {
	res, err := c.file.GetAllFiles(ctx,
		&filepb.GetAllFilesRequest{
			NoteId: noteID.String(),
		},
	)
	if err != nil {
		return nil, grpc.ParseError(err).Error()
	}

	return res.Files, nil
}

func (c *Client) DownloadFile(ctx context.Context, id uuid.UUID, noteID uuid.UUID) error {
	_, err := c.file.DownloadFile(ctx,
		&filepb.DownloadFileRequest{
			Id:     id.String(),
			NoteId: noteID.String(),
		},
	)
	if err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.NotFound {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}

func (c *Client) RemoveFile(ctx context.Context, id uuid.UUID, noteID uuid.UUID) error {
	if _, err := c.file.RemoveFile(ctx,
		&filepb.RemoveFileRequest{
			Id:     id.String(),
			NoteId: noteID.String(),
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
