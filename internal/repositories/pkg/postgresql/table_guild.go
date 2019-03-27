package postgresql

import (
	"context"
	"encoding/json"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type pgGuildRepository struct {
	adapter *db.Default
}

// NewGuildRepository returns an initialized PostgreSQL repository for guilds
func NewGuildRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Guild {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "label", "meta", "member_ids",
	}

	// Sortable columns
	sortableColumns := []string{
		"name",
	}

	return &pgGuildRepository{
		adapter: db.NewCRUDTable(session, "", GuildTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlGuild struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	Meta    string `db:"meta"`
	Members string `db:"member_ids"`
}

func toGuildSQL(entity *models.Guild) (*sqlGuild, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := json.Marshal(entity.Members)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlGuild{
		ID:      entity.ID,
		Name:    entity.Name,
		Meta:    string(meta),
		Members: string(members),
	}, nil
}

func (dto *sqlGuild) ToEntity() (*models.Guild, error) {
	entity := &models.Guild{
		ID:   dto.ID,
		Name: dto.Name,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Membership
	err = json.Unmarshal([]byte(dto.Members), &entity.Members)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

// ------------------------------------------------------------

func (r *pgGuildRepository) Create(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}
	// Convert to DTO
	data, err := toGuildSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgGuildRepository) Get(ctx context.Context, id string) (*models.Guild, error) {
	var entity sqlGuild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgGuildRepository) Update(ctx context.Context, entity *models.Guild) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toGuildSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name":       obj.Name,
		"meta":       obj.Meta,
		"member_ids": obj.Members,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgGuildRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgGuildRepository) Search(ctx context.Context, filter *repositories.GuildSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Guild, int, error) {
	panic("Not implemented")
}

func (r *pgGuildRepository) FindByName(ctx context.Context, name string) (*models.Guild, error) {
	var entity sqlGuild

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}
