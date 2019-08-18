package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// DeleteHandler handles DeleteRequest for entity
var DeleteHandler = func(chapters repositories.Chapter) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &chapterv1.DeleteResponse{}

		// Check non-nil request
		if r == nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "request must not be nil",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request must not be nil")
		}

		// Check request type
		req, ok := r.(*chapterv1.DeleteRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

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

		if err := chapters.Delete(ctx, req.Id); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to delete Chapter object",
			}
			return res, errors.Newf(errors.Internal, err, "unable to delete entity")
		}

		// Return expected result
		return res, nil
	}
}
