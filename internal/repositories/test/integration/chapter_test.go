// +build integration

package integration

import (
	"fmt"
	"testing"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/test/specs"
)

var (
	chapterRepositories = map[string]repositories.Chapter{}
)

func TestChapterRepository(t *testing.T) {

	if !*database {
		t.Skip()
	}

	for name, repo := range chapterRepositories {
		t.Run(fmt.Sprintf("Chapter repository on %s", name), specs.Chapter(repo))
	}
}
