package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotimap/internal/helpers"
)

// Chapter describes chapter attributes holder
type Chapter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewChapter returns a chapter instance
func NewChapter(name string) *Chapter {
	return &Chapter{
		ID:   helpers.IDGeneratorFunc(),
		Name: name,
		Slug: slug.Make(name),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Chapter) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(c.ID, validation.Required, is.Alphanumeric),
		validation.Field(c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
		validation.Field(c.Slug, validation.Required, is.PrintableASCII, validation.Length(3, 0)),
	)
}

// URN returns an uniform resource name for external linking
func (c *Chapter) URN() string {
	return fmt.Sprintf("urn:spom:v1:chapter:%s:%s", c.ID, c.Slug)
}
