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

// CreateHandler handles CreateRequest for entity
var CreateHandler = func(persons repositories.Person) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &personv1.CreateResponse{}

		// Check request type
		req, ok := r.(*personv1.CreateRequest)
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
			constraints.PersonPrincipalMustBeUnique(persons, req.Principal),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		// Prepare Person creation
		entity := models.NewPerson(req.Principal)

		// Create use in database
		if err := persons.Create(ctx, entity); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to create Person",
			}
			return res, errors.Newf(errors.Internal, err, "unable to create entity")
		}

		// Prepare response
		res.Entity = mapper.FromEntity(entity)

		// Return result
		return res, nil
	}
}
