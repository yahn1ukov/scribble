package adapters

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/yahn1ukov/scribble/libs/grpc"
	storagepb "github.com/yahn1ukov/scribble/libs/grpc/storage"
	"google.golang.org/grpc/codes"
)

type StorageGRPCClient struct {
	client storagepb.StorageServiceClient
}

func NewStorageGRPCClient(client storagepb.StorageServiceClient) *StorageGRPCClient {
	return &StorageGRPCClient{
		client: client,
	}
}

func (c *StorageGRPCClient) Upload(ctx context.Context, noteId string, files []*multipart.FileHeader) *grpc.Error {
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
			&storagepb.UploadFileRequest{
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
