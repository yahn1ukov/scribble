package dto

import "time"

type CreateInput struct {
	Title string `json:"title" binding:"required"`
}

type GetOutput struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateInput struct {
	Title string `json:"title"`
}
