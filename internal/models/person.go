package models

import (
	"go.zenithar.org/pkg/types"
	"go.zenithar.org/spotigraph/internal/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Person describes person attributes holder
type Person struct {
	ID        string         `json:"id" bson:"_id" rethinkdb:"id"`
	Principal string         `json:"principal" bson:"principal" rethinkdb:"principal"`
	Meta      types.Metadata `json:"meta" bson:"meta" rethinkdb:"meta"`
	FirstName string
	LastName  string
}

// NewPerson returns an person instance
func NewPerson(principal string) *Person {
	return &Person{
		ID:        helpers.IDGeneratorFunc(),
		Principal: helpers.PrincipalHashFunc(principal),
	}
}

// ------------------------------------------------------------------

// Validate entity contraints
func (u *Person) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.ID, helpers.IDValidationRules...),
		validation.Field(&u.Principal, validation.Required, is.PrintableASCII),
	)
}
