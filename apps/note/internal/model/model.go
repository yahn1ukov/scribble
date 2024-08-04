package model

import "time"

type Note struct {
	ID         string    `db:"id"`
	NotebookID string    `db:"notebook_id"`
	Title      string    `db:"title"`
	Content    *string   `db:"body"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
