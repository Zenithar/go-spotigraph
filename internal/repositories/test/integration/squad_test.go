// +build integration

package integration

import (
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/integration/specs"
)

var squadRepositories = map[string]repositories.Squad{}

func TestSquadRepository(t *testing.T) {
	for name, repo := range squadRepositories {
		t.Run(name, specs.Squad(repo))
	}
}
