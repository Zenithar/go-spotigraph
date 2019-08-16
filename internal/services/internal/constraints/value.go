package constraints

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.zenithar.org/spotigraph/internal/helpers"
)

// MustBeAnIdentifier returns a ID contraint validator
func MustBeAnIdentifier(value string) func(context.Context) error {
	return func(ctx context.Context) error {
		return validation.Validate(value, helpers.IDValidationRules...)
	}
}

// MustBeAName returns a ID contraint validator
func MustBeAName(value string) func(context.Context) error {
	return func(ctx context.Context) error {
		return validation.Validate(value, validation.Required, is.PrintableASCII, validation.Length(2, 50))
	}
}
