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

type pgUserRepository struct {
	adapter *db.Default
}

// NewUserRepository returns an initialized PostgreSQL repository for users
func NewUserRepository(cfg *db.Configuration, session *sqlx.DB) repositories.User {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "prn", "meta",
	}

	// Sortable columns
	sortableColumns := []string{
		"prn",
	}

	return &pgUserRepository{
		adapter: db.NewCRUDTable(session, "", UserTableName, defaultColumns, sortableColumns),
	}
}

// ------------------------------------------------------------

type sqlUser struct {
	ID        string `db:"id"`
	Principal string `db:"name"`
	Meta      string `db:"meta"`
}

func toUserSQL(entity *models.User) (*sqlUser, error) {
	meta, err := json.Marshal(entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqlUser{
		ID:        entity.ID,
		Principal: entity.Principal,
		Meta:      string(meta),
	}, nil
}

func (dto *sqlUser) ToEntity() (*models.User, error) {
	entity := &models.User{
		ID:        dto.ID,
		Principal: dto.Principal,
	}

	// Decode JSON columns

	// Metadata
	err := json.Unmarshal([]byte(dto.Meta), &entity.Meta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

// ------------------------------------------------------------

func (r *pgUserRepository) Create(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Convert to DTO
	data, err := toUserSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Create(ctx, data)
}

func (r *pgUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	var entity sqlUser

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}

func (r *pgUserRepository) Update(ctx context.Context, entity *models.User) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	// Intermediary DTO
	obj, err := toUserSQL(entity)
	if err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"meta": obj.Meta,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *pgUserRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.RemoveOne(ctx, map[string]interface{}{
		"id": id,
	})
}

func (r *pgUserRepository) Search(ctx context.Context, filter *repositories.UserSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.User, int, error) {
	panic("Not implemented")
}

func (r *pgUserRepository) FindByPrincipal(ctx context.Context, principal string) (*models.User, error) {
	var entity sqlUser

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"prn": principal,
	}, &entity); err != nil {
		return nil, err
	}

	return entity.ToEntity()
}
