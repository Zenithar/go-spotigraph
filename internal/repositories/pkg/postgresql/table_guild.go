package postgresql

import (
	"context"
	"encoding/json"
	"strings"

	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	sq "github.com/Masterminds/squirrel"
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
		"id", "name", "meta", "member_ids", "leader_id",
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
	ID        string `db:"id"`
	Name      string `db:"name"`
	Meta      string `db:"meta"`
	MemberIDs string `db:"member_ids"`
	LeaderID  string `db:"leader_id"`
}

func toGuildSQL(entity *models.Guild) (*sqlGuild, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := json.Marshal(entity.MemberIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlGuild{
		ID:        entity.ID,
		Name:      entity.Name,
		Meta:      string(meta),
		MemberIDs: string(members),
		LeaderID:  entity.LeaderID,
	}, nil
}

func (dto *sqlGuild) ToEntity() (*models.Guild, error) {
	entity := &models.Guild{
		ID:       dto.ID,
		Name:     dto.Name,
		LeaderID: dto.LeaderID,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Membership
	err = json.Unmarshal([]byte(dto.MemberIDs), &entity.MemberIDs)
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
		"member_ids": obj.MemberIDs,
		"leader_id":  obj.LeaderID,
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
	var results []sqlGuild

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	entities := make([]*models.Guild, len(results))
	if len(results) == 0 {
		return entities, count, api.ErrNoResult
	}

	for i, entity := range results {
		e, err := entity.ToEntity()
		if err != nil {
			continue
		}
		entities[i] = e
	}

	// Return results and total count
	return entities, count, nil
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

// -----------------------------------------------------------------------------

func (r *pgGuildRepository) buildFilter(filter *repositories.GuildSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.GuildID)) > 0 {
			clauses["guild_id"] = filter.GuildID
		}
		if len(strings.TrimSpace(filter.Name)) > 0 {
			clauses["name"] = filter.Name
		}

		return clauses
	}

	return nil
}
