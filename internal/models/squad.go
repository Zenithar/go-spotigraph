package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

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
func (c *Squad) SetProductOwner(u *User) {
	c.ProductOwnerID = u.ID
}

// URN returns an uniform resource label for external linking
func (c *Squad) URN() string {
	return fmt.Sprintf("spfg:v1::squad:%s:%s", c.ID, slug.Make(c.Label))
}
