package dto

import (
	"mime/multipart"
	"time"
)

type CreateInput struct {
	Title string                  `form:"title" binding:"required"`
	Body  string                  `form:"body"`
	Files []*multipart.FileHeader `form:"files"`
}

type File struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetOutput struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Files     []*File   `json:"files,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
