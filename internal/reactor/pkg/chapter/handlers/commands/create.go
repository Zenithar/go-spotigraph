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
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/reactor/internal/constraints"
	"go.zenithar.org/spotigraph/internal/reactor/internal/publisher"
	"go.zenithar.org/spotigraph/internal/reactor/pkg/chapter/events"
	"go.zenithar.org/spotigraph/internal/reactor/pkg/chapter/mapper"
	"go.zenithar.org/spotigraph/internal/repositories"

	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
	systemv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
)

// CreateHandler handles CreateRequest for entity
var CreateHandler = func(chapters repositories.Chapter, persons repositories.Person, broker publisher.Publisher) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &chapterv1.CreateResponse{}

		// Check request type
		req, ok := r.(*chapterv1.CreateRequest)
		if !ok {
			res.Error = &systemv1.Error{
				Code:    http.StatusBadRequest,
				Message: "unexpected request type",
			}
			return res, errors.Newf(errors.InvalidArgument, nil, "request has invalid type (%T)", req)
		}

		var person models.Person

		// Validate service constraints
		if err := constraints.Validate(ctx,
			// Request must be syntaxically valid
			constraints.MustBeValid(req),
			// Label must be unique
			constraints.ChapterLabelMustBeUnique(chapters, req.Label),
			// Leader must exists
			constraints.PersonMustExists(persons, req.LeaderId, &person),
		); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
			}
			return res, err
		}

		// Prepare Chapter creation
		entity := models.NewChapter(req.Label)

		// Assign leader
		entity.SetLeader(&person)

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

		// Publish event
		if err := broker.Publish(ctx, events.ChapterCreated(res.Entity.Urn, res.Entity.Label, res.Entity.LeaderId)); err != nil {
			res.Error = &systemv1.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to publish event",
			}
			return res, errors.Newf(errors.Internal, err, "unable to publish event")
		}

		// Return result
		return res, nil
	}
}
