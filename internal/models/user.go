package models

import (
	"go.zenithar.org/spotimap/internal/helpers"
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
