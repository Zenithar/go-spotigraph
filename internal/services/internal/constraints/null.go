package constraints

import (
	"context"
	"errors"
)

// MustNotBeNil specification checks that given object is not nil
func MustNotBeNil(object interface{}, message string) func(context.Context) error {
	return func(ctx context.Context) error {
		if object == nil {
			return errors.New(message)
		}

		// Return no error
		return nil
	}
}
