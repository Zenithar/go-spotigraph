// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace

var deps = []string{
	"github.com/izumin5210/gex/cmd/gex",
}

// Tools updates tools from package
func (Go) Tools() error {
	fmt.Println("## Intalling tools")

	for _, dep := range deps {
		fmt.Printf(" > %s\n", dep)
		sh.RunV("go", "install", dep)
	}

	return sh.RunV("gex", "--build")
}
