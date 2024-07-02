package app

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yahn1ukov/scribble/services/file/internal/config"
	"go.uber.org/fx"
)

func NewMinIO(cfg *config.Config) (*minio.Client, error) {
	client, err := minio.New(
		cfg.Storage.MinIO.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.Storage.MinIO.AccessKey, cfg.Storage.MinIO.SecretKey, ""),
			Secure: cfg.Storage.MinIO.UseSSL,
		},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func RunMinIO(lc fx.Lifecycle, cfg *config.Config, client *minio.Client) {
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
