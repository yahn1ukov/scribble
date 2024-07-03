package services

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	pb "github.com/yahn1ukov/scribble/libs/grpc/file"
	"github.com/yahn1ukov/scribble/services/file/internal/config"
	"github.com/yahn1ukov/scribble/services/file/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/file/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/file/internal/core/ports"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service struct {
	cfg         *config.Config
	minioClient *minio.Client
	repository  ports.Repository
}

func NewService(cfg *config.Config, minioClient *minio.Client, repository ports.Repository) ports.Service {
	return &service{
		cfg:         cfg,
		minioClient: minioClient,
		repository:  repository,
	}
}

func (s *service) Upload(ctx context.Context, in *dto.UploadInput) error {
	file := &domain.File{
		Name:        in.Name,
		Size:        in.Size,
		ContentType: in.ContentType,
		URL:         fmt.Sprintf("notes/%s/%s", in.NoteID, in.Name),
		NoteID:      in.NoteID,
	}

	content := bytes.NewReader(in.Content)

	if _, err := s.minioClient.PutObject(ctx, s.cfg.Storage.MinIO.Bucket, file.URL, content, file.Size,
		minio.PutObjectOptions{
			ContentType: file.ContentType,
		},
	); err != nil {
		return err
	}

	return s.repository.Create(ctx, file)
}

func (s *service) GetAll(ctx context.Context, noteId string) ([]*pb.File, error) {
	files, err := s.repository.GetAll(ctx, noteId)
	if err != nil {
		return nil, err
	}

	var out []*pb.File
	for _, file := range files {
		out = append(
			out,
			&pb.File{
				Id:          file.ID,
				Name:        file.Name,
				Size:        file.Size,
				ContentType: file.ContentType,
				CreatedAt:   timestamppb.New(file.CreatedAt),
			},
		)
	}

	return out, nil
}

func (s *service) Get(ctx context.Context, id string) (*dto.GetOutput, error) {
	file, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	content, err := s.minioClient.GetObject(ctx, s.cfg.Storage.MinIO.Bucket, file.URL, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	out := &dto.GetOutput{
		Name:        file.Name,
		ContentType: file.ContentType,
		Content:     content,
	}

	return out, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	file, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	if err = s.minioClient.RemoveObject(ctx, s.cfg.Storage.MinIO.Bucket, file.URL, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return s.repository.Delete(ctx, id)
}
