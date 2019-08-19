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

// UserRetriever is the user contract definition for read-only operation.
type UserRetriever interface {
	Get(ctx context.Context, id string) (*models.User, error)
	Search(ctx context.Context, filter *UserSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.User, int, error)
	FindByPrincipal(ctx context.Context, principal string) (*models.User, error)
}

// User describes user repository contract
type User interface {
	UserRetriever

	Create(ctx context.Context, entity *models.User) error
	Update(ctx context.Context, entity *models.User) error
	Delete(ctx context.Context, id string) error
}

// ChapterSearchFilter represents chapter entity collection search criteria
type ChapterSearchFilter struct {
	ChapterID string
	Label     string
	Slug      string
}

//go:generate mockgen -destination test/mock/chapter.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Chapter

// ChapterRetriever is the chapter contract definition for read-only operation.
type ChapterRetriever interface {
	Get(ctx context.Context, id string) (*models.Chapter, error)
	Search(ctx context.Context, filter *ChapterSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Chapter, int, error)
	FindByLabel(ctx context.Context, name string) (*models.Chapter, error)
}

// Chapter describes chapter repository contract
type Chapter interface {
	ChapterRetriever

	Create(ctx context.Context, entity *models.Chapter) error
	Update(ctx context.Context, entity *models.Chapter) error
	Delete(ctx context.Context, id string) error
}

// GuildSearchFilter represents guild entity collection search criteria
type GuildSearchFilter struct {
	GuildID string
	Label   string
	Slug    string
}

//go:generate mockgen -destination test/mock/guild.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Guild

// GuildRetriever is the guild contract definition for read-only operation.
type GuildRetriever interface {
	Get(ctx context.Context, id string) (*models.Guild, error)
	Search(ctx context.Context, filter *GuildSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Guild, int, error)
	FindByLabel(ctx context.Context, name string) (*models.Guild, error)
}

// Guild describes guild repository contract
type Guild interface {
	GuildRetriever

	Create(ctx context.Context, entity *models.Guild) error
	Update(ctx context.Context, entity *models.Guild) error
	Delete(ctx context.Context, id string) error
}

// SquadSearchFilter represents squad entity collection search criteria
type SquadSearchFilter struct {
	SquadID string
	Label   string
	Slug    string
}

//go:generate mockgen -destination test/mock/squad.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Squad

// SquadRetriever is the squad contract definition for read-only operation.
type SquadRetriever interface {
	Get(ctx context.Context, id string) (*models.Squad, error)
	Search(ctx context.Context, filter *SquadSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Squad, int, error)
	FindByLabel(ctx context.Context, name string) (*models.Squad, error)
}

// Squad describes squad repository contract
type Squad interface {
	SquadRetriever

	Create(ctx context.Context, entity *models.Squad) error
	Update(ctx context.Context, entity *models.Squad) error
	Delete(ctx context.Context, id string) error
}

// TribeSearchFilter represents tribe entity collection search criteria
type TribeSearchFilter struct {
	TribeID string
	Label   string
	Slug    string
}

//go:generate mockgen -destination test/mock/tribe.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Tribe

// TribeRetriever is the tribe contract definition for read-only operation.
type TribeRetriever interface {
	Get(ctx context.Context, id string) (*models.Tribe, error)
	Search(ctx context.Context, filter *TribeSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Tribe, int, error)
	FindByLabel(ctx context.Context, name string) (*models.Tribe, error)
}

// Tribe describes tribe repository contract
type Tribe interface {
	TribeRetriever

	Create(ctx context.Context, entity *models.Tribe) error
	Update(ctx context.Context, entity *models.Tribe) error
	Delete(ctx context.Context, id string) error
}

//go:generate mockgen -destination test/mock/membership.gen.go -package mock go.zenithar.org/spotigraph/internal/repositories Membership

// Membership describes membership repository contract
type Membership interface {
	Join(ctx context.Context, entity *models.User, ug models.UserGroup) error
	Leave(ctx context.Context, entity *models.User, ug models.UserGroup) error
}
