package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yahn1ukov/scribble/apps/note/internal/model"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, notebookID string, note *model.Note) (string, error) {
	query := "INSERT INTO notes(title, body, notebook_id) VALUES($1, $2, $3) RETURNING id"

	var id string
	if err := r.db.QueryRowContext(ctx, query, note.Title, note.Body, notebookID).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *postgresRepository) GetAll(ctx context.Context, notebookID string) ([]*model.Note, error) {
	query := "SELECT id, title, body, created_at FROM notes WHERE notebook_id = $1"

	rows, err := r.db.QueryContext(ctx, query, notebookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *postgresRepository) Get(ctx context.Context, id string, notebookID string) (*model.Note, error) {
	query := "SELECT id, title, body, created_at FROM notes WHERE id = $1 AND notebook_id = $2 LIMIT 1"

	var note model.Note
	if err := r.db.QueryRowContext(ctx, query, id, notebookID).Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &note, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, notebookID string, note *model.Note) error {
	query := "UPDATE notes SET title = $1, body = $2 WHERE id = $3 AND notebook_id = $4"

	result, err := r.db.ExecContext(ctx, query, note.Title, note.Body, id, notebookID)
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

func (r *postgresRepository) Delete(ctx context.Context, id string, notebookID string) error {
	query := "DELETE FROM notes WHERE id = $1 AND notebook_id = $2"

	result, err := r.db.ExecContext(ctx, query, id, notebookID)
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
