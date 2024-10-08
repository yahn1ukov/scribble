package repositories

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yahn1ukov/scribble/apps/note/internal/model"
)

type Repository interface {
	Create(context.Context, string, *model.Note) (string, error)
	GetAll(context.Context, string) ([]*model.Note, error)
	GetByID(context.Context, string, string) (*model.Note, error)
	Update(context.Context, string, string, map[string]interface{}) error
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

func (r *repository) Create(ctx context.Context, notebookID string, note *model.Note) (string, error) {
	query, args, err := sq.
		Insert("notes").
		Columns("notebook_id", "title", "content").
		Values(notebookID, note.Title, note.Content).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", err
	}

	var id string
	if err = r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) GetAll(ctx context.Context, notebookID string) ([]*model.Note, error) {
	query, args, err := sq.
		Select("id", "title", "content", "created_at").
		From("notes").
		Where(sq.Eq{"notebook_id": notebookID}).
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

	var notes []*model.Note
	for rows.Next() {
		var note model.Note
		if err = rows.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
		); err != nil {
			return nil, err
		}

		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) GetByID(ctx context.Context, id string, notebookID string) (*model.Note, error) {
	query, args, err := sq.
		Select("id", "title", "content", "created_at").
		From("notes").
		Where(sq.Eq{"id": id, "notebook_id": notebookID}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var note model.Note
	if err = r.pool.
		QueryRow(ctx, query, args...).
		Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &note, nil
}

func (r *repository) Update(ctx context.Context, id string, notebookID string, updatedFields map[string]interface{}) error {
	query, args, err := sq.
		Update("notes").
		SetMap(updatedFields).
		Where(sq.Eq{"id": id, "notebook_id": notebookID}).
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

func (r *repository) Delete(ctx context.Context, id string, notebookID string) error {
	query, args, err := sq.
		Delete("notes").
		Where(sq.Eq{"id": id, "notebook_id": notebookID}).
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
