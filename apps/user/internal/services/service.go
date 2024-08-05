package services

import (
	"context"
	"github.com/yahn1ukov/scribble/apps/user/internal/dto"
	"github.com/yahn1ukov/scribble/apps/user/internal/model"
	"github.com/yahn1ukov/scribble/apps/user/internal/repositories"
	"github.com/yahn1ukov/scribble/libs/hash"
	pb "github.com/yahn1ukov/scribble/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Create(context.Context, *dto.CreateInput) (string, error)
	FindByEmail(context.Context, string) (*pb.FindUserResponse, error)
	GetByID(context.Context, string) (*pb.User, error)
	Update(context.Context, string, *dto.UpdateInput) error
	UpdatePassword(context.Context, string, *dto.UpdatePasswordInput) error
	Delete(context.Context, string) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, input *dto.CreateInput) (string, error) {
	if input.Email == "" {
		return "", ErrEmailIsRequired
	}

	if input.Password == "" {
		return "", ErrPasswordIsRequired
	}

	hashedPassword, err := hash.Hash(input.Password)
	if err != nil {
		return "", err
	}

	user := &model.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  hashedPassword,
	}

	return s.repository.Create(ctx, user)
}

func (s *service) FindByEmail(ctx context.Context, email string) (*pb.FindUserResponse, error) {
	user, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	output := &pb.FindUserResponse{
		Id:       user.ID,
		Password: user.Password,
	}

	return output, nil
}

func (s *service) GetByID(ctx context.Context, id string) (*pb.User, error) {
	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	output := &pb.User{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}

	return output, nil
}

func (s *service) Update(ctx context.Context, id string, input *dto.UpdateInput) error {
	updatedFields := make(map[string]interface{})

	if input.Email != nil && *input.Email != "" {
		updatedFields["email"] = *input.Email
	}

	if input.FirstName != nil {
		updatedFields["first_name"] = *input.FirstName
	}

	if input.LastName != nil {
		updatedFields["last_name"] = *input.LastName
	}

	if len(updatedFields) == 0 {
		return ErrNoFieldsToUpdate
	}

	return s.repository.Update(ctx, id, updatedFields)
}

func (s *service) UpdatePassword(ctx context.Context, id string, input *dto.UpdatePasswordInput) error {
	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if input.OldPassword == "" {
		return ErrOldPasswordIsRequired
	}

	if input.NewPassword == "" {
		return ErrNewPasswordIsRequired
	}

	if !hash.Verify(user.Password, input.OldPassword) {
		return ErrInvalidPassword
	}

	if input.NewPassword == input.OldPassword {
		return ErrPasswordsAreSame
	}

	hashedPassword, err := hash.Hash(input.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return s.repository.UpdatePassword(ctx, id, user)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
