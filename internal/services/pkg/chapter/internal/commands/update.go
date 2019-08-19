package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/mapper"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// UpdateHandler handles UpdateRequest for entity
var UpdateHandler = func(chapters repositories.Chapter) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &chapterv1.UpdateResponse{}

		// Check request type
		req, ok := r.(*chapterv1.UpdateRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		// Prepare expected results
		var entity models.Chapter

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Chapter must exists
			constraints.ChapterMustExists(chapters, req.Id, &entity),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		updated := false

		if req.Label != nil {
			if err := constraints.Validate(ctx,
				// Check acceptable name value
				constraints.MustBeAName(req.Label.Value),
				// Is already used ?
				constraints.ChapterLabelMustBeUnique(chapters, req.Label.Value),
			); err != nil {
				res.Error = &systemv1.Error{
					Code:    http.StatusConflict,
					Message: err.Error(),
				}
				return res, errors.Newf(errors.Internal, err, "unable to check label uniqueness")
			}
			entity.Label = req.Label.Value
			updated = true
		}

		// Skip operation when no updates
		if updated {
			if err := chapters.Update(ctx, &entity); err != nil {
				res.Error = &systemv1.Error{
					Code:    http.StatusInternalServerError,
					Message: "Unable to update Chapter object",
				}
				return res, errors.Newf(errors.Internal, err, "unable to update entity")
			}
		}

		// Prepare response
		res.Entity = mapper.FromEntity(&entity)

		// Return expected result
		return res, nil
	}
}
