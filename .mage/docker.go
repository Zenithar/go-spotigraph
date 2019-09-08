// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zmage

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Build docker image.
func (Docker) Build() error {
	color.Red("# Docker -------------------------------------------------------------------")
	fmt.Printf("BUILD_DATE : %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("VERSION : %s\n", tag())
	fmt.Printf("VCS_REF : %s\n", hash())

	fmt.Printf(" > Production image\n")
	return sh.RunV("docker", "build",
		"-f", "deployment/docker/Dockerfile",
		"--build-arg", fmt.Sprintf("BUILD_DATE=%s", time.Now().Format(time.RFC3339)),
		"--build-arg", fmt.Sprintf("VERSION=%s", tag()),
		"--build-arg", fmt.Sprintf("VCS_REF=%s", hash()),
		"-t", "spotigraph:latest",
		".")
}
