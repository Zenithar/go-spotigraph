package constraints

import (
	"context"
	"fmt"

	"go.zenithar.org/spotigraph/internal/repositories"
)

// mustBeUnique specification checks if the given name already exists
func mustBeUnique(finder EntityRetrieverFunc, attribute string) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		// Retrieve object from repository
		object, err := finder(ctx)
		if err != nil {
			return err
		}
		if object != nil {
			return fmt.Errorf("%s already used", attribute)
		}

		return nil
	}
}

// UserPrincipalMustBeUnique returns specification for user principal uniqueness
func UserPrincipalMustBeUnique(users repositories.User, principal string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return users.FindByPrincipal(ctx, principal)
		}, "User principal")
}

// SquadNameMustBeUnique returns specification for squad name uniqueness
func SquadNameMustBeUnique(squads repositories.Squad, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return squads.FindByName(ctx, name)
		}, "Squad name")
}
