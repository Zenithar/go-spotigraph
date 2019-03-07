package constraints

import (
	"context"
)

// Builder defines matcher constraints
type Builder func(context.Context) error

// Validate service constraints
func Validate(ctx context.Context, constraints ...Builder) error {
	var err error

	for _, validator := range constraints {
		if err = validator(ctx); err != nil {
			break
		}
	}

	return err
}
