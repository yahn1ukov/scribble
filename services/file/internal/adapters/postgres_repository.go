package adapters

import (
	"context"
	"database/sql"

	"github.com/yahn1ukov/scribble/services/file/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/file/internal/core/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, file *domain.File) error {
	query := "INSERT INTO files(name, size, content_type, url, note_id) VALUES($1, $2, $3, $4, $5)"

	if _, err := r.db.ExecContext(ctx, query, file.Name, file.Size, file.ContentType, file.URL, file.NoteID); err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) GetAll(ctx context.Context, noteId string) ([]*domain.File, error) {
	query := "SELECT id, name, size, content_type, url, created_at FROM files WHERE note_id = $1"

	rows, err := r.db.QueryContext(ctx, query, noteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*domain.File
	for rows.Next() {
		var file domain.File
		if err := rows.Scan(&file.ID, &file.Name, &file.Size, &file.ContentType, &file.URL, &file.CreatedAt); err != nil {
			return nil, err
		}
		files = append(files, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}
