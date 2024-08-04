package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	"go.uber.org/fx"
)

func New(cfg *config.Config) (*minio.Client, error) {
	return minio.New(
		cfg.Storage.MinIO.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.Storage.MinIO.AccessKey, cfg.Storage.MinIO.SecretKey, ""),
			Secure: cfg.Storage.MinIO.UseSSL,
		},
	)
}

func Run(lc fx.Lifecycle, cfg *config.Config, client *minio.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := client.MakeBucket(
				ctx,
				cfg.Storage.MinIO.Bucket,
				minio.MakeBucketOptions{
					Region: cfg.Storage.MinIO.Region,
				},
			); err != nil {
				exists, errBucketExists := client.BucketExists(ctx, cfg.Storage.MinIO.Bucket)
				if errBucketExists != nil && !exists {
					return err
				}
			}

			return nil
		},
	})
}
