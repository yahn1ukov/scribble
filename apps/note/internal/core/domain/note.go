package domain

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}
