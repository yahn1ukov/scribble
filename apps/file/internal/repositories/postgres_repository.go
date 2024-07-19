package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yahn1ukov/scribble/apps/file/internal/model"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, noteID string, file *model.File) error {
	query := "INSERT INTO files(name, size, content_type, url, note_id) VALUES($1, $2, $3, $4, $5)"

	if _, err := r.db.ExecContext(ctx, query, file.Name, file.Size, file.ContentType, file.URL, noteID); err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) GetAll(ctx context.Context, noteID string) ([]*model.File, error) {
	query := "SELECT id, name, size, content_type, created_at FROM files WHERE note_id = $1"

	rows, err := r.db.QueryContext(ctx, query, noteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*model.File
	for rows.Next() {
		var file model.File
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

func (r *postgresRepository) Get(ctx context.Context, id string, noteID string) (*model.File, error) {
	query := "SELECT name, content_type, url FROM files WHERE id = $1 AND note_id = $2 LIMIT 1"

	var file model.File
	if err := r.db.QueryRowContext(ctx, query, id, noteID).Scan(&file.Name, &file.ContentType, &file.URL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &file, nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string, noteID string) error {
	query := "DELETE FROM files WHERE id = $1 AND note_id = $2"

	result, err := r.db.ExecContext(ctx, query, id, noteID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
