package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/reactor/internal/constraints"
	"go.zenithar.org/spotigraph/internal/reactor/internal/publisher"
	"go.zenithar.org/spotigraph/internal/reactor/pkg/chapter/events"
	"go.zenithar.org/spotigraph/internal/reactor/pkg/chapter/mapper"
	"go.zenithar.org/spotigraph/internal/repositories"

	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	eventsv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/events/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// UpdateHandler handles UpdateRequest for entity
var UpdateHandler = func(chapters repositories.Chapter, persons repositories.PersonRetriever, broker publisher.Publisher) reactor.HandlerFunc {
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

		// Translate as DTO
		dto := mapper.FromEntity(&entity)

		// Prepare event list to publish
		eventList := []*eventsv1.Event{}

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

			// Add event to list
			eventList = append(eventList, events.ChapterLabelUpdated(dto.Urn, entity.Label, req.Label.Value))

			// Update attribute
			entity.Label = req.Label.Value
		}

		if req.LeaderId != nil {
			var person models.Person

			if err := constraints.Validate(ctx,
				// Check acceptable id value
				constraints.MustBeAnIdentifier(req.LeaderId.Value),
				// Person exists ?
				constraints.PersonMustExists(persons, req.LeaderId.Value, &person),
			); err != nil {
				res.Error = &systemv1.Error{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				}
				return res, errors.Newf(errors.InvalidArgument, nil, "person not found")
			}

			// Add event to list
			eventList = append(eventList, events.ChapterLeaderUpdated(dto.Urn, entity.LeaderID, req.LeaderId.Value))

			// Update attribute
			entity.SetLeader(&person)
		}

		// Skip operation when no updates
		if len(eventList) > 0 {
			if err := chapters.Update(ctx, &entity); err != nil {
				res.Error = &systemv1.Error{
					Code:    http.StatusInternalServerError,
					Message: "Unable to update Chapter object",
				}
				return res, errors.Newf(errors.Internal, err, "unable to update entity")
			}

			// Publish all events
			for _, e := range eventList {
				broker.Publish(ctx, e)
			}
		}

		// Prepare response
		res.Entity = mapper.FromEntity(&entity)

		// Return expected result
		return res, nil
	}
}
