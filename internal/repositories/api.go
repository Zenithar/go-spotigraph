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

//go:generate mockgen -destination test/mock/user.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories User

// User describes user repository contract
type User interface {
	Create(ctx context.Context, entity *models.User) error
	Get(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, entity *models.User) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *UserSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.User, int, error)
	FindByPrincipal(ctx context.Context, principal string) (*models.User, error)
}

// ChapterSearchFilter represents chapter entity collection search criteria
type ChapterSearchFilter struct {
	ChapterID string
	Name      string
	Slug      string
}

//go:generate mockgen -destination test/mock/chapter.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Chapter

// Chapter describes chapter repository contract
type Chapter interface {
	Create(ctx context.Context, entity *models.Chapter) error
	Get(ctx context.Context, id string) (*models.Chapter, error)
	Update(ctx context.Context, entity *models.Chapter) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *ChapterSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Chapter, int, error)
	FindByName(ctx context.Context, name string) (*models.Chapter, error)
}

// GuildSearchFilter represents guild entity collection search criteria
type GuildSearchFilter struct {
	GuildID string
	Name    string
	Slug    string
}

//go:generate mockgen -destination test/mock/guild.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Guild

// Guild describes guild repository contract
type Guild interface {
	Create(ctx context.Context, entity *models.Guild) error
	Get(ctx context.Context, id string) (*models.Guild, error)
	Update(ctx context.Context, entity *models.Guild) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *GuildSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Guild, int, error)
	FindByName(ctx context.Context, name string) (*models.Guild, error)
}

// SquadSearchFilter represents squad entity collection search criteria
type SquadSearchFilter struct {
	SquadID string
	Name    string
	Slug    string
}

//go:generate mockgen -destination test/mock/squad.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Squad

// Squad describes squad repository contract
type Squad interface {
	Create(ctx context.Context, entity *models.Squad) error
	Get(ctx context.Context, id string) (*models.Squad, error)
	Update(ctx context.Context, entity *models.Squad) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *SquadSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Squad, int, error)
	FindByName(ctx context.Context, name string) (*models.Squad, error)
}

// TribeSearchFilter represents tribe entity collection search criteria
type TribeSearchFilter struct {
	TribeID string
	Name    string
	Slug    string
}

//go:generate mockgen -destination test/mock/tribe.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Tribe

// Tribe describes tribe repository contract
type Tribe interface {
	Create(ctx context.Context, entity *models.Tribe) error
	Get(ctx context.Context, id string) (*models.Tribe, error)
	Update(ctx context.Context, entity *models.Tribe) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filter *TribeSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Tribe, int, error)
	FindByName(ctx context.Context, name string) (*models.Tribe, error)
}
