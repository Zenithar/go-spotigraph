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

package models

import (
	"go.zenithar.org/pkg/types"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/helpers"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Person describes person attributes holder
type Person struct {
	ID        string         `json:"id" bson:"_id" rethinkdb:"id"`
	Principal string         `json:"principal" bson:"principal" rethinkdb:"principal"`
	Meta      types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`
	FirstName string
	LastName  string
}

// NewPerson returns an person instance
func NewPerson(principal string) *Person {
	return &Person{
		ID:        helpers.IDGeneratorFunc(),
		Principal: helpers.PrincipalHashFunc(principal),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (u *Person) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.ID, helpers.IDValidationRules...),
		validation.Field(&u.Principal, validation.Required, is.PrintableASCII),
	)
}
