package adapters

import (
	"context"
	"database/sql"
	"errors"

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
	query := "SELECT id, name, size, content_type, created_at FROM files WHERE note_id = $1"

	rows, err := r.db.QueryContext(ctx, query, noteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*domain.File
	for rows.Next() {
		var file domain.File
		if err := rows.Scan(&file.ID, &file.Name, &file.Size, &file.ContentType, &file.CreatedAt); err != nil {
			return nil, err
		}
		files = append(files, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *postgresRepository) Get(ctx context.Context, id string) (*domain.File, error) {
	query := "SELECT name, content_type, url FROM files WHERE id = $1 LIMIT 1"

	var file domain.File
	if err := r.db.
		QueryRowContext(ctx, query, id).
		Scan(&file.Name, &file.ContentType, &file.URL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &file, nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM files WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
