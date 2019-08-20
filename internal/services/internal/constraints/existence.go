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
		if isNil(object) {
			return errors.New("object not found")
		}
		return nil
	}
}

// PersonMustExists specification checks if given user exists
func PersonMustExists(users repositories.PersonRetriever, id string, entity *models.Person) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := users.Get(ctx, id)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}

// SquadMustExists specification checks if given squad exists
func SquadMustExists(squads repositories.SquadRetriever, id string, entity *models.Squad) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := squads.Get(ctx, id)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}

// GuildMustExists specification checks if given guild exists
func GuildMustExists(guilds repositories.GuildRetriever, id string, entity *models.Guild) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := guilds.Get(ctx, id)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}

// ChapterMustExists specification checks if given chapter exists
func ChapterMustExists(chapters repositories.ChapterRetriever, id string, entity *models.Chapter) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := chapters.Get(ctx, id)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}

// TribeMustExists specification checks if given tribe exists
func TribeMustExists(tribes repositories.TribeRetriever, id string, entity *models.Tribe) func(ctx context.Context) error {
	return mustExists(
		func(ctx context.Context) (interface{}, error) {
			object, err := tribes.Get(ctx, id)
			if object != nil {
				*entity = *object
			}
			return entity, err
		},
	)
}
