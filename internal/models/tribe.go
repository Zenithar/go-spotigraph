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
	ID   string   `json:"id" bson:"_id" rethinkdb:"id"`
	Name string   `json:"name" bson:"name" rethinkdb:"name"`
	Meta Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`

	Squads StringArray `json:"squad_ids" bson:"squad_ids" rethinkdb:"squad_ids"`
}

// NewTribe returns a tribe instance
func NewTribe(name string) *Tribe {
	return &Tribe{
		ID:   helpers.IDGeneratorFunc(),
		Name: name,
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Tribe) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, validation.Required, is.Alphanumeric),
		validation.Field(&c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// AddSquad adds the given squad as member of tribe
func (c *Tribe) AddSquad(s *Squad) {
	c.Squads.AddIfNotContains(s.ID)
}

// RemoveSquad removes the given squad as member of tribe
func (c *Tribe) RemoveSquad(s *Squad) {
	c.Squads.Remove(s.ID)
}

// URN returns an uniform resource name for external linking
func (c *Tribe) URN() string {
	return fmt.Sprintf("urn:spfg:v1:tribe:%s:%s", c.ID, slug.Make(c.Name))
}
