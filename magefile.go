// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default    = Build
	goFiles    = getGoFiles()
	goSrcFiles = getGoSrcFiles()
)

func init() {
	time.Local = time.UTC
}

func Build() {
	fmt.Println("# Core packages")
	mg.SerialDeps(Proto.Service, Proto.GRPC, Go.Format, Go.Lint, Go.Test)

	fmt.Println("")
	fmt.Println("# Artifacts")
	mg.Deps(Bin.Spotigraph)
}

// -----------------------------------------------------------------------------

type CI mg.Namespace

// Validate circleci configuration file (circleci/config.yml).
func (CI) Validate() error {
	return sh.RunV("circleci-cli", "config", "validate")
}

// execute circleci job build on local.
func (ci CI) Build() error {
	return ci.localExecute("build")
}

func (ci CI) localExecute(job string) error {
	return sh.RunV("circleci-cli", "local", "execute", "--job", job)
}

// -----------------------------------------------------------------------------

type Go mg.Namespace

var deps = []string{
	"github.com/izumin5210/gex/cmd/gex",
}

// Generate go code
func (Go) Generate() {
	fmt.Println("## Generate code")
	sh.RunV("go", "generate", "./...")
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
	sh.RunV("go", "mod", "vendor")

	for _, dep := range deps {
		fmt.Printf(" > %s\n", dep)
		sh.RunV("go", "install", dep)
	}

	sh.RunV("gex", "--build")
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
		"--cobra_out", "Mpkg/protocol/v1/spotigraph/spotigraph.proto=go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph:.",
		"pkg/grpc/v1/spotigraph/pb/spotigraph.proto",
	)
}

// -----------------------------------------------------------------------------

type Bin mg.Namespace

func (Bin) Spotigraph() error {
	return goBuild("go.zenithar.org/spotigraph/cmd/spotigraph", "spotigraph")
}

func goBuild(packageName, out string) error {
	fmt.Printf(" > Building %s [%s]\n", out, packageName)

	varsSetByLinker := map[string]string{
		"go.zenithar.org/spotigraph/internal/version.Version":   tag(),
		"go.zenithar.org/spotigraph/internal/version.Revision":  hash(),
		"go.zenithar.org/spotigraph/internal/version.Branch":    branch(),
		"go.zenithar.org/spotigraph/internal/version.BuildUser": "jenkins",
		"go.zenithar.org/spotigraph/internal/version.BuildDate": time.Now().Format(time.RFC3339),
		"go.zenithar.org/spotigraph/internal/version.GoVersion": runtime.Version(),
	}
	var linkerArgs string
	for name, value := range varsSetByLinker {
		linkerArgs += fmt.Sprintf(" -X %s=%s", name, value)
	}
	linkerArgs = fmt.Sprintf("-s -w %s", linkerArgs)

	return sh.Run("go", "build", "-tags", "netgo", "-ldflags", linkerArgs, "-o", fmt.Sprintf("bin/%s", out), packageName)
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

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}

// branch returns the git branch for current repo
func branch() string {
	hash, _ := sh.Output("git", "rev-parse", "--abbrev-ref", "HEAD")
	return hash
}