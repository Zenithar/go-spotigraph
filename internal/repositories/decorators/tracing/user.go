package tracing

import (
	"context"

	"go.zenithar.org/spotimap/internal/models"
	"go.zenithar.org/spotimap/internal/repositories"

	"go.opencensus.io/trace"
)

type userRepositoryDecorator struct {
	next repositories.User
}

// UserRepository returns an user repository tracing decorator
func UserRepository(next repositories.User) repositories.User {
	return &userRepositoryDecorator{
		next: next,
	}
}

// ---------------------------------------------
func (d *userRepositoryDecorator) Create(ctx context.Context, entity *models.User) error {
	// Start span
	ctx, span := trace.StartSpan(ctx, "user.Create")

	// Delegate to next service
	err := d.next.Create(ctx, entity)

	// Set span status
	spanStatus(span, err)

	// End the span
	span.End()

	// Pass result to next decorator
	return err
}

func (d *userRepositoryDecorator) Get(ctx context.Context, id string) (*models.User, error) {
	ctx, span := trace.StartSpan(ctx, "user.Get")

	// Delegate to next service
	entity, err := d.next.Get(ctx, id)

	// Set span status
	spanStatus(span, err)

	// End the span
	span.End()

	// Pass result to next decorator
	return entity, err
}

func (d *userRepositoryDecorator) Update(ctx context.Context, entity *models.User) error {
	ctx, span := trace.StartSpan(ctx, "user.Update")

	// Delegate to next service
	err := d.next.Update(ctx, entity)

	// Set span status
	spanStatus(span, err)

	// End the span
	span.End()

	// Pass result to next decorator
	return err
}

func (d *userRepositoryDecorator) Delete(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "user.Delete")

	// Delegate to next service
	err := d.next.Delete(ctx, id)

	// Set span status
	spanStatus(span, err)

	// End the span
	span.End()

	// Pass result to next decorator
	return err
}
