package domain

import "time"

type Note struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}
