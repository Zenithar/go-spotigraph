package constraints

import (
	"context"
	"fmt"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/helpers"
	"go.zenithar.org/spotigraph/internal/repositories"
)

// mustBeUnique specification checks if the given name already exists
func mustBeUnique(finder EntityRetrieverFunc, attribute string) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		// Retrieve object from repository
		object, err := finder(ctx)
		if err != nil && err != db.ErrNoResult {
			return err
		}
		if !isNil(object) {
			return fmt.Errorf("%s is already used", attribute)
		}

		return nil
	}
}

// UserPrincipalMustBeUnique returns specification for user principal uniqueness
func UserPrincipalMustBeUnique(users repositories.User, principal string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return users.FindByPrincipal(ctx, helpers.PrincipalHashFunc(principal))
		}, "User principal")
}

// SquadNameMustBeUnique returns specification for squad name uniqueness
func SquadNameMustBeUnique(squads repositories.Squad, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return squads.FindByName(ctx, name)
		}, "Squad name")
}

// GuildNameMustBeUnique returns specification for chapter name uniqueness
func GuildNameMustBeUnique(guilds repositories.Guild, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return guilds.FindByName(ctx, name)
		}, "Guild name")
}

// ChapterNameMustBeUnique returns specification for chapter name uniqueness
func ChapterNameMustBeUnique(chapters repositories.Chapter, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return chapters.FindByName(ctx, name)
		}, "Chapter name")
}

// TribeNameMustBeUnique returns specification for tribe name uniqueness
func TribeNameMustBeUnique(tribes repositories.Tribe, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return tribes.FindByName(ctx, name)
		}, "Tribe name")
}
