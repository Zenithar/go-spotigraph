package main

import (
	"time"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cli/spotigraph/cmd"
)

func init() {
	time.Local = time.UTC
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.CheckErr("Unable to complete command execution", err)
	}
}
