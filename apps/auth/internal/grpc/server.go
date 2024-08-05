package grpc

import (
	"context"
	"github.com/yahn1ukov/scribble/apps/auth/internal/config"
	"github.com/yahn1ukov/scribble/libs/hash"
	"github.com/yahn1ukov/scribble/libs/jwt"
	pb "github.com/yahn1ukov/scribble/proto/auth"
	userpb "github.com/yahn1ukov/scribble/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	cfg        *config.Config
	userClient userpb.UserServiceClient
}

func NewServer(cfg *config.Config, userClient userpb.UserServiceClient) *Server {
	return &Server{
		cfg:        cfg,
		userClient: userClient,
	}
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userClient.FindUser(
		ctx,
		&userpb.FindUserRequest{
			Email: req.Email,
		},
	)
	if err != nil {
		return nil, err
	}

	if !hash.Verify(user.Password, req.Password) {
		return nil, status.Error(codes.InvalidArgument, ErrInvalidPassword.Error())
	}

	token, err := jwt.Generate(user.Id, s.cfg.JWT.Secret, s.cfg.JWT.Expiry)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.userClient.CreateUser(
		ctx,
		&userpb.CreateUserRequest{
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Password:  req.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Generate(user.Id, s.cfg.JWT.Secret, s.cfg.JWT.Expiry)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: token,
	}, nil
}
