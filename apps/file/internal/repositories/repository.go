package repositories

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yahn1ukov/scribble/apps/file/internal/model"
)

type Repository interface {
	Create(context.Context, string, *model.File) error
	GetAll(context.Context, string) ([]*model.File, error)
	GetByID(context.Context, string, string) (*model.File, error)
	Delete(context.Context, string, string) error
}

type repository struct {
	pool *pgxpool.Pool
}

var _ Repository = (*repository)(nil)

func New(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) Create(ctx context.Context, noteID string, file *model.File) error {
	query, args, err := sq.
		Insert("files").
		Columns("note_id", "name", "size", "content_type", "url").
		Values(noteID, file.Name, file.Size, file.ContentType, file.URL).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context, noteID string) ([]*model.File, error) {
	query, args, err := sq.
		Select("id", "name", "size", "content_type", "created_at").
		From("files").
		Where(sq.Eq{"note_id": noteID}).
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

	var files []*model.File
	for rows.Next() {
		var file model.File
		if err = rows.Scan(
			&file.ID,
			&file.Name,
			&file.Size,
			&file.ContentType,
			&file.CreatedAt,
		); err != nil {
			return nil, err
		}

		files = append(files, &file)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *repository) GetByID(ctx context.Context, id string, noteID string) (*model.File, error) {
	query, args, err := sq.
		Select("name", "content_type", "url").
		From("files").
		Where(sq.Eq{"id": id, "note_id": noteID}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var file model.File
	if err = r.pool.
		QueryRow(ctx, query, args...).
		Scan(
			&file.Name,
			&file.ContentType,
			&file.URL,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &file, nil
}

func (r *repository) Delete(ctx context.Context, id string, noteID string) error {
	query, args, err := sq.
		Delete("files").
		Where(sq.Eq{"id": id, "note_id": noteID}).
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
