package constraints

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
)

// UserMustExists specification checks if the given user exists
func UserMustExists(users repositories.User, id string, result *models.User) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		// Retrieve users details
		entity, err := users.Get(ctx, id)
		if err != nil {
			result = nil
			return err
		}

		// Return entity
		*result = *entity

		return nil
	}
}
