package constraints

import (
	"context"
	"errors"
	"reflect"
)

func isNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

// MustNotBeNil specification checks that given object is not nil
func MustNotBeNil(object interface{}, message string) func(context.Context) error {
	return func(ctx context.Context) error {
		if isNil(object) {
			return errors.New(message)
		}

		// Return no error
		return nil
	}
}
