package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"

	"go.zenithar.org/spotigraph/internal/helpers"
)

// Chapter describes chapter attributes holder
type Chapter struct {
	ID   string   `json:"id" bson:"_id" rethinkdb:"id" db:"chapter_id"`
	Name string   `json:"name" bson:"name" rethinkdb:"name" db:"name"`
	Meta Metadata `json:"meta" bson:"meta" rethinkdb:"meta" db:"meta"`

	Leader  string      `json:"leader_id" bson:"leader_id" rethinkdb:"leader_id" db:"leader_id"`
	Members StringArray `json:"member_ids" bson:"member_ids" rethinkdb:"member_ids" db:"members"`
}

// NewChapter returns a chapter instance
func NewChapter(name string) *Chapter {
	return &Chapter{
		ID:   helpers.IDGeneratorFunc(),
		Name: name,
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (c *Chapter) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, helpers.IDValidationRules...),
		validation.Field(&c.Name, validation.Required, is.PrintableASCII, validation.Length(3, 50)),
	)
}

// AddMember adds the given user as member of chapter
func (c *Chapter) AddMember(u *User) {
	c.Members.AddIfNotContains(u.ID)
}

// RemoveMember removes the given user as member of chapter
func (c *Chapter) RemoveMember(u *User) {
	c.Members.Remove(u.ID)
}

// URN returns an uniform resource name for external linking
func (c *Chapter) URN() string {
	return fmt.Sprintf("urn:spfg:v1:chapter:%s:%s", c.ID, slug.Make(c.Name))
}
