package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/internal/services/pkg/person/internal/mapper"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// UpdateHandler handles UpdateRequest for entity
var UpdateHandler = func(persons repositories.Person) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &personv1.UpdateResponse{}

		// Check request type
		req, ok := r.(*personv1.UpdateRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		// Prepare expected results
		var entity models.Person

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Person must exists
			constraints.PersonMustExists(persons, req.Id, &entity),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		updated := false

		if req.FirstName != nil {
			entity.FirstName = req.FirstName.Value
			updated = true
		}
		if req.LastName != nil {
			entity.LastName = req.LastName.Value
			updated = true
		}

		// Skip operation when no updates
		if updated {
			if err := persons.Update(ctx, &entity); err != nil {
				res.Error = &systemv1.Error{
					Code:    http.StatusInternalServerError,
					Message: "Unable to update Person object",
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
