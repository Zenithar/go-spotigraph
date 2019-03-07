package constraints

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
)

// SquadMustExists specification checks if the given squad exists
func SquadMustExists(squads repositories.Squad, id string, result *models.Squad) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		// Retrieve users details
		entity, err := squads.Get(ctx, id)
		if err != nil {
			result = nil
			return err
		}

		// Return entity
		*result = *entity

		return nil
	}
}

// SquadNameMustBeUnique specification checks if the given name already exists
func SquadNameMustBeUnique(squads repositories.Squad, name string) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		// Check principal existence
		_, err := squads.FindByName(ctx, name)
		if err != nil {
			return err
		}

		return nil
	}
}
