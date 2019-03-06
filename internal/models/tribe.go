package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotimap/internal/helpers"
)

// Tribe describes tribe attributes holder
type Tribe struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewTribe returns a tribe instance
func NewTribe(name string) *Tribe {
	return &Tribe{
		ID:   helpers.IDGeneratorFunc(),
		Name: name,
		Slug: slug.Make(name),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Tribe) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(c.ID, validation.Required, is.Alphanumeric),
		validation.Field(c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
		validation.Field(c.Slug, validation.Required, is.PrintableASCII, validation.Length(3, 0)),
	)
}
