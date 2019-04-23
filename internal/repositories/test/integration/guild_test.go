// +build integration

package integration

import (
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/integration/specs"
)

var guildRepositories = map[string]repositories.Guild{}

func TestGuildRepository(t *testing.T) {
	for name, repo := range guildRepositories {
		t.Run(name, specs.Guild(repo))
	}
}
