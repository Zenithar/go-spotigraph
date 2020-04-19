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

package helpers

import (
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// IDGeneratedLength defines the length of the id string
const IDGeneratedLength = 32

// IDGeneratorFunc returns a randomly generated string useable as identifier
var IDGeneratorFunc = func() string {
	return uniuri.NewLen(IDGeneratedLength)
}

// IDValidationRules describes identifier contract for syntaxic validation
var IDValidationRules = []validation.Rule{
	validation.Required,
	validation.Length(IDGeneratedLength, IDGeneratedLength),
	is.Alphanumeric,
}
