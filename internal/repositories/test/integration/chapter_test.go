// +build integration

package integration

import (
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/integration/specs"
)

var chapterRepositories = map[string]repositories.Chapter{}

func TestChapterRepository(t *testing.T) {
	for name, repo := range chapterRepositories {
		t.Run(name, specs.Chapter(repo))
	}
}
