package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotigraph/internal/helpers"
)

// Squad describes squad attributes holder
type Squad struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewSquad returns a squad instance
func NewSquad(name string) *Squad {
	return &Squad{
		ID:   helpers.IDGeneratorFunc(),
		Name: name,
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Squad) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(c.ID, validation.Required, is.Alphanumeric),
		validation.Field(c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// URN returns an uniform resource name for external linking
func (c *Squad) URN() string {
	return fmt.Sprintf("urn:spom:v1:squad:%s:%s", c.ID, slug.Make(c.Name))
}
