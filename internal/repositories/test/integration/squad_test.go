// +build integration

package integration

import (
	"fmt"
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/specs"
)

var squadRepositories = map[string]repositories.Squad{}

func TestSquadRepository(t *testing.T) {
	if !*database {
		t.Skip()
	}

	for name, repo := range squadRepositories {
		t.Run(fmt.Sprintf("Squad repository on %s", name), specs.Squad(repo))
	}
}
