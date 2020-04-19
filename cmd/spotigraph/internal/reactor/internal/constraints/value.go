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

package constraints

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/helpers"
)

// MustBeAnIdentifier returns a ID contraint validator
func MustBeAnIdentifier(value string) func(context.Context) error {
	return func(ctx context.Context) error {
		return validation.Validate(value, helpers.IDValidationRules...)
	}
}

// MustBeAName returns a ID contraint validator
func MustBeAName(value string) func(context.Context) error {
	return func(ctx context.Context) error {
		return validation.Validate(value, validation.Required, is.PrintableASCII, validation.Length(2, 50))
	}
}
