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
package router

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	tenantv1 "zntr.io/spotigraph/api/gen/go/spotigraph/tenant/v1"
	"zntr.io/spotigraph/pkg/serr"
	"zntr.io/spotigraph/pkg/web/request"
	"zntr.io/spotigraph/pkg/web/respond"
)

type ServiceHandler[REQ any, RES serr.ServiceError] func(context.Context, REQ) (RES, error)

func CreateTenant(logger *zerolog.Logger, h ServiceHandler[*tenantv1.CreateRequest, *tenantv1.CreateResponse]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate method
		if r.Method != http.MethodPost {
			respond.WithJSON(w,
				http.StatusMethodNotAllowed,
				serr.InvalidRequest().Description("This endpoint only support POST method.").Build(),
			)
			return
		}

		// Decode body
		var req tenantv1.CreateRequest
		if err := request.DecodeJSON(w, r, &req); err != nil {
			logger.Error().Err(err)
			respond.WithJSON(w,
				http.StatusBadRequest,
				serr.InvalidRequest().Build(),
			)
			return
		}

		// Delegate to handler
		res, err := h(r.Context(), &req)
		switch {
		case res == nil:
			respond.WithJSON(w, http.StatusInternalServerError, serr.ServerError().Description("The handler returned an unexpected response."))
			return
		case res.GetError() != nil:
			respond.WithJSON(w, int(res.Error.StatusCode), res.Error)
			fallthrough
		case err != nil:
			logger.Error().Err(err).Msg("error occured during handler execution")
			return
		default:
		}

		// Return response
		respond.WithJSON(w, http.StatusCreated, res.Tenant)
	}
}
