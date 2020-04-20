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
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"go.zenithar.org/pkg/types"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/helpers"
)

// Tribe describes tribe attributes holder
type Tribe struct {
	ID    string         `json:"id" bson:"_id" rethinkdb:"id"`
	Label string         `json:"label" bson:"label" rethinkdb:"label"`
	Meta  types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`

	LeaderID string            `json:"leader_id" bson:"leader_id" rethinkdb:"leader_id"`
	SquadIDs types.StringArray `json:"squad_ids" bson:"squad_ids" rethinkdb:"squad_ids"`
}

// NewTribe returns a tribe instance
func NewTribe(label string) *Tribe {
	return &Tribe{
		ID:    helpers.IDGeneratorFunc(),
		Label: label,
	}
}

// ------------------------------------------------------------------

// GetGroupType returns user group type
func (c *Tribe) GetGroupType() string {
	return "tribe"
}

// GetGroupID returns user group type
func (c *Tribe) GetGroupID() string {
	return c.ID
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Tribe) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Label, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// AddSquad adds the given squad as member of tribe
func (c *Tribe) AddSquad(s *Squad) {
	c.SquadIDs.AddIfNotContains(s.ID)
}

// RemoveSquad removes the given squad as member of tribe
func (c *Tribe) RemoveSquad(s *Squad) {
	c.SquadIDs.Remove(s.ID)
}

// SetLeader defines the chapter leader
func (c *Tribe) SetLeader(u *Person) {
	c.LeaderID = u.ID
}
