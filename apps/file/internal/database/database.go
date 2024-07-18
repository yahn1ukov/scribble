package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Cfg *config.Config
}

func New(p Params) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), p.Cfg.DB.Postgres.URL)
}
