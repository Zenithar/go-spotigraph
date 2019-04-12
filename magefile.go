// +build mage

package main

import (
	"fmt"
	"os"
	"path"
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

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// Calculate file paths
var toolsBinDir = normalizePath(path.Join(curDir, "bin"))

func init() {
	time.Local = time.UTC

	// Add local bin in PATH
	err := os.Setenv("PATH", fmt.Sprintf("%s:%s", toolsBinDir, os.Getenv("PATH")))
	if err != nil {
		panic(err)
	}
}

func Build() {
	fmt.Println("# Core packages")
	mg.SerialDeps(Prototool.Generate, Go.Generate, Go.Format, Go.Import, Go.Lint, Go.Test)

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
func (Go) Generate() error {
	fmt.Println("## Generate code")
	return sh.RunV("go", "generate", "./...")
}

// Test run go test
func (Go) Test() error {
	fmt.Println("## Running tests")
	return sh.RunV("gotestsum", "--junitfile", "unit-tests.xml", "--", "-short", "-race", "-cover", "./...")
}

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	fmt.Println("## Cleaning go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	fmt.Println("## Intalling dependencies")
	sh.RunV("go", "mod", "vendor")

	for _, dep := range deps {
		fmt.Printf(" > %s\n", dep)
		sh.RunV("go", "install", dep)
	}

	return sh.RunV("gex", "--build")
}

// Format runs gofmt on everything
func (Go) Format() error {
	fmt.Println("## Format everything")
	args := []string{"-s", "-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

// Import runs goimports on everything
func (Go) Import() error {
	fmt.Println("## Process imports")
	args := []string{"-w"}
	args = append(args, goSrcFiles...)
	return sh.RunV("goreturns", args...)
}

// Lint run linter.
func (Go) Lint() error {
	mg.Deps(Go.Format)
	fmt.Println("## Lint go code")
	return sh.RunV("golangci-lint", "run")
}

// -----------------------------------------------------------------------------

type Prototool mg.Namespace

func (Prototool) Lint() error {
	fmt.Println("## Protobuf lint")
	return sh.RunV("prototool", "generate")
}

func (Prototool) Generate() error {
	mg.Deps(Prototool.Lint)
	fmt.Println("## Protobuf code generation")
	return sh.RunV("prototool", "generate")
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
		"go.zenithar.org/spotigraph/internal/version.BuildUser": os.Getenv("USER"),
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

func mustStr(r string, err error) string {
	if err != nil {
		panic(err)
	}
	return r
}

// normalizePath turns a path into an absolute path and removes symlinks
func normalizePath(name string) string {
	absPath := mustStr(filepath.Abs(name))
	return absPath
}
