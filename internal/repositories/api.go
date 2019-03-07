package repositories

import (
	"context"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
)

// UserSearchFilter represents user entity collection search criteria
type UserSearchFilter struct {
	UserID    string
	Principal string
}

// User describes user repository contract
type User interface {
	Create(ctx context.Context, entity *models.User) error
	Get(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, entity *models.User) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *UserSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.User, int, error)
	FindByPrincipal(ctx context.Context, principal string) (*models.User, error)
}

// Chapter describes chapter repository contract
type Chapter interface {
	Create(ctx context.Context, entity *models.Chapter) error
	Get(ctx context.Context, id string) (*models.Chapter, error)
	Update(ctx context.Context, entity *models.Chapter) error
	Delete(ctx context.Context, id string) error
}

// Guild describes guild repository contract
type Guild interface {
	Create(ctx context.Context, entity *models.Guild) error
	Get(ctx context.Context, id string) (*models.Guild, error)
	Update(ctx context.Context, entity *models.Guild) error
	Delete(ctx context.Context, id string) error
}

// Squad describes squad repository contract
type Squad interface {
	Create(ctx context.Context, entity *models.Squad) error
	Get(ctx context.Context, id string) (*models.Squad, error)
	Update(ctx context.Context, entity *models.Squad) error
	Delete(ctx context.Context, id string) error
}

// Tribe describes tribe repository contract
type Tribe interface {
	Create(ctx context.Context, entity *models.Tribe) error
	Get(ctx context.Context, id string) (*models.Tribe, error)
	Update(ctx context.Context, entity *models.Tribe) error
	Delete(ctx context.Context, id string) error
}
