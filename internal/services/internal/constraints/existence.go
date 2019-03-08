package constraints

import (
	"context"
	"errors"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
)

// EntityRetrieverFunc describes function indirection for repositories
type EntityRetrieverFunc func(context.Context) (interface{}, error)

func mustExists(finder EntityRetrieverFunc) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		object, err := finder(ctx)
		if err != nil {
			return err
		}
		if object == nil {
			return errors.New("Object not found")
		}
		return nil
	}
}

// UserMustExists specification checks if given user exists
func UserMustExists(users repositories.User, id string, entity *models.User) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			return users.Get(ctx, id)
		},
	)
}

// SquadMustExists specification checks if given squad exists
func SquadMustExists(squads repositories.Squad, id string, entity *models.Squad) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			return squads.Get(ctx, id)
		},
	)
}
