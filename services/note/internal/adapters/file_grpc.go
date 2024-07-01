package adapters

import (
	"context"
	"mime/multipart"

	"github.com/yahn1ukov/scribble/libs/grpc"
	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
)

type FileGRPCClient struct {
	client filepb.FileServiceClient
}

func NewFileGRPCClient(client filepb.FileServiceClient) *FileGRPCClient {
	return &FileGRPCClient{
		client: client,
	}
}

func (c *FileGRPCClient) Create(ctx context.Context, noteId string, files []*multipart.FileHeader) *grpc.Error {
	var metadata []*filepb.Metadata
	for _, file := range files {
		metadata = append(
			metadata,
			&filepb.Metadata{
				Name:        file.Filename,
				Size:        file.Size,
				ContentType: file.Header.Get("Content-Type"),
			},
		)
	}

	if _, err := c.client.Create(ctx,
		&filepb.CreateFileRequest{
			NoteId: noteId,
			Files:  metadata,
		},
	); err != nil {
		return grpc.ParseError(err)
	}

	return nil
}

func (c *FileGRPCClient) GetAll(ctx context.Context, noteId string) ([]*filepb.File, *grpc.Error) {
	files, err := c.client.GetAll(ctx,
		&filepb.GetAllFileRequest{
			NoteId: noteId,
		},
	)
	if err != nil {
		return nil, grpc.ParseError(err)
	}

	return files.Files, nil
}
