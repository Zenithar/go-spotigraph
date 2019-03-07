package user

import (
	"context"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type service struct {
	users repositories.User
}

// New returns a service instance
func New(users repositories.User) services.User {
	return &service{
		users: users,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *spotigraph.UserCreateReq) spotigraph.SingleUserRes {
	panic("not implemented")
}

func (s *service) Get(ctx context.Context, req *spotigraph.UserGetReq) spotigraph.SingleUserRes {
	panic("not implemented")
}

func (s *service) Update(ctx context.Context, req *spotigraph.UserCreateReq) spotigraph.SingleUserRes {
	panic("not implemented")
}

func (s *service) Delete(ctx context.Context, req *spotigraph.UserGetReq) spotigraph.EmptyRes {
	panic("not implemented")
}
