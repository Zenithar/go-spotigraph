package tracing

import (
	"context"

	"go.zenithar.org/spotimap/internal/models"
	"go.zenithar.org/spotimap/internal/repositories"

	"go.opencensus.io/trace"
)

type tracingUserRepositoryDecorator struct {
	next repositories.User
}

// UserRepositoryTracing returns an user repository tracing decorator
func UserRepositoryTracing(next repositories.User) repositories.User {
	return &tracingUserRepositoryDecorator{
		next: next,
	}
}

// ---------------------------------------------
func (d *tracingUserRepositoryDecorator) Create(ctx context.Context, entity *models.User) error {
	ctx, span := trace.StartSpan(ctx, "user.Create")
	defer span.End()

	return d.next.Create(ctx, entity)
}

func (d *tracingUserRepositoryDecorator) Get(ctx context.Context, id string) (*models.User, error) {
	ctx, span := trace.StartSpan(ctx, "user.Get")
	defer span.End()

	return d.next.Get(ctx, id)
}

func (d *tracingUserRepositoryDecorator) Update(ctx context.Context, entity *models.User) error {
	ctx, span := trace.StartSpan(ctx, "user.Update")
	defer span.End()

	return d.next.Update(ctx, entity)
}

func (d *tracingUserRepositoryDecorator) Delete(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "user.Delete")
	defer span.End()

	return d.next.Delete(ctx, id)
}
