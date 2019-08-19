package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/pkg/types"
	"go.zenithar.org/spotigraph/internal/helpers"
)

// Chapter describes chapter attributes holder
type Chapter struct {
	ID       string         `json:"id" bson:"_id" rethinkdb:"id"`
	Label    string         `json:"label" bson:"label" rethinkdb:"label"`
	Meta     types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`
	LeaderID string         `json:"leader_id" bson:"leader_id" rethinkdb:"leader_id"`
}

// NewChapter returns a chapter instance
func NewChapter(label string) *Chapter {
	return &Chapter{
		ID:    helpers.IDGeneratorFunc(),
		Label: label,
	}
}

// ------------------------------------------------------------------

// GetGroupType returns user group type
func (c *Chapter) GetGroupType() string {
	return "chapter"
}

// GetGroupID returns user group type
func (c *Chapter) GetGroupID() string {
	return c.ID
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Chapter) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Label, validation.Required, is.PrintableASCII, validation.Length(2, 50)),
	)
}

// SetLeader defines the chapter leader
func (c *Chapter) SetLeader(u *User) {
	c.LeaderID = u.ID
}

// URN returns an uniform resource name for external linking
func (c *Chapter) URN() string {
	return fmt.Sprintf("spfg:v1::chapter:%s:%s", c.ID, slug.Make(c.Label))
}
