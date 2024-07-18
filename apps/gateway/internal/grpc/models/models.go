package models

type CreateNotebookInput struct {
	Title string
}

type UpdateNotebookInput struct {
	Title string
}

type CreateNoteInput struct {
	Title string
	Body  string
}

type UpdateNoteInput struct {
	Title string
	Body  string
}
