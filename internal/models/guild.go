package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotigraph/internal/helpers"
)

// Guild describes guild attributes holder
type Guild struct {
	ID    string   `json:"id" bson:"_id" rethinkdb:"id"`
	Label string   `json:"label" bson:"label" rethinkdb:"label"`
	Meta  Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`

	LeaderID  string      `json:"leader_id" bson:"leader_id" rethinkdb:"leader_id"`
	MemberIDs StringArray `json:"member_ids" bson:"member_ids" rethinkdb:"member_ids"`
}

// NewGuild returns a guild instance
func NewGuild(label string) *Guild {
	return &Guild{
		ID:    helpers.IDGeneratorFunc(),
		Label: label,
	}
}

// ------------------------------------------------------------------

// GetGroupType returns user group type
func (c *Guild) GetGroupType() string {
	return "guild"
}

// GetGroupID returns user group type
func (c *Guild) GetGroupID() string {
	return c.ID
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Guild) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Label, validation.Required, is.PrintableASCII, validation.Length(2, 50)),
	)
}

// AddMember adds the given user as member of guild
func (c *Guild) AddMember(u *User) {
	c.MemberIDs.AddIfNotContains(u.ID)
}

// RemoveMember removes the given user as member of guild
func (c *Guild) RemoveMember(u *User) {
	c.MemberIDs.Remove(u.ID)
}

// SetLeader defines the chapter leader
func (c *Guild) SetLeader(u *User) {
	c.LeaderID = u.ID
}

// URN returns an uniform resource label for external linking
func (c *Guild) URN() string {
	return fmt.Sprintf("spfg:v1::guild:%s:%s", c.ID, slug.Make(c.Label))
}
