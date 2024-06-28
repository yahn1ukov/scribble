package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateInput struct {
	Title string `json:"title" binding:"required"`
}

type GetOutput struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateInput struct {
	Title string `json:"title"`
}
