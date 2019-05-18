package main

import (
	"math/rand"
	"time"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cli/spotigraph/internal/cmd"
	_ "go.zenithar.org/spotigraph/cli/spotigraph/internal/features"
)

// -----------------------------------------------------------------------------

func init() {
	// Set time locale
	time.Local = time.UTC

	// Initialize random seed
	rand.Seed(time.Now().UTC().Unix())
}

// -----------------------------------------------------------------------------

func main() {
	if err := cmd.Execute(); err != nil {
		log.CheckErr("Unable to complete command execution", err)
	}
}
