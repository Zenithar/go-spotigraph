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
func UserPrincipalMustBeUnique(users repositories.UserRetriever, principal string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return users.FindByPrincipal(ctx, helpers.PrincipalHashFunc(principal))
		}, "User principal")
}

// SquadLabelMustBeUnique returns specification for squad name uniqueness
func SquadLabelMustBeUnique(squads repositories.SquadRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return squads.FindByLabel(ctx, name)
		}, "Squad name")
}

// GuildLabelMustBeUnique returns specification for chapter name uniqueness
func GuildLabelMustBeUnique(guilds repositories.GuildRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return guilds.FindByLabel(ctx, name)
		}, "Guild name")
}

// ChapterLabelMustBeUnique returns specification for chapter name uniqueness
func ChapterLabelMustBeUnique(chapters repositories.ChapterRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return chapters.FindByLabel(ctx, name)
		}, "Chapter name")
}

// TribeLabelMustBeUnique returns specification for tribe name uniqueness
func TribeLabelMustBeUnique(tribes repositories.TribeRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return tribes.FindByLabel(ctx, name)
		}, "Tribe name")
}
