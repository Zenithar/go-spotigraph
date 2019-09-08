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
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Gen mg.Namespace

// Generate initializers
func (Gen) Wire() {
	color.Blue("### Wiring dispatchers")
	mustGoGenerate("gRPC", "go.zenithar.org/spotigraph/cli/spotigraph/internal/dispatchers/grpc")
}

// Generate mocks for tests
func (Gen) Mocks() {
	color.Blue("### Mocks")
	mustGoGenerate("Repositories", "go.zenithar.org/spotigraph/internal/repositories")
	mustGoGenerate("Publisher", "go.zenithar.org/spotigraph/internal/reactor/internal/publisher")
}

// Generate mocks for tests
func (Gen) Decorators() {
	color.Blue("### Decorators")
	mustGoGenerate("Repositories", "go.zenithar.org/spotigraph/internal/repositories/pkg/...")
	mustGoGenerate("Services", "go.zenithar.org/spotigraph/internal/services/pkg/...")
}

// Generate initializers
func (Gen) Migrations() {
	color.Blue("### Database migrations")
	mustGoGenerate("PostgreSQL", "go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql")
}

// Generate protobuf
func (Gen) Protobuf() error {
	color.Blue("### Protobuf")
	return sh.RunV("prototool", "all", "--fix", "pkg/protocol")
}
