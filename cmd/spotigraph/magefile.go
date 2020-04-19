// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build mage

package main

import (
	"fmt"
	"runtime"

	"go.zenithar.org/spotigraph/build/mage/golang"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
)

var Default = Build

// Build the artefact
func Build() {
	banner := figure.NewFigure("Spotigraph", "", true)
	banner.Print()

	fmt.Println("")
	color.Red("# Build Info ---------------------------------------------------------------")
	fmt.Printf("Go version : %s\n", runtime.Version())

	color.Red("# Pipeline -----------------------------------------------------------------")
	mg.SerialDeps(golang.Vendor, golang.License, Generate, golang.Lint("../../"), Test)

	color.Red("# Artefact(s) --------------------------------------------------------------")
	mg.Deps(Compile)
}

// Generate code
func Generate() {
	color.Cyan("## Generate code")
}

// Test application
func Test() {
	color.Cyan("## Tests")
	mg.Deps(golang.UnitTest("./..."))
}

// Compile artefacts
func Compile() {
	mg.Deps(
		golang.Build("spotigraph", "go.zenithar.org/spotigraph/cmd/spotigraph"),
	)
}

// Release
func Release() {
	color.Red("# Releasing ---------------------------------------------------------------")

	color.Cyan("## Cross compiling artifact")

	mg.SerialDeps(
		func() error {
			return golang.Release(
				"spotigraph",
				"go.zenithar.org/spotigraph/cmd/spotigraph",
				golang.GOOS("darwin"), golang.GOARCH("amd64"),
			)()
		},
		func() error {
			return golang.Release(
				"spotigraph",
				"go.zenithar.org/spotigraph/cmd/spotigraph",
				golang.GOOS("linux"), golang.GOARCH("amd64"),
			)()
		},
		func() error {
			return golang.Release(
				"spotigraph",
				"go.zenithar.org/spotigraph/cmd/spotigraph",
				golang.GOOS("windows"), golang.GOARCH("amd64"),
			)()
		},
	)
}
