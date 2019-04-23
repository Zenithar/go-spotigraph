// +build integration

package integration

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"go.zenithar.org/pkg/testing/containers/database"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/version"
)

var databases = flag.String("databases", "postgresql", "Repositories backend to use, splitted with a coma ','. Example: postgresql,mongodb,rethinkdb")

func init() {
	flag.Parse()

	ctx := context.Background()

	// Prepare logger
	log.Setup(ctx, &log.Options{
		Debug:     true,
		AppName:   "spotigraph-integration-tests",
		AppID:     "123456",
		Version:   version.Version,
		Revision:  version.Revision,
		SentryDSN: "",
	})

	// Initialize random seed
	rand.Seed(time.Now().UTC().Unix())

	// Set UTC for all time
	time.Local = time.UTC
}

func testMainWrapper(m *testing.M) int {
	if testing.Short() {
		fmt.Println("Skipping integration tests")
		return 0
	}

	log.Bg().Info("Initializing test DB for integration test (disable with `go test -short`)")

	ctx := context.Background()
	backends := strings.Split(strings.ToLower(*databases), ",")

	for _, back := range backends {
		switch back {
		case "postgresql":
			// Initialize postgresql
			cancel, err := postgreSQLConnection(ctx)
			if err != nil {
				log.Bg().Fatal("Unable to initialize repositories", zap.Error(err))
			}
			defer func() {
				cancel()
			}()
		default:
			log.Bg().Fatal("Unsupported backend", zap.String("backend", back))
		}
	}

	defer func() {
		database.KillAll(ctx)
	}()

	return m.Run()
}

// TestMain is the test entrypoint
func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}
