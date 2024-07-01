package domain

import "time"

type Notebook struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}
