package user

import (
	"context"
	"net/http"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
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

func (s *service) Create(ctx context.Context, req *spotigraph.UserCreateReq) (*spotigraph.SingleUserRes, error) {
	res := &spotigraph.SingleUserRes{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Prepare user creation
	entity := models.NewUser(req.Principal)

	// Validate entity
	if err := entity.Validate(); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "Unable to prepare user object",
		}
		return res, err
	}

	// Create use in database
	if err := s.users.Create(ctx, entity); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create user",
		}
		return res, err
	}

	// Prepare response
	res.Entity = spotigraph.FromUser(entity)

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *spotigraph.UserGetReq) *spotigraph.SingleUserRes {
	panic("not implemented")
}

func (s *service) Update(ctx context.Context, req *spotigraph.UserCreateReq) *spotigraph.SingleUserRes {
	panic("not implemented")
}

func (s *service) Delete(ctx context.Context, req *spotigraph.UserGetReq) *spotigraph.EmptyRes {
	panic("not implemented")
}
