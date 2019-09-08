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
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"go.zenithar.org/pkg/types"
	"go.zenithar.org/spotigraph/internal/helpers"
)

// Squad describes squad attributes holder
type Squad struct {
	ID             string         `json:"id" bson:"_id" rethinkdb:"id"`
	Label          string         `json:"label" bson:"label" rethinkdb:"label"`
	Meta           types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`
	ProductOwnerID string         `json:"product_owner_id" bson:"product_owner_id" rethinkdb:"product_owner_id"`
}

// NewSquad returns a squad instance
func NewSquad(label string) *Squad {
	return &Squad{
		ID:    helpers.IDGeneratorFunc(),
		Label: label,
	}
}

// ------------------------------------------------------------------

// GetGroupType returns user group type
func (c *Squad) GetGroupType() string {
	return "squad"
}

// GetGroupID returns user group type
func (c *Squad) GetGroupID() string {
	return c.ID
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Squad) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Label, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// SetProductOwner defines the squad product owner
func (c *Squad) SetProductOwner(u *Person) {
	c.ProductOwnerID = u.ID
}
