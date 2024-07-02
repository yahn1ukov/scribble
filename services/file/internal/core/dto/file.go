package dto

import "io"

type UploadInput struct {
	Name        string
	Size        int64
	ContentType string
	NoteID      string
	Content     io.Reader
}
