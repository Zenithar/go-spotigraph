package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/mapper"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// GetHandler handles CreateRequest for entity
var GetHandler = func(chapters repositories.ChapterRetriever) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &chapterv1.GetResponse{}

		// Check request type
		req, ok := r.(*chapterv1.GetRequest)
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

		// Retrieve Chapter from database
		entity, err := chapters.Get(ctx, req.Id)
		if err != nil && err != db.ErrNoResult {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to retrieve Chapter",
			}
			return res, err
		}
		if entity == nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusNotFound,
				Message: "Chapter not found",
			}
			return res, errors.Newf(errors.NotFound, nil, "entity not found")
		}

		// Prepare response
		res.Entity = mapper.FromEntity(entity)

		// Return result
		return res, nil
	}
}
