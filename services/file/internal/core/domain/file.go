package domain

import "time"

type File struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Size        int64     `db:"size"`
	ContentType string    `db:"content_type"`
	URL         string    `db:"url"`
	NoteID      string    `db:"note_id"`
	CreatedAt   time.Time `db:"created_at"`
}
