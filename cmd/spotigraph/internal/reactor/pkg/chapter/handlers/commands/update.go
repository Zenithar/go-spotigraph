// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"context"
	"net/http"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/internal/constraints"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/internal/publisher"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/pkg/chapter/events"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/pkg/chapter/mapper"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories"

	chapterv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/chapter/v1"
	eventsv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/events/v1"
	systemv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/system/v1"
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
				if err := broker.Publish(ctx, e); err != nil {
					res.Error = &systemv1.Error{
						Code:    http.StatusInternalServerError,
						Message: "Unable to publish event",
					}
					return res, errors.Newf(errors.Internal, err, "unable to publish event")
				}
			}
		}

		// Prepare response
		res.Entity = mapper.FromEntity(&entity)

		// Return expected result
		return res, nil
	}
}
