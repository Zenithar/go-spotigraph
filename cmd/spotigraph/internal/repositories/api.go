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

package repositories

import (
	"context"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
)

// PersonSearchFilter represents person entity collection search criteria
type PersonSearchFilter struct {
	PersonID  string
	Principal string
}

//go:generate mockgen -destination test/mock/person.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories Person

// PersonRetriever is the person contract definition for read-only operation.
type PersonRetriever interface {
	Get(ctx context.Context, id string) (*models.Person, error)
	Search(ctx context.Context, filter *PersonSearchFilter, pagination *db.Pagination, sortParams *db.SortParameters) ([]*models.Person, int, error)
	FindByPrincipal(ctx context.Context, principal string) (*models.Person, error)
}

// Person describes person repository contract
type Person interface {
	PersonRetriever

	Create(ctx context.Context, entity *models.Person) error
	Update(ctx context.Context, entity *models.Person) error
	Delete(ctx context.Context, id string) error
}

// ChapterSearchFilter represents chapter entity collection search criteria
type ChapterSearchFilter struct {
	ChapterID string
	Label     string
	Slug      string
}

//go:generate mockgen -destination test/mock/chapter.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories Chapter

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

//go:generate mockgen -destination test/mock/guild.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories Guild

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

//go:generate mockgen -destination test/mock/squad.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories Squad

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

//go:generate mockgen -destination test/mock/tribe.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories Tribe

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

//go:generate mockgen -destination test/mock/membership.gen.go -package mock go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories Membership

// Membership describes membership repository contract
type Membership interface {
	Join(ctx context.Context, entity *models.Person, ug models.PersonGroup) error
	Leave(ctx context.Context, entity *models.Person, ug models.PersonGroup) error
}
