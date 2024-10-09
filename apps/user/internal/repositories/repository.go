package repositories

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yahn1ukov/scribble/apps/user/internal/model"
)

type Repository interface {
	Create(context.Context, *model.User) (string, error)
	FindByEmail(context.Context, string) (*model.User, error)
	GetByID(context.Context, string) (*model.User, error)
	Update(context.Context, string, map[string]interface{}) error
	UpdatePassword(context.Context, string, *model.User) error
	Delete(context.Context, string) error
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

func (r *repository) Create(ctx context.Context, user *model.User) (string, error) {
	query, args, err := sq.
		Insert("users").
		Columns("email", "first_name", "last_name", "password").
		Values(user.Email, user.FirstName, user.LastName, user.Password).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", err
	}

	var id string
	if err = r.pool.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", ErrAlreadyExists
		}

		return "", err
	}

	return id, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query, args, err := sq.
		Select("id", "password").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	if err = r.pool.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Password); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.User, error) {
	query, args, err := sq.
		Select("id", "email", "first_name", "last_name", "password", "created_at").
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	if err = r.pool.
		QueryRow(ctx, query, args...).
		Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(ctx context.Context, id string, updatedFields map[string]interface{}) error {
	query, args, err := sq.
		Update("users").
		SetMap(updatedFields).
		Where(sq.Eq{"id": id}).
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

func (r *repository) UpdatePassword(ctx context.Context, id string, user *model.User) error {
	query, args, err := sq.
		Update("users").
		Set("password", user.Password).
		Where(sq.Eq{"id": id}).
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

func (r *repository) Delete(ctx context.Context, id string) error {
	query, args, err := sq.
		Delete("users").
		Where(sq.Eq{"id": id}).
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
