package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// DeleteHandler handles DeleteRequest for entity
var DeleteHandler = func(squads repositories.Squad) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &squadv1.DeleteResponse{}

		// Check request type
		req, ok := r.(*squadv1.DeleteRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		var entity models.Squad

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Squad must exists
			constraints.SquadMustExists(squads, req.Id, &entity),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		if err := squads.Delete(ctx, req.Id); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to delete Squad object",
			}
			return res, errors.Newf(errors.Internal, err, "unable to delete entity")
		}

		// Return expected result
		return res, nil
	}
}
