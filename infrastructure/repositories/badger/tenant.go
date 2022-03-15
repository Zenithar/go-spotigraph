// Licensed to Thibault Normand under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Thibault Normand licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package badger

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/timshannon/badgerhold/v4"

	"zntr.io/spotigraph/domain/tenant"
	"zntr.io/spotigraph/infrastructure/repositories"
	"zntr.io/spotigraph/pkg/cursor"
)

func Tenants(store *badgerhold.Store, log *zerolog.Logger) tenant.Repository {
	return &badgerTenantRepository{
		store: store,
		log:   log,
	}
}

// -----------------------------------------------------------------------------

type badgerTenantRepository struct {
	store *badgerhold.Store
	log   *zerolog.Logger
}

// -----------------------------------------------------------------------------

type tenantEntity struct {
	ID     string `json:"id" badgerhold:"key"`
	Label  string `json:"label" badgerhold:"unique"`
	Slug   string `json:"slug" badgerhold:"unique"`
	Locked bool   `json:"locked"`
}

var _ tenant.Tenant = (*tenantEntity)(nil)

func (s *tenantEntity) GetID() tenant.ID { return tenant.ID(s.ID) }
func (s *tenantEntity) GetLabel() string { return s.Label }
func (s *tenantEntity) GetSlug() string  { return s.Slug }
func (s *tenantEntity) IsLocked() bool   { return s.Locked }

// -----------------------------------------------------------------------------

func (r *badgerTenantRepository) NextID(ctx context.Context) (tenant.ID, error) {
	return tenant.ID(xid.New().String()), nil
}

// -----------------------------------------------------------------------------

func (r *badgerTenantRepository) List(ctx context.Context, filter tenant.SearchFilter) ([]tenant.Tenant, *cursor.PageInfo, error) {
	collection := []tenant.Tenant{}

	if err := r.store.ForEach(nil, func(record *tenantEntity) error {
		// Append to collection
		collection = append(collection, record)

		// No error
		return nil
	}); err != nil {
		return nil, nil, fmt.Errorf("badger: unable to list tenants: %w", err)
	}

	// No error
	return collection, nil, nil
}

func (r *badgerTenantRepository) GetByID(ctx context.Context, id tenant.ID) (tenant.Tenant, error) {
	var res tenantEntity

	if err := r.store.Get(string(id), &res); err != nil {
		if errors.Is(err, badgerhold.ErrNotFound) {
			return nil, repositories.ErrNoResult
		}
		return nil, fmt.Errorf("badger: unable to lookup by id: %w", err)
	}

	// No error
	return &res, nil
}

func (r *badgerTenantRepository) GetByLabel(ctx context.Context, label string) (tenant.Tenant, error) {
	var res tenantEntity

	if err := r.store.FindOne(&res, badgerhold.Where("Label").Eq(label)); err != nil {
		if errors.Is(err, badgerhold.ErrNotFound) {
			return nil, repositories.ErrNoResult
		}
		return nil, fmt.Errorf("badger: unable to lookup by label: %w", err)
	}

	// No error
	return &res, nil
}

func (r *badgerTenantRepository) GetBySlug(ctx context.Context, slug string) (tenant.Tenant, error) {
	var res tenantEntity

	if err := r.store.FindOne(&res, badgerhold.Where("Slug").Eq(slug)); err != nil {
		if errors.Is(err, badgerhold.ErrNotFound) {
			return nil, repositories.ErrNoResult
		}
		return nil, fmt.Errorf("badger: unable to lookup by slug: %w", err)
	}

	// No error
	return &res, nil
}

// -----------------------------------------------------------------------------

func (r *badgerTenantRepository) Save(ctx context.Context, model tenant.Tenant) error {
	// Prepare entity
	persistable := &tenantEntity{
		ID:     string(model.GetID()),
		Label:  model.GetLabel(),
		Slug:   model.GetSlug(),
		Locked: model.IsLocked(),
	}

	// Save in persistence
	if err := r.store.Insert(persistable.ID, persistable); err != nil {
		return fmt.Errorf("badger: unable to save tenant: %v: %w", err, tenant.ErrTenantNotSaved)
	}

	// No error
	return nil
}

func (r *badgerTenantRepository) Remove(ctx context.Context, model tenant.Tenant) error {
	// Save in persistence
	if err := r.store.Delete(model.GetID(), &tenantEntity{}); err != nil {
		return fmt.Errorf("badger: unable to delete tenant: %v: %w", err, tenant.ErrTenantNotDeleted)
	}

	// No error
	return nil
}
