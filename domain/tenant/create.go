// Licensed to Thibault Normand under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Thibault Normand licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package tenant

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/gosimple/slug"
	tenantv1 "zntr.io/spotigraph/api/gen/go/spotigraph/tenant/v1"
	"zntr.io/spotigraph/infrastructure/publisher"
	"zntr.io/spotigraph/infrastructure/repositories"
	"zntr.io/spotigraph/pkg/serr"
	"zntr.io/spotigraph/pkg/types"
)

var slugExpression = regexp.MustCompile("^[a-z0-9]+(?:-[a-z0-9]+)*$")

func CreateHandler(tenants Repository, eventBroker publisher.Publisher) func(context.Context, *tenantv1.CreateRequest) (*tenantv1.CreateResponse, error) {
	return func(ctx context.Context, req *tenantv1.CreateRequest) (*tenantv1.CreateResponse, error) {
		res := &tenantv1.CreateResponse{}

		// Check arguments
		if req == nil {
			res.Error = serr.ServerError().Build()
			return res, errors.New("unable to process nil request")
		}

		// Check label uniqueness
		saved, err := tenants.GetByLabel(ctx, req.Label)
		switch {
		case errors.Is(err, repositories.ErrNoResult) && types.IsNil(saved):
			// No label match
		case err != nil:
			res.Error = serr.ServerError().Build()
			return res, fmt.Errorf("unable check label uniqueness: %w", err)
		case !types.IsNil(saved):
			res.Error = serr.InvalidRequest().Fields("label").Descriptionf("The label %q is already used.", req.Label).Build()
			return res, nil
		}

		// Validate slug
		if req.Slug == nil {
			req.Slug = types.StringRef(slug.Make(req.Label))
		} else if !slug.IsSlug(*req.Slug) {
			res.Error = serr.InvalidRequest().Fields("slug").Descriptionf("The slug %q is not a valid slug format.", *req.Slug).Build()
			return res, nil
		}

		// Check slug uniqueness
		saved, err = tenants.GetBySlug(ctx, *req.Slug)
		switch {
		case errors.Is(err, repositories.ErrNoResult) && types.IsNil(saved):
			// No slug match
		case err != nil:
			res.Error = serr.ServerError().Build()
			return res, fmt.Errorf("unable check slug uniqueness: %w", err)
		case !types.IsNil(saved):
			res.Error = serr.InvalidRequest().Fields("slug").Descriptionf("The slug %q is already used.", *req.Slug).Build()
			return res, nil
		}

		// Stop execution if validation only
		if req.ValidateOnly {
			return res, nil
		}

		// Create a new domain identifier
		newTenantID, err := tenants.NextID(ctx)
		if err != nil {
			res.Error = serr.ServerError().Build()
			return res, fmt.Errorf("unable generate next tenant identifier: %w", err)
		}

		// Create a new domain object
		do := New(
			WithID(newTenantID),
			WithLabel(req.Label),
			WithSlug(*req.Slug),
		)

		// Delegate to persistence
		if err := tenants.Save(ctx, do); err != nil {
			res.Error = serr.ServerError().Build()
			return res, fmt.Errorf("unable to save tenant: %w", err)
		}

		// Prepare response
		res.Tenant = &tenantv1.Tenant{
			Id:    string(newTenantID),
			Label: do.GetLabel(),
			Slug:  do.GetSlug(),
		}

		// Generate event
		if err := eventBroker.Publish(ctx, EventCreated(do.GetID(), do.GetLabel(), do.GetSlug())); err != nil {
			res.Error = serr.ServerError().Build()
			return res, fmt.Errorf("unable to publish tenant created event: %w", err)
		}

		// No error
		return res, nil
	}
}
