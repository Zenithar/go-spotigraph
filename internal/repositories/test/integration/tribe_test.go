// +build integration

package integration

import (
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/integration/specs"
)

var tribeRepositories = map[string]repositories.Tribe{}

func TestTribeRepository(t *testing.T) {
	for name, repo := range tribeRepositories {
		t.Run(name, specs.Tribe(repo))
	}
}
