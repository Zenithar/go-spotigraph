// +build integration

package integration

import (
	"flag"
	"os"
	"testing"
)

var (
	database = flag.Bool("database", false, "run database integration tests")
)

func setupDatabase() {}

func teardownDatabase() {}

func TestMain(m *testing.M) {
	flag.Parse()

	if *database {
		setupDatabase()
	}

	result := m.Run()

	if *database {
		teardownDatabase()
	}

	os.Exit(result)
}
