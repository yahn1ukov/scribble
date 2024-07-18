package clients

import (
	filepb "github.com/yahn1ukov/scribble/libs/grpc/file"
	notepb "github.com/yahn1ukov/scribble/libs/grpc/note"
	notebookpb "github.com/yahn1ukov/scribble/libs/grpc/notebook"
)

type Client struct {
	notebook notebookpb.NotebookServiceClient
	note     notepb.NoteServiceClient
	file     filepb.FileServiceClient
}

func NewClient(
	notebook notebookpb.NotebookServiceClient,
	note notepb.NoteServiceClient,
	file filepb.FileServiceClient,
) *Client {
	return &Client{
		notebook: notebook,
		note:     note,
		file:     file,
	}
}
