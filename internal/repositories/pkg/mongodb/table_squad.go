package mongodb

import (
	"context"

	mongowrapper "github.com/opencensus-integrations/gomongowrapper"
	api "go.zenithar.org/pkg/db"
	db "go.zenithar.org/pkg/db/adapter/mongodb"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
)

type mgoSquadRepository struct {
	adapter *db.Default
}

// NewSquadRepository returns an initialized MongoDB repository for squads
func NewSquadRepository(cfg *db.Configuration, session *mongowrapper.WrappedClient) repositories.Squad {
	return &mgoSquadRepository{
		adapter: db.NewCRUDTable(session, cfg.DatabaseName, SquadTableName),
	}
}

// ------------------------------------------------------------

func (r *mgoSquadRepository) Create(ctx context.Context, entity *models.Squad) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Insert(ctx, entity)
}

func (r *mgoSquadRepository) Get(ctx context.Context, id string) (*models.Squad, error) {
	var entity models.Squad

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"id": id,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mgoSquadRepository) Update(ctx context.Context, entity *models.Squad) error {
	// Validate entity first
	if err := entity.Validate(); err != nil {
		return err
	}

	return r.adapter.Update(ctx, map[string]interface{}{
		"name":             entity.Name,
		"meta":             entity.Meta,
		"product_owner_id": entity.ProductOwnerID,
		"member_ids":       entity.MemberIDs,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *mgoSquadRepository) Delete(ctx context.Context, id string) error {
	return r.adapter.Delete(ctx, id)
}

func (r *mgoSquadRepository) Search(ctx context.Context, filter *repositories.SquadSearchFilter, pagination *api.Pagination, sortParams *api.SortParameters) ([]*models.Squad, int, error) {
	panic("Not implemented")
}

func (r *mgoSquadRepository) FindByName(ctx context.Context, name string) (*models.Squad, error) {
	var entity models.Squad

	if err := r.adapter.WhereAndFetchOne(ctx, map[string]interface{}{
		"name": name,
	}, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mgoSquadRepository) AddMembers(ctx context.Context, id string, users ...*models.User) error {
	// Retrieve squad entity
	entity, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	// Add user as members
	for _, u := range users {
		entity.AddMember(u)
	}

	// Update members
	return r.adapter.Update(ctx, map[string]interface{}{
		"member_ids": entity.MemberIDs,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}

func (r *mgoSquadRepository) RemoveMembers(ctx context.Context, id string, users ...*models.User) error {
	// Retrieve squad entity
	entity, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	// Remove user from members
	for _, u := range users {
		entity.RemoveMember(u)
	}

	// Update members
	return r.adapter.Update(ctx, map[string]interface{}{
		"member_ids": entity.MemberIDs,
	}, map[string]interface{}{
		"id": entity.ID,
	})
}
