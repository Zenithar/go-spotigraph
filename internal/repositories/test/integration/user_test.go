// +build integration

package integration

import (
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/integration/specs"
)

var userRepositories = map[string]repositories.User{}

func TestUserRepository(t *testing.T) {
	for name, repo := range userRepositories {
		t.Run(name, specs.User(repo))
	}
}
