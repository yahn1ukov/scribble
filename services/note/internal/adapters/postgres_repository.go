package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yahn1ukov/scribble/services/note/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/note/internal/core/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, notebookId string, note *domain.Note) (string, error) {
	query := "INSERT INTO notes(title, body, notebook_id) VALUES($1, $2, $3) RETURNING id"

	var id string
	if err := r.db.
		QueryRowContext(ctx, query, note.Title, note.Body, notebookId).
		Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *postgresRepository) GetAll(ctx context.Context, notebookId string) ([]*domain.Note, error) {
	query := "SELECT id, title, body, created_at FROM notes WHERE notebook_id = $1"

	rows, err := r.db.QueryContext(ctx, query, notebookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*domain.Note
	for rows.Next() {
		var note domain.Note
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

func (r *postgresRepository) Get(ctx context.Context, id string, notebookId string) (*domain.Note, error) {
	query := "SELECT id, title, body, created_at FROM notes WHERE id = $1 AND notebook_id = $2 LIMIT 1"

	var note domain.Note
	if err := r.db.
		QueryRowContext(ctx, query, id, notebookId).
		Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &note, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, notebookId string, note *domain.Note) error {
	query := "UPDATE notes SET title = $1, body = $2 WHERE id = $3 AND notebook_id = $4"

	result, err := r.db.ExecContext(ctx, query, note.Title, note.Body, id, notebookId)
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

func (r *postgresRepository) Delete(ctx context.Context, id string, notebookId string) error {
	query := "DELETE FROM notes WHERE id = $1 AND notebook_id = $2"

	result, err := r.db.ExecContext(ctx, query, id, notebookId)
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
