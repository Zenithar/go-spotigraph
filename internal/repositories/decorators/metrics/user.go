package metrics

import (
	"context"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"go.opencensus.io/stats"
)

// -----------------------------------------------------------------------------

var (
	mUserCreatedCount      = stats.Int64("zenithar.org/measures/repositories/user/created_count", "How many user has been created", stats.UnitDimensionless)
	mUserCreatedCountError = stats.Int64("zenithar.org/measures/repositories/user/created_error_count", "How many user has been created", stats.UnitDimensionless)
)

// -----------------------------------------------------------------------------

type metricUserRepositoryDecorator struct {
	next repositories.User
}

// UserRepository returns an user repository tracing decorator
func UserRepository(next repositories.User) repositories.User {
	return &metricUserRepositoryDecorator{
		next: next,
	}
}

// ---------------------------------------------
func (d *metricUserRepositoryDecorator) Create(ctx context.Context, entity *models.User) error {
	// Delegate to next service
	err := d.next.Create(ctx, entity)

	// Set span status
	if err != nil {
		stats.Record(ctx, mUserCreatedCount.M(1))
	} else {
		stats.Record(ctx, mUserCreatedCountError.M(1))
	}

	return err
}

func (d *metricUserRepositoryDecorator) Get(ctx context.Context, id string) (*models.User, error) {
	return d.next.Get(ctx, id)
}

func (d *metricUserRepositoryDecorator) Update(ctx context.Context, entity *models.User) error {
	return d.next.Update(ctx, entity)
}

func (d *metricUserRepositoryDecorator) Delete(ctx context.Context, id string) error {
	return d.next.Delete(ctx, id)
}
