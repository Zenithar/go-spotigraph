// +build integration

package integration

import (
	"fmt"
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/specs"
)

var tribeRepositories = map[string]repositories.Tribe{}

func TestTribeRepository(t *testing.T) {
	if !*database {
		t.Skip()
	}

	for name, repo := range tribeRepositories {
		t.Run(fmt.Sprintf("Tribe repository on %s", name), specs.Tribe(repo))
	}
}
