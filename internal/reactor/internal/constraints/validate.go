package constraints

import "context"

// Validable interface used to defines Validation protocol
type Validable interface {
	Validate() error
}

// MustBeValid specification checks that given object is valid
func MustBeValid(validable Validable) func(context.Context) error {
	return func(ctx context.Context) error {
		// Validate request
		return validable.Validate()
	}
}
