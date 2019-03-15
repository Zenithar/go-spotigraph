package tribe

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/db"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type service struct {
	Tribes repositories.Tribe
}

// New returns a service instance
func New(Tribes repositories.Tribe) services.Tribe {
	return &service{
		Tribes: Tribes,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *spotigraph.TribeCreateReq) (*spotigraph.SingleTribeRes, error) {
	res := &spotigraph.SingleTribeRes{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Name must be unique
		constraints.TribeNameMustBeUnique(s.Tribes, req.Name),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Prepare Tribe creation
	entity := models.NewTribe(req.Name)

	// Validate entity
	if err := entity.Validate(); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "Unable to prepare Tribe object",
		}
		return res, err
	}

	// Create use in database
	if err := s.Tribes.Create(ctx, entity); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create Tribe",
		}
		return res, err
	}

	// Prepare response
	res.Entity = spotigraph.FromTribe(entity)

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *spotigraph.TribeGetReq) (*spotigraph.SingleTribeRes, error) {
	res := &spotigraph.SingleTribeRes{}

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

	// Retrieve Tribe from database
	entity, err := s.Tribes.Get(ctx, req.Id)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve Tribe",
		}
		return res, err
	}
	if entity == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusNotFound,
			Message: "Tribe not found",
		}
		return res, db.ErrNoResult
	}

	// Prepare response
	res.Entity = spotigraph.FromTribe(entity)

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *spotigraph.TribeUpdateReq) (*spotigraph.SingleTribeRes, error) {
	res := &spotigraph.SingleTribeRes{}

	// Prepare expected results
	var (
		entity models.Tribe
	)

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Tribe must exists
		constraints.TribeMustExists(s.Tribes, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	updated := false

	if req.Name != nil {
		if err := constraints.Validate(ctx,
			// Check acceptable name value
			constraints.MustBeAName(req.Name.Value),
			// Is already used ?
			constraints.TribeNameMustBeUnique(s.Tribes, req.Name.Value),
		); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusConflict,
				Message: "Tribe name already used",
			}
			return res, err
		}
		entity.Name = req.Name.Value
		updated = true
	}

	// Skip operation when no updates
	if updated {

		// Create account in database
		if err := s.Tribes.Update(ctx, &entity); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to update Tribe object",
			}
			return res, err
		}

	}

	// Prepare response
	res.Entity = spotigraph.FromTribe(&entity)

	// Return expected result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *spotigraph.TribeGetReq) (*spotigraph.EmptyRes, error) {
	res := &spotigraph.EmptyRes{}

	// Prepare expected results
	var (
		entity models.Tribe
	)

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Tribe must exists
		constraints.TribeMustExists(s.Tribes, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	if err := s.Tribes.Delete(ctx, req.Id); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to delete Tribe object",
		}
		return res, err
	}

	// Return expected result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *spotigraph.TribeSearchReq) (*spotigraph.PaginatedTribeRes, error) {
	res := &spotigraph.PaginatedTribeRes{}

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
	filter := &repositories.TribeSearchFilter{}
	if req.TribeId != nil {
		filter.TribeID = req.TribeId.Value
	}
	if req.Slug != nil {
		filter.Slug = req.Slug.Value
	}
	if req.Name != nil {
		filter.Name = req.Name.Value
	}

	// Do the search
	entities, total, err := s.Tribes.Search(ctx, filter, pagination, sortParams)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
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
	if err != db.ErrNoResult {
		res.Members = spotigraph.FromTribes(entities)
	}

	// Return results
	return res, nil
}
