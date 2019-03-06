package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotimap/internal/helpers"
)

// Guild describes guild attributes holder
type Guild struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewGuild returns a guild instance
func NewGuild(name string) *Guild {
	return &Guild{
		ID:   helpers.IDGeneratorFunc(),
		Name: name,
		Slug: slug.Make(name),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Guild) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(c.ID, validation.Required, is.Alphanumeric),
		validation.Field(c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
		validation.Field(c.Slug, validation.Required, is.PrintableASCII, validation.Length(3, 0)),
	)
}
