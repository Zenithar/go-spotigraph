package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild/internal/mapper"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// SearchHandler handles SearchRequest for entity
var SearchHandler = func(guilds repositories.GuildRetriever) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &guildv1.SearchResponse{}

		// Check request type
		req, ok := r.(*guildv1.SearchRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
		); err != nil {
			res.Error = &systemv1.Error{
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
		if req.Label != nil {
			filter.Label = req.Label.Value
		}

		// Do the search
		entities, total, err := guilds.Search(ctx, filter, pagination, sortParams)
		if err != nil && err != db.ErrNoResult {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to process request",
			}
			return res, errors.Newf(errors.Internal, err, "unable to query database")
		}

		// Set pagination total for paging calcul
		pagination.SetTotal(uint(total))
		res.Total = uint32(pagination.Total())
		res.Count = uint32(pagination.CurrentPageCount())
		res.PerPage = uint32(pagination.PerPage)
		res.CurrentPage = uint32(pagination.Page)
		res.Members = mapper.FromCollection(entities)

		// Return results
		return res, nil
	}
}
