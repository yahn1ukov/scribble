package app

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yahn1ukov/scribble/apps/note/internal/config"
	"go.uber.org/fx"
)

func NewPostgres(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		cfg.DB.Postgres.Driver,
		fmt.Sprintf(
			"%s://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.DB.Postgres.Driver,
			cfg.DB.Postgres.User,
			cfg.DB.Postgres.Password,
			cfg.DB.Postgres.Host,
			cfg.DB.Postgres.Port,
			cfg.DB.Postgres.Name,
			cfg.DB.Postgres.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunPostgres(lc fx.Lifecycle, db *sql.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(_ context.Context) error {
			return db.Close()
		},
	})
}
