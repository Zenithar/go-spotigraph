// +build integration

package integration

import (
	"fmt"
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/specs"
)

var guildRepositories = map[string]repositories.Guild{}

func TestGuildRepository(t *testing.T) {
	if !*database {
		t.Skip()
	}

	for name, repo := range guildRepositories {
		t.Run(fmt.Sprintf("Guild repository on %s", name), specs.Guild(repo))
	}
}
