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

type pgTribeRepository struct {
	adapter *db.Default
}

// NewTribeRepository returns an initialized PostgreSQL repository for tribes
func NewTribeRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Tribe {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "name", "meta", "squad_ids", "leader_id",
	}

	// Sortable columns
	sortableColumns := []string{
		"name",
	}

	return &pgTribeRepository{
		adapter: db.NewCRUDTable(session, "", TribeTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlTribe struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Meta     string `db:"meta"`
	SquadIDs string `db:"squad_ids"`
	LeaderID string `db:"leader_id"`
}

func toTribeSQL(entity *models.Tribe) (*sqlTribe, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	squads, err := json.Marshal(entity.SquadIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlTribe{
		ID:       entity.ID,
		Name:     entity.Name,
		Meta:     string(meta),
		SquadIDs: string(squads),
	}, nil
}

func (dto *sqlTribe) ToEntity() (*models.Tribe, error) {
	entity := &models.Tribe{
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
	err = json.Unmarshal([]byte(dto.SquadIDs), &entity.SquadIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

// ------------------------------------------------------------

func (r *pgTribeRepository) Create(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toTribeSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgTribeRepository) Get(ctx context.Context, id string) (*models.Tribe, error) {
	var entity sqlTribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgTribeRepository) Update(ctx context.Context, entity *models.Tribe) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toTribeSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name":      obj.Name,
		"meta":      obj.Meta,
		"squad_ids": obj.SquadIDs,
		"leader_id": obj.LeaderID,
	}, map[string]interface{}{
		"id": obj.ID,
	})
}

func (r *pgTribeRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgTribeRepository) Search(ctx context.Context, filter *repositories.TribeSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Tribe, int, error) {
	var results []*models.Tribe

	count, err := r.adapter.Search(ctx, r.buildFilter(filter), pagination, sortParams, &results)
	if err != nil {
		return nil, count, err
	}

	if len(results) == 0 {
		return results, count, api.ErrNoResult
	}

	// Return results and total count
	return results, count, nil
}

func (r *pgTribeRepository) FindByName(ctx context.Context, name string) (*models.Tribe, error) {
	var entity sqlTribe

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

// -----------------------------------------------------------------------------

func (r *pgTribeRepository) buildFilter(filter *repositories.TribeSearchFilter) interface{} {
	if filter != nil {
		clauses := sq.Eq{
			"1": "1",
		}

		if len(strings.TrimSpace(filter.TribeID)) > 0 {
			clauses["id"] = filter.TribeID
		}
		if len(strings.TrimSpace(filter.Name)) > 0 {
			clauses["name"] = filter.Name
		}

		return clauses
	}

	return nil
}
