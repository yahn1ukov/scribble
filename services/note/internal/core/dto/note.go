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

type GetOutput struct {
	ID        string                  `json:"id"`
	Title     string                  `json:"title"`
	Body      string                  `json:"body"`
	Files     []*multipart.FileHeader `json:"files,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
}

type UpdateInput struct {
	Title string                  `form:"title"`
	Body  string                  `form:"body"`
	Files []*multipart.FileHeader `form:"files"`
}
