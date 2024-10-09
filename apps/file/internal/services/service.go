package services

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/yahn1ukov/scribble/apps/file/internal/config"
	"github.com/yahn1ukov/scribble/apps/file/internal/dto"
	"github.com/yahn1ukov/scribble/apps/file/internal/model"
	"github.com/yahn1ukov/scribble/apps/file/internal/repositories"
	pb "github.com/yahn1ukov/scribble/proto/file"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	GetAll(context.Context, string) ([]*pb.File, error)
	GetByID(context.Context, string, string) (*pb.DownloadFileResponse, error)
	Upload(context.Context, string, *dto.UploadInput) error
	Remove(context.Context, string, string) error
}

type service struct {
	repository  repositories.Repository
	cfg         *config.Config
	minioClient *minio.Client
}

var _ Service = (*service)(nil)

func New(repository repositories.Repository, cfg *config.Config, minioClient *minio.Client) *service {
	return &service{
		repository:  repository,
		cfg:         cfg,
		minioClient: minioClient,
	}
}

func (s *service) Upload(ctx context.Context, noteID string, input *dto.UploadInput) error {
	file := &model.File{
		Name:        input.Name,
		Size:        input.Size,
		ContentType: input.ContentType,
		URL:         fmt.Sprintf("notes/%s/%s", noteID, input.Name),
	}

	content := bytes.NewReader(input.Content)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if _, err := s.minioClient.PutObject(ctx, s.cfg.Storage.MinIO.Bucket, file.URL, content, file.Size,
			minio.PutObjectOptions{
				ContentType: file.ContentType,
			},
		); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		return s.repository.Create(ctx, noteID, file)
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context, noteID string) ([]*pb.File, error) {
	files, err := s.repository.GetAll(ctx, noteID)
	if err != nil {
		return nil, err
	}

	var output []*pb.File
	for _, file := range files {
		output = append(
			output,
			&pb.File{
				Id:          file.ID,
				Name:        file.Name,
				Size:        file.Size,
				ContentType: file.ContentType,
				CreatedAt:   timestamppb.New(file.CreatedAt),
			},
		)
	}

	return output, nil
}

func (s *service) GetByID(ctx context.Context, id string, noteID string) (*pb.DownloadFileResponse, error) {
	file, err := s.repository.GetByID(ctx, id, noteID)
	if err != nil {
		return nil, err
	}

	object, err := s.minioClient.GetObject(ctx, s.cfg.Storage.MinIO.Bucket, file.URL, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	output := &pb.DownloadFileResponse{
		Name:        file.Name,
		ContentType: file.ContentType,
		Content:     content,
	}

	return output, nil
}

func (s *service) Remove(ctx context.Context, id string, noteID string) error {
	file, err := s.repository.GetByID(ctx, id, noteID)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.minioClient.RemoveObject(ctx, s.cfg.Storage.MinIO.Bucket, file.URL, minio.RemoveObjectOptions{})
	})

	g.Go(func() error {
		return s.repository.Delete(ctx, id, noteID)
	})

	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}
