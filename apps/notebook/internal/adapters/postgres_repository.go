package adapters

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/domain"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Exists(ctx context.Context, id uuid.UUID) error {
	query := "SELECT CASE WHEN EXISTS (SELECT 1 FROM notebooks WHERE id = $1) THEN TRUE ELSE FALSE END"

	var exists bool
	if err := r.db.
		QueryRowContext(ctx, query, id).
		Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return domain.ErrNotFound
	}

	return nil
}

func (r *postgresRepository) Create(ctx context.Context, notebook *domain.Notebook) error {
	query := "INSERT INTO notebooks(title) VALUES($1)"

	if _, err := r.db.ExecContext(ctx, query, notebook.Title); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return domain.ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.Notebook, error) {
	query := "SELECT * FROM notebooks"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notebooks []*domain.Notebook
	for rows.Next() {
		var notebook domain.Notebook
		if err := rows.Scan(&notebook.ID, &notebook.Title, &notebook.CreatedAt); err != nil {
			return nil, err
		}
		notebooks = append(notebooks, &notebook)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notebooks, nil
}

func (r *postgresRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Notebook, error) {
	query := "SELECT * FROM notebooks WHERE id = $1 LIMIT 1"

	var notebook domain.Notebook
	if err := r.db.
		QueryRowContext(ctx, query, id).
		Scan(&notebook.ID, &notebook.Title, &notebook.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &notebook, nil
}

func (r *postgresRepository) Update(ctx context.Context, id uuid.UUID, notebook *domain.Notebook) error {
	query := "UPDATE notebooks SET title = $1 WHERE id = $2"

	result, err := r.db.ExecContext(ctx, query, notebook.Title, id)
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

func (r *postgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM notebooks WHERE id = $1"

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
