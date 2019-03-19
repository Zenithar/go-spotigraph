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
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Meta Metadata `json:"meta"`
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
	return validation.ValidateStruct(&c,
		validation.Field(c.ID, validation.Required, is.Alphanumeric),
		validation.Field(c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// URN returns an uniform resource name for external linking
func (c *Tribe) URN() string {
	return fmt.Sprintf("urn:spfg:v1:tribe:%s:%s", c.ID, slug.Make(c.Name))
}
