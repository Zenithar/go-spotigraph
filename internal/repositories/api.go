package repositories

import (
	"context"

	"go.zenithar.org/spotimap/internal/models"
)

// User describes user repository contract
type User interface {
	Create(ctx context.Context, entity *models.User) error
	Get(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, entity *models.User) error
	Delete(ctx context.Context, id string) error
}
