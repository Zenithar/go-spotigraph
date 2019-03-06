package models

import (
	"fmt"

	"go.zenithar.org/spotimap/internal/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User describes user attributes holder
type User struct {
	ID        string `json:"id"`
	Principal string `json:"prn"`
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
	return validation.ValidateStruct(&u,
		validation.Field(u.ID, validation.Required, is.Alphanumeric),
		validation.Field(u.Principal, validation.Required, is.Base64),
	)
}

// URN returns an uniform resource name for external linking
func (c *User) URN() string {
	return fmt.Sprintf("urn:spom:v1:user:%s", c.ID)
}
