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

package constraints

import "context"

// Builder defines matcher constraints
type Builder func(context.Context) error

// Validate service constraints
func Validate(ctx context.Context, constraints ...Builder) error {
	for _, validator := range constraints {
		if err := validator(ctx); err != nil {
			return err
		}
	}

	return nil
}
