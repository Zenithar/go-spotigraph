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

package constraints_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/internal/constraints"
)

func TestNullConstraint(t *testing.T) {
	// Testcases
	tc := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Nil",
			input:   nil,
			wantErr: true,
		}, {
			name:    "Not nil pointer",
			input:   &struct{}{},
			wantErr: false,
		}, {
			name:    "Not nil object",
			input:   struct{}{},
			wantErr: false,
		},
	}

	// Run as subtests
	for _, tt := range tc {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)
			ctx := context.Background()

			err := constraints.Validate(ctx, constraints.MustNotBeNil(testCase.input, "Property"))
			// assert results expectations
			if testCase.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
			}
		})
	}
}
