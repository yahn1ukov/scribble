package model

import "time"

type File struct {
	ID          string    `db:"id"`
	NoteID      string    `db:"note_id"`
	Name        string    `db:"name"`
	Size        int64     `db:"size"`
	ContentType string    `db:"content_type"`
	URL         string    `db:"url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
