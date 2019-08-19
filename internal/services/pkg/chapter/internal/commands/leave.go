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

// LeaveHandler handles LeaveRequest for entity
var LeaveHandler = func(chapters repositories.ChapterRetriever, users repositories.UserRetriever, memberships repositories.Membership) reactor.HandlerFunc {
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
			user    models.User
		)

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Chapter must exists
			constraints.ChapterMustExists(chapters, req.ChapterId, &chapter),
			// User must exists
			constraints.UserMustExists(users, req.UserId, &user),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		// Create use in database
		if err := memberships.Leave(ctx, &user, &chapter); err != nil {
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