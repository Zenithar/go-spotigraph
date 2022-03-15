// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
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

package serr

import "net/http"

// ServerError returns a compliant `server_error` error.
func ServerError() Builder {
	return &defaultErrorBuilder{
		statusCode:       http.StatusInternalServerError,
		errorCode:        "server_error",
		errorDescription: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
	}
}

// InvalidRequest returns a compliant `invalid_request` error.
func InvalidRequest() Builder {
	return &defaultErrorBuilder{
		statusCode:       http.StatusBadRequest,
		errorCode:        "invalid_request",
		errorDescription: "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.",
	}
}

// ResourceNotFound returns a compliant `resource_not_found` error.
func ResourceNotFound() Builder {
	return &defaultErrorBuilder{
		statusCode:       http.StatusNotFound,
		errorCode:        "resource_not_found",
		errorDescription: "The request could not be fulfilled because the expected resource is not found.",
	}
}

// AccessDenied returns a compliant `access_denied` error.
func AccessDenied() Builder {
	return &defaultErrorBuilder{
		statusCode:       http.StatusForbidden,
		errorCode:        "access_denied",
		errorDescription: "The request is not successfully authorized.",
	}
}
