package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotigraph/internal/helpers"
)

// Tribe describes tribe attributes holder
type Tribe struct {
	ID    string   `json:"id" bson:"_id" rethinkdb:"id"`
	Label string   `json:"label" bson:"label" rethinkdb:"label"`
	Meta  Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`

	LeaderID string      `json:"leader_id" bson:"leader_id" rethinkdb:"leader_id"`
	SquadIDs StringArray `json:"squad_ids" bson:"squad_ids" rethinkdb:"squad_ids"`
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
func (c *Tribe) SetLeader(u *User) {
	c.LeaderID = u.ID
}

// URN returns an uniform resource label for external linking
func (c *Tribe) URN() string {
	return fmt.Sprintf("spfg:v1::tribe:%s:%s", c.ID, slug.Make(c.Label))
}
