package adapters

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/yahn1ukov/scribble/services/storage/internal/config"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/storage/internal/core/ports"
)

type minioRepository struct {
	cfg    *config.Config
	client *minio.Client
}

func NewMinIORepository(cfg *config.Config, client *minio.Client) ports.Repository {
	return &minioRepository{
		cfg:    cfg,
		client: client,
	}
}

func (r *minioRepository) Upload(ctx context.Context, file *domain.File) error {
	if _, err := r.client.PutObject(ctx, r.cfg.Storage.MinIO.Bucket, file.URL, file.Content, file.Size,
		minio.PutObjectOptions{
			ContentType: file.ContentType,
		},
	); err != nil {
		return err
	}

	return nil
}
