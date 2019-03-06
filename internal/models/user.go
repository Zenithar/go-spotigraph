package models

import (
	"go.zenithar.org/spotimap/internal/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User describes user attributes holder
type User struct {
	ID string `json:"id"`
}

// NewUser returns an user instance
func NewUser(principal string) *User {
	return &User{
		ID: helpers.IDGeneratorFunc(),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (u *User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(u.ID, validation.Required, is.Alphanumeric),
	)
}
