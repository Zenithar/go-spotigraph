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

package golang

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Generate invoke the go:generate task on given package.
func Generate(name, packageName string) func() error {
	return func() error {
		fmt.Printf(" > %s [%s]\n", name, packageName)
		return sh.RunV("go", "generate", packageName)
	}
}
