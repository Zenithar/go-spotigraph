package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/mapper"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// CreateHandler handles CreateRequest for entity
var CreateHandler = func(chapters repositories.Chapter) HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &chapterv1.CreateResponse{}

		// Check non-nil request
		if isNil(r) {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "request must not be nil",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request must not be nil")
		}

		// Check request type
		req, ok := r.(*chapterv1.CreateRequest)
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
			// Label must be unique
			constraints.ChapterLabelMustBeUnique(chapters, req.Label),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		// Prepare Chapter creation
		entity := models.NewChapter(req.Label)

		// Create use in database
		if err := chapters.Create(ctx, entity); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to create Chapter",
			}
			return res, errors.Newf(errors.Internal, err, "unable to create entity")
		}

		// Prepare response
		res.Entity = mapper.FromEntity(entity)

		// Return result
		return res, nil
	}
}
