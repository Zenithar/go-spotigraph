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

package main

import (
	"math/rand"
	"time"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/cmd"
	_ "go.zenithar.org/spotigraph/cmd/spotigraph/internal/features"
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
