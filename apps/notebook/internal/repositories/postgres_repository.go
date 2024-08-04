package repositories

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *postgresRepository) Create(ctx context.Context, userID string, notebook *model.Notebook) error {
	query, args, err := sq.
		Insert("notebooks").
		Columns("user_id", "title", "description").
		Values(userID, notebook.Title, notebook.Description).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, query, args...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (r *postgresRepository) GetAll(ctx context.Context, userID string) ([]*model.Notebook, error) {
	query, args, err := sq.
		Select("id", "title", "description", "created_at").
		From("notebooks").
		Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notebooks []*model.Notebook
	for rows.Next() {
		var notebook model.Notebook
		if err = rows.Scan(
			&notebook.ID,
			&notebook.Title,
			&notebook.Description,
			&notebook.CreatedAt,
		); err != nil {
			return nil, err
		}

		notebooks = append(notebooks, &notebook)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notebooks, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, userID string, updatedFields map[string]interface{}) error {
	query, args, err := sq.
		Update("notebooks").
		SetMap(updatedFields).
		Where(sq.Eq{"id": id, "user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	result, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrAlreadyExists
		}

		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string, userID string) error {
	query, args, err := sq.
		Delete("notebooks").
		Where(sq.Eq{"id": id, "user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	result, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
