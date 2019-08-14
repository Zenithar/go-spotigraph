package user

import (
	"context"
	"fmt"
	"net/http"

	"go.zenithar.org/pkg/db"

	"go.zenithar.org/spotigraph/internal/helpers"
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

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Principal must be unique
		constraints.UserPrincipalMustBeUnique(s.users, req.Principal),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	// Prepare user creation
	entity := models.NewUser(req.Principal)

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

func (s *service) Get(ctx context.Context, req *spotigraph.UserGetReq) (*spotigraph.SingleUserRes, error) {
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

	// Retrieve user from database
	entity, err := s.users.Get(ctx, req.Id)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve user",
		}
		return res, err
	}
	if entity == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}
		return res, db.ErrNoResult
	}

	// Prepare response
	res.Entity = spotigraph.FromUser(entity)

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *spotigraph.UserUpdateReq) (*spotigraph.SingleUserRes, error) {
	res := &spotigraph.SingleUserRes{}

	// Prepare expected results

	var entity models.User

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// User must exists
		constraints.UserMustExists(s.users, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	updated := false

	if req.Principal != nil {
		// Compute hash first
		hash := helpers.PrincipalHashFunc(req.Principal.Value)

		// Check usage
		if err := constraints.Validate(ctx,
			// Principal must be unique
			constraints.UserPrincipalMustBeUnique(s.users, hash),
		); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusConflict,
				Message: "Principal already used",
			}
			return res, err
		}
		entity.Principal = hash
		updated = true
	}

	// Skip operation when no updates
	if updated {
		// Create account in database
		if err := s.users.Update(ctx, &entity); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to update user object",
			}
			return res, err
		}
	}

	// Prepare response
	res.Entity = spotigraph.FromUser(&entity)

	// Return expected result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *spotigraph.UserGetReq) (*spotigraph.EmptyRes, error) {
	res := &spotigraph.EmptyRes{}

	// Prepare expected results

	var entity models.User

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// User must exists
		constraints.UserMustExists(s.users, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	if err := s.users.Delete(ctx, req.Id); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to delete user object",
		}
		return res, err
	}

	// Return expected result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *spotigraph.UserSearchReq) (*spotigraph.PaginatedUserRes, error) {
	res := &spotigraph.PaginatedUserRes{}

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

	// Prepare request parameters
	sortParams := db.SortConverter(req.Sorts)
	pagination := db.NewPaginator(uint(req.Page), uint(req.PerPage))

	// Build search filter
	filter := &repositories.UserSearchFilter{}
	if req.UserId != nil {
		filter.UserID = req.UserId.Value
	}
	if req.Principal != nil {
		filter.Principal = req.Principal.Value
	}

	// Do the search
	entities, total, err := s.users.Search(ctx, filter, pagination, sortParams)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to process request",
		}
		return res, err
	}

	// Set pagination total for paging calcul
	pagination.SetTotal(uint(total))
	res.Total = uint32(pagination.Total())
	res.Count = uint32(pagination.CurrentPageCount())
	res.PerPage = uint32(pagination.PerPage)
	res.CurrentPage = uint32(pagination.Page)

	// If no result back to first page
	res.Members = spotigraph.FromUsers(entities)

	// Return results
	return res, nil
}
