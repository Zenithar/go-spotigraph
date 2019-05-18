package guild

import (
	"context"
	"fmt"
	"net/http"

	"go.zenithar.org/pkg/db"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type service struct {
	guilds repositories.Guild
}

// New returns a service instance
func New(guilds repositories.Guild) services.Guild {
	return &service{
		guilds: guilds,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *spotigraph.GuildCreateReq) (*spotigraph.SingleGuildRes, error) {
	res := &spotigraph.SingleGuildRes{}

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
		// Name must be unique
		constraints.GuildNameMustBeUnique(s.guilds, req.Name),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	// Prepare Guild creation
	entity := models.NewGuild(req.Name)

	// Create use in database
	if err := s.guilds.Create(ctx, entity); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create Guild",
		}
		return res, err
	}

	// Prepare response
	res.Entity = spotigraph.FromGuild(entity)

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *spotigraph.GuildGetReq) (*spotigraph.SingleGuildRes, error) {
	res := &spotigraph.SingleGuildRes{}

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

	// Retrieve Guild from database
	entity, err := s.guilds.Get(ctx, req.Id)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve Guild",
		}
		return res, err
	}
	if entity == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusNotFound,
			Message: "Guild not found",
		}
		return res, db.ErrNoResult
	}

	// Prepare response
	res.Entity = spotigraph.FromGuild(entity)

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *spotigraph.GuildUpdateReq) (*spotigraph.SingleGuildRes, error) {
	res := &spotigraph.SingleGuildRes{}

	// Prepare expected results

	var entity models.Guild

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
		// Guild must exists
		constraints.GuildMustExists(s.guilds, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	updated := false

	if req.Name != nil {
		if err := constraints.Validate(ctx,
			// Check acceptable name value
			constraints.MustBeAName(req.Name.Value),
			// Is already used ?
			constraints.GuildNameMustBeUnique(s.guilds, req.Name.Value),
		); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusConflict,
				Message: err.Error(),
			}
			return res, err
		}
		entity.Name = req.Name.Value
		updated = true
	}

	// Skip operation when no updates
	if updated {
		// Create account in database
		if err := s.guilds.Update(ctx, &entity); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to update Guild object",
			}
			return res, err
		}
	}

	// Prepare response
	res.Entity = spotigraph.FromGuild(&entity)

	// Return expected result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *spotigraph.GuildGetReq) (*spotigraph.EmptyRes, error) {
	res := &spotigraph.EmptyRes{}

	// Prepare expected results

	var entity models.Guild

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
		// Guild must exists
		constraints.GuildMustExists(s.guilds, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	if err := s.guilds.Delete(ctx, req.Id); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to delete Guild object",
		}
		return res, err
	}

	// Return expected result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *spotigraph.GuildSearchReq) (*spotigraph.PaginatedGuildRes, error) {
	res := &spotigraph.PaginatedGuildRes{}

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
	filter := &repositories.GuildSearchFilter{}
	if req.GuildId != nil {
		filter.GuildID = req.GuildId.Value
	}
	if req.Name != nil {
		filter.Name = req.Name.Value
	}

	// Do the search
	entities, total, err := s.guilds.Search(ctx, filter, pagination, sortParams)
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
	if err != db.ErrNoResult {
		res.Members = spotigraph.FromGuilds(entities)
	}

	// Return results
	return res, nil
}
