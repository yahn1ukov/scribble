package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/model"
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{
		pool: pool,
	}
}

func (r *postgresRepository) Create(ctx context.Context, notebook *model.Notebook) error {
	query := "INSERT INTO notebooks(title) VALUES($1)"

	if _, err := r.pool.Exec(ctx, query, notebook.Title); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*model.Notebook, error) {
	query := "SELECT id, title, created_at FROM notebooks"

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notebooks []*model.Notebook
	for rows.Next() {
		var notebook model.Notebook
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

func (r *postgresRepository) Get(ctx context.Context, id string) (*model.Notebook, error) {
	query := "SELECT id, title, created_at FROM notebooks WHERE id = $1 LIMIT 1"

	var notebook model.Notebook
	if err := r.pool.QueryRow(ctx, query, id).Scan(&notebook.ID, &notebook.Title, &notebook.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &notebook, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, notebook *model.Notebook) error {
	query := "UPDATE notebooks SET title = $1 WHERE id = $2"

	result, err := r.pool.Exec(ctx, query, notebook.Title, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM notebooks WHERE id = $1"

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
