package models

import "io"

type CreateNotebookRequest struct {
	Title string
}

type UpdateNotebookRequest struct {
	Title string
}

type CreateNoteRequest struct {
	Title string
	Body  string
}

type UpdateNoteRequest struct {
	Title string
	Body  string
}

type UploadFileRequest struct {
	Name        string
	Size        int64
	ContentType string
	Content     io.Reader
}
