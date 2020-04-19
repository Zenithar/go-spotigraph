//+build integration

package specs

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories"
)

// Squad returns squad repository full test scenario builder
func Squad(underTest repositories.Squad) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		g := NewGomegaWithT(t)

		// Stub context
		ctx := context.Background()

		// Prepare a new entity
		created := models.NewSquad("foo")
		g.Expect(created).ToNot(BeNil(), "Newly created entity should not be nil")

		// Create the entity using repository
		err := underTest.Create(ctx, created)
		g.Expect(err).To(BeNil(), "Error creation should be nil")

		// -------------------------------------------------------------------------------------------------------------

		// Retrieve by id from repository
		saved, err := underTest.Get(ctx, created.ID)
		g.Expect(err).To(BeNil(), "Retrieval error should be nil")
		g.Expect(saved).ToNot(BeNil(), "Saved entity should not be nil")

		// Compare objects
		g.Expect(cmp.Equal(created, saved)).To(BeTrue(), "Saved and Created should be equals")

		// Retrieve by non-existent id
		nonExistent, err := underTest.Get(ctx, "non-existent-id")
		g.Expect(err).ToNot(BeNil(), "Error should be raised on non-existent entity")
		g.Expect(err).To(Equal(db.ErrNoResult), "Error ErrNoResult should be raised")
		g.Expect(nonExistent).To(BeNil(), "Non-existent entity should be nil")

		// -------------------------------------------------------------------------------------------------------------

		// Update an entity
		saved, err = underTest.Get(ctx, created.ID)
		g.Expect(err).To(BeNil(), "Retrieval error should be nil")
		g.Expect(saved).ToNot(BeNil(), "Saved entity should not be nil")

		// Update properties

		// Update with repository
		err = underTest.Update(ctx, saved)
		g.Expect(err).To(BeNil(), "Update error should be nil")

		// Retrieve from repository to check updated properties
		updated, err := underTest.Get(ctx, created.ID)
		g.Expect(err).To(BeNil(), "Retrieval error should be nil")
		g.Expect(updated).ToNot(BeNil(), "Saved entity should not be nil")

		// Compare objects
		g.Expect(cmp.Equal(created, updated)).To(BeTrue(), "Saved and Updated should be equals")

		// -------------------------------------------------------------------------------------------------------------

		// Remove an entity
		err = underTest.Delete(ctx, created.ID)
		g.Expect(err).To(BeNil(), "Removal error should be nil")

		// Retrieve from repository to check deletion
		deleted, err := underTest.Get(ctx, created.ID)
		g.Expect(err).ToNot(BeNil(), "Deletion error should not be nil")
		g.Expect(err).To(Equal(db.ErrNoResult), "Error ErrNoResult should be raised")
		g.Expect(deleted).To(BeNil(), "Deleted entity should be nil")

		// Remove a non-existent entity
		err = underTest.Delete(ctx, "non-existent-id")
		g.Expect(err).ToNot(BeNil(), "Removal error should not be nil")
		g.Expect(err).To(Equal(db.ErrNoModification), "Error ErrNoModification should be raised")
	}
}
