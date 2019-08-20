package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// JoinHandler handles JoinRequest for entity
var JoinHandler = func(guilds repositories.Guild, persons repositories.Person, memberships repositories.Membership) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &guildv1.JoinResponse{}

		// Check request type
		req, ok := r.(*guildv1.JoinRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		var (
			guild  models.Guild
			person models.Person
		)

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Guild must exists
			constraints.GuildMustExists(guilds, req.GuildId, &guild),
			// Person must exists
			constraints.PersonMustExists(persons, req.PersonId, &person),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		// Create use in database
		if err := memberships.Join(ctx, &person, &guild); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to join Guild",
			}
			return res, errors.Newf(errors.Internal, err, "unable to join guild")
		}

		// Return result
		return res, nil
	}
}
