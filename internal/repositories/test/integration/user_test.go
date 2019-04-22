// +build integration

package integration

import (
	"fmt"
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/specs"
)

var (
	userRepositories = map[string]repositories.User{}
)

func TestUserRepository(t *testing.T) {

	if !*database {
		t.Skip()
	}

	for name, repo := range userRepositories {
		t.Run(fmt.Sprintf("User repository on %s", name), specs.User(repo))
	}
}
