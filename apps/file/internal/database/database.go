package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	"go.uber.org/fx"
)

func New(cfg *config.Config) (*sql.DB, error) {
	return sql.Open(
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
}

func Run(lc fx.Lifecycle, db *sql.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(_ context.Context) error {
			return db.Close()
		},
	})
}
