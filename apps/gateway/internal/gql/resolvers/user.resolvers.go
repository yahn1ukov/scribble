package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/gqlmodels"
	"github.com/yahn1ukov/scribble/libs/grpc"
	userpb "github.com/yahn1ukov/scribble/proto/user"
)

func (r *mutationResolver) UpdateUser(ctx context.Context, input gqlmodels.UpdateUserInput) (bool, error) {
	id := r.middleware.GetUserIDFromCtx(ctx)

	if _, err := r.userClient.UpdateUser(
		ctx,
		&userpb.UpdateUserRequest{
			Id:        id,
			Email:     input.Email,
			FirstName: input.FirstName,
			LastName:  input.LastName,
		},
	); err != nil {
		return false, grpc.ParseError(err).Error()
	}

	return true, nil
}

func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input gqlmodels.UpdateUserPasswordInput) (bool, error) {
	id := r.middleware.GetUserIDFromCtx(ctx)

	if _, err := r.userClient.UpdateUserPassword(
		ctx,
		&userpb.UpdateUserPasswordRequest{
			Id:          id,
			OldPassword: input.OldPassword,
			NewPassword: input.NewPassword,
		},
	); err != nil {
		return false, grpc.ParseError(err).Error()
	}

	return true, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	id := r.middleware.GetUserIDFromCtx(ctx)

	if _, err := r.userClient.DeleteUser(
		ctx,
		&userpb.DeleteUserRequest{
			Id: id,
		},
	); err != nil {
		return false, grpc.ParseError(err).Error()
	}

	return true, nil
}

func (r *queryResolver) User(ctx context.Context) (*gqlmodels.User, error) {
	id := r.middleware.GetUserIDFromCtx(ctx)

	res, err := r.userClient.GetUser(
		ctx,
		&userpb.GetUserRequest{
			Id: id,
		},
	)
	if err != nil {
		return nil, grpc.ParseError(err).Error()
	}

	output := &gqlmodels.User{
		ID:        uuid.MustParse(res.Id),
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		CreatedAt: res.CreatedAt.AsTime(),
	}

	return output, nil
}
