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
	ID   string   `json:"id" bson:"_id" rethinkdb:"id"`
	Name string   `json:"name" bson:"name" rethinkdb:"name"`
	Meta Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`

	ProductOwnerID string      `json:"product_owner_id" bson:"product_owner_id" rethinkdb:"product_owner_id"`
	MemberIDs      StringArray `json:"member_ids" bson:"member_ids" rethinkdb:"member_ids"`
}

// NewSquad returns a squad instance
func NewSquad(name string) *Squad {
	return &Squad{
		ID:        helpers.IDGeneratorFunc(),
		Name:      name,
		MemberIDs: make([]string, 0),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Squad) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// AddMember adds the given user as member of squad
func (c *Squad) AddMember(u *User) {
	c.MemberIDs.AddIfNotContains(u.ID)
}

// RemoveMember removes the given user as member of squad
func (c *Squad) RemoveMember(u *User) {
	c.MemberIDs.Remove(u.ID)
}

// SetProductOwner defines the squad product owner
func (c *Squad) SetProductOwner(u *User) {
	c.ProductOwnerID = u.ID
}

// URN returns an uniform resource name for external linking
func (c *Squad) URN() string {
	return fmt.Sprintf("urn:spfg:v1:squad:%s:%s", c.ID, slug.Make(c.Name))
}
