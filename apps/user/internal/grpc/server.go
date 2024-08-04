package grpc

import (
	"context"
	"errors"
	"github.com/yahn1ukov/scribble/apps/user/internal/dto"
	"github.com/yahn1ukov/scribble/apps/user/internal/repositories"
	"github.com/yahn1ukov/scribble/apps/user/internal/services"
	pb "github.com/yahn1ukov/scribble/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedUserServiceServer

	service services.Service
}

func NewServer(service services.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	input := &dto.CreateInput{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	id, err := s.service.Create(ctx, input)
	if err != nil {
		if errors.Is(err, services.ErrEmailIsRequired) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, services.ErrPasswordIsRequired) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}

func (s *Server) FindUser(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	user, err := s.service.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := s.service.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := s.service.Update(ctx, req.Id, input); err != nil {
		if errors.Is(err, services.ErrNoFieldsToUpdate); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*emptypb.Empty, error) {
	input := &dto.UpdatePasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	if err := s.service.UpdatePassword(ctx, req.Id, input); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if errors.Is(err, services.ErrOldPasswordIsRequired) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, services.ErrNewPasswordIsRequired) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, services.ErrInvalidPassword) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, services.ErrPasswordsAreSame) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := s.service.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
