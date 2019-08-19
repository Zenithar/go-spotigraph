package models

import (
	"fmt"

	"go.zenithar.org/pkg/types"
	"go.zenithar.org/spotigraph/internal/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User describes user attributes holder
type User struct {
	ID        string         `json:"id" bson:"_id" rethinkdb:"id"`
	Principal string         `json:"principal" bson:"principal" rethinkdb:"principal"`
	Meta      types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`
}

// NewUser returns an user instance
func NewUser(principal string) *User {
	return &User{
		ID:        helpers.IDGeneratorFunc(),
		Principal: helpers.PrincipalHashFunc(principal),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.ID, helpers.IDValidationRules...),
		validation.Field(&u.Principal, validation.Required, is.PrintableASCII),
	)
}

// URN returns an uniform resource name for external linking
func (u *User) URN() string {
	return fmt.Sprintf("spfg:v1::user:%s", u.ID)
}
