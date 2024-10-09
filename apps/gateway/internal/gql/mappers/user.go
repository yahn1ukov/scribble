package mappers

import (
	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/gqlmodels"
	userpb "github.com/yahn1ukov/scribble/proto/user"
)

func (m *Mapper) GRPCUserToUser(user *userpb.User) gqlmodels.User {
	return gqlmodels.User{
		ID:        uuid.MustParse(user.Id),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.AsTime(),
	}
}
