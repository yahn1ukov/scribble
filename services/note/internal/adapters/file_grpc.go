package adapters

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/yahn1ukov/scribble/libs/grpc"
	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"github.com/yahn1ukov/scribble/services/note/internal/core/dto"
	"google.golang.org/grpc/codes"
)

type FileGRPCClient struct {
	client filepb.FileServiceClient
}

func NewFileGRPCClient(client filepb.FileServiceClient) *FileGRPCClient {
	return &FileGRPCClient{
		client: client,
	}
}

func (c *FileGRPCClient) Upload(ctx context.Context, noteId string, files []*multipart.FileHeader) *grpc.Error {
	stream, err := c.client.Upload(ctx)
	if err != nil {
		return grpc.ParseError(err)
	}

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return grpc.CreateError(codes.Internal, err.Error())
		}
		defer f.Close()

		content, err := io.ReadAll(f)
		if err != nil {
			return grpc.CreateError(codes.Internal, err.Error())
		}

		if err = stream.Send(
			&filepb.UploadFileRequest{
				Name:        file.Filename,
				Size:        file.Size,
				ContentType: file.Header.Get("Content-Type"),
				NoteId:      noteId,
				Content:     content,
			},
		); err != nil {
			return grpc.ParseError(err)
		}
	}

	if err = stream.CloseSend(); err != nil {
		return grpc.CreateError(codes.Internal, err.Error())
	}

	if _, err = stream.CloseAndRecv(); err != nil {
		return grpc.ParseError(err)
	}

	return nil
}

func (c *FileGRPCClient) GetAll(ctx context.Context, noteId string) ([]*dto.File, *grpc.Error) {
	files, err := c.client.GetAll(ctx,
		&filepb.GetAllFileRequest{
			NoteId: noteId,
		},
	)
	if err != nil {
		return nil, grpc.ParseError(err)
	}

	var out []*dto.File
	for _, file := range files.Files {
		out = append(
			out,
			&dto.File{
				ID:          file.Id,
				Name:        file.Name,
				Size:        file.Size,
				ContentType: file.ContentType,
				URL:         file.Url,
				CreatedAt:   file.CreatedAt.AsTime(),
			},
		)
	}

	return out, nil
}
