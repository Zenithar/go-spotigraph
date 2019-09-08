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

// Guild describes guild attributes holder
type Guild struct {
	ID       string         `json:"id" bson:"_id" rethinkdb:"id"`
	Label    string         `json:"label" bson:"label" rethinkdb:"label"`
	Meta     types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`
	LeaderID string         `json:"leader_id" bson:"leader_id" rethinkdb:"leader_id"`
}

// NewGuild returns a guild instance
func NewGuild(label string) *Guild {
	return &Guild{
		ID:    helpers.IDGeneratorFunc(),
		Label: label,
	}
}

// ------------------------------------------------------------------

// GetGroupType returns user group type
func (c *Guild) GetGroupType() string {
	return "guild"
}

// GetGroupID returns user group type
func (c *Guild) GetGroupID() string {
	return c.ID
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Guild) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Label, validation.Required, is.PrintableASCII, validation.Length(2, 50)),
	)
}

// SetLeader defines the chapter leader
func (c *Guild) SetLeader(u *Person) {
	c.LeaderID = u.ID
}
