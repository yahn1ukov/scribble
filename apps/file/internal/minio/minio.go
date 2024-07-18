package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc     fx.Lifecycle
	Cfg    *config.Config
	Client *minio.Client
}

func New(p Params) (*minio.Client, error) {
	return minio.New(
		p.Cfg.Storage.MinIO.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(p.Cfg.Storage.MinIO.AccessKey, p.Cfg.Storage.MinIO.SecretKey, ""),
			Secure: p.Cfg.Storage.MinIO.UseSSL,
		},
	)
}

func Run(p Params) {
	p.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := p.Client.MakeBucket(
				ctx,
				p.Cfg.Storage.MinIO.Bucket,
				minio.MakeBucketOptions{
					Region: p.Cfg.Storage.MinIO.Region,
				},
			); err != nil {
				exists, errBucketExists := p.Client.BucketExists(ctx, p.Cfg.Storage.MinIO.Bucket)
				if errBucketExists != nil && !exists {
					return err
				}
			}
			return nil
		},
	})
}
