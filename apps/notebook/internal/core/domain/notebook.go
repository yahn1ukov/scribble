package domain

import (
	"time"

	"github.com/google/uuid"
)

type Notebook struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}
