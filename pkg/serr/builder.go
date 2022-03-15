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

import (
	"errors"
	"fmt"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	systemv1 "zntr.io/spotigraph/api/gen/go/spotigraph/system/v1"
	"zntr.io/spotigraph/pkg/types"
)

// -----------------------------------------------------------------------------

// ErrorBuilder describes error builder contract.
type ErrorBuilder interface {
	State(value string) ErrorBuilder
	Description(value string) ErrorBuilder
	Build() systemv1.Error
}

type defaultErrorBuilder struct {
	statusCode       int32
	errorCode        string
	errorDescription string
	errorURI         string
	state            string
	internalError    error
	fields           []string
}

func (eb *defaultErrorBuilder) State(value string) Builder {
	eb.state = value
	return eb
}

func (eb *defaultErrorBuilder) StatusCode(value int32) Builder {
	eb.statusCode = value
	return eb
}

func (eb *defaultErrorBuilder) ErrorCode(value string) Builder {
	eb.errorCode = value
	return eb
}

func (eb *defaultErrorBuilder) Description(value string) Builder {
	eb.errorDescription = value
	return eb
}

func (eb *defaultErrorBuilder) Descriptionf(format string, args ...interface{}) Builder {
	eb.errorDescription = fmt.Sprintf(format, args...)
	return eb
}

func (eb *defaultErrorBuilder) InternalErr(err error) Builder {
	eb.internalError = err
	eb.fields = fieldsFromError(err)
	return eb
}

func (eb *defaultErrorBuilder) ErrorURI(value string) Builder {
	eb.errorURI = value
	return eb
}

func (eb *defaultErrorBuilder) Fields(values ...string) Builder {
	fields := types.StringArray(eb.fields)
	fields.AddIfNotContains(values...)
	eb.fields = fields
	return eb
}

func (eb *defaultErrorBuilder) Build() *systemv1.Error {
	// Create error object
	err := &systemv1.Error{
		StatusCode:       eb.statusCode,
		ErrorCode:        eb.errorCode,
		ErrorDescription: eb.errorDescription,
	}
	if len(eb.fields) > 0 {
		err.Fields = eb.fields
	}

	// Return error instance
	return err
}

// -----------------------------------------------------------------------------

func fieldsFromError(err error) []string {
	var errs validation.Errors
	// Check if it's a validation error
	if errors.As(err, &errs) {
		var fields types.StringArray
		for k := range errs {
			fields.AddIfNotContains(k)
		}
		sort.Strings(fields)
		return fields
	}

	// No fields
	return nil
}
