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
	var err error
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			entity, err = users.Get(ctx, id)
			return entity, err
		},
	)
}

// SquadMustExists specification checks if given squad exists
func SquadMustExists(squads repositories.Squad, id string, entity *models.Squad) func(ctx context.Context) error {
	var err error
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			entity, err = squads.Get(ctx, id)
			return entity, err
		},
	)
}

// ChapterMustExists specification checks if given chapter exists
func ChapterMustExists(chapters repositories.Chapter, id string, entity *models.Chapter) func(ctx context.Context) error {
	var err error
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			entity, err = chapters.Get(ctx, id)
			return entity, err
		},
	)
}

// TribeMustExists specification checks if given tribe exists
func TribeMustExists(tribes repositories.Tribe, id string, entity *models.Tribe) func(ctx context.Context) error {
	var err error
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			entity, err = tribes.Get(ctx, id)
			return entity, err
		},
	)
}
