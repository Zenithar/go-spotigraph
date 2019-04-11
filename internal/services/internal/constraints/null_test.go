package constraints_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
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
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			ctx := context.Background()

			err := constraints.Validate(ctx, constraints.MustNotBeNil(tt.input, "Property"))
			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
			}
		})
	}
}
