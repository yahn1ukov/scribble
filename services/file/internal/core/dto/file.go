package dto

import "io"

type UploadInput struct {
	Name        string
	Size        int64
	ContentType string
	NoteID      string
	Content     []byte
}

type GetOutput struct {
	Name        string
	ContentType string
	Content     io.Reader
}
