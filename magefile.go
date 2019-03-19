// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default    = Build
	goFiles    = getGoFiles()
	goSrcFiles = getGoSrcFiles()
)

func Build() error {
	mg.Deps(Go.Lint)

	fmt.Println("## Building")

	return sh.Run("go", "build", "-o", "bin/spotigraph", "go.zenithar.org/spotigraph/cmd/spotigraph")
}

func CI() {
	mg.SerialDeps(Proto.Service, Proto.GRPC, Go.Format, Go.Lint, Go.Test, Build)
}

// -----------------------------------------------------------------------------

type Go mg.Namespace

var deps = []string{
	"github.com/golangci/golangci-lint/cmd/golangci-lint",
	"github.com/gotestyourself/gotestsum",
}

// Test run go test
func (Go) Test() {
	fmt.Println("## Running tests")
	sh.RunV("gotestsum", "--", "-short", "-race", "-cover", "./...")
}

// Tidy add/remove depenedencies.
func (Go) Tidy() {
	fmt.Println("## Cleaning go modules")
	sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() {
	fmt.Println("## Intalling dependencies")
	for _, dep := range deps {
		fmt.Printf(" > %s\n", dep)
		sh.RunV("go", "install", dep)
	}
}

// Format runs goimports on everything
func (Go) Format() {
	fmt.Println("## Format everything")
	args := []string{"-w"}
	args = append(args, goFiles...)
	sh.RunV("goimports", args...)
}

// Lint run linter.
func (Go) Lint() {
	mg.Deps(Go.Format)
	fmt.Println("## Lint go code")
	sh.RunV("golangci-lint", "run")
}

// -----------------------------------------------------------------------------

type Proto mg.Namespace

// Service generate service protobuf objects
func (Proto) Service() error {
	fmt.Println("## Generating service DTO")
	return sh.Run(
		"protoc",
		"-I", ".",
		"-I", "${GOPATH}/src",
		"-I", "${GOPATH}/src/github.com/gogo/protobuf/protobuf",
		"--gogo_out", "Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.",
		"--validate_out", "lang=gogo:.",
		"pkg/protocol/v1/spotigraph/spotigraph.proto",
	)
}

// GRPC generate grpc stubs
func (Proto) GRPC() error {
	fmt.Println("## Generating gRPC stubs")
	return sh.Run(
		"protoc",
		"-I", ".",
		"-I", "${GOPATH}/src",
		"-I", "${GOPATH}/src/github.com/gogo/protobuf/protobuf",
		"--gogo_out", "plugins=grpc,Mpkg/protocol/v1/spotigraph/spotigraph.proto=go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph:.",
		"pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
	)
}

// -----------------------------------------------------------------------------

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}

func getGoSrcFiles() []string {
	var goSrcFiles []string

	for _, path := range goFiles {
		if !strings.HasSuffix(path, "_test.go") {
			continue
		}

		goSrcFiles = append(goSrcFiles, path)
	}

	return goSrcFiles
}
