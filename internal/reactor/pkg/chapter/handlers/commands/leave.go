package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/reactor/internal/constraints"
	"go.zenithar.org/spotigraph/internal/repositories"

	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// LeaveHandler handles LeaveRequest for entity
var LeaveHandler = func(chapters repositories.ChapterRetriever, persons repositories.PersonRetriever, memberships repositories.Membership) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &chapterv1.LeaveResponse{}

		// Check request type
		req, ok := r.(*chapterv1.LeaveRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		var (
			chapter models.Chapter
			person  models.Person
		)

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Chapter must exists
			constraints.ChapterMustExists(chapters, req.ChapterId, &chapter),
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
		if err := memberships.Leave(ctx, &person, &chapter); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to leave Chapter",
			}
			return res, errors.Newf(errors.Internal, err, "unable to leave chapter")
		}

		// Return result
		return res, nil
	}
}
