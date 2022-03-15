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

package tenant

import (
	"context"
	"errors"

	"zntr.io/spotigraph/pkg/cursor"
)

type ID string

// Tenant describes the Tenant domain object contract.
type Tenant interface {
	// GetID returns the domain object identifier created by the repository.
	GetID() ID
	// GetLabel returns the tenant label.
	GetLabel() string
	// GetSlug returns the slugified label for url usages.
	GetSlug() string
	// IsLocked returns the tenant lockage status.
	IsLocked() bool
}

var (
	ErrTenantNotSaved   = errors.New("unable to save the given tenant in the repository")
	ErrTenantNotDeleted = errors.New("unable to delete the given tenant from the repository")
)

// SearchFilter represents tenant entity collection search criteria
type SearchFilter struct {
	Limit     *uint64
	Cursor    *string
	ObjectIDs []string
	Label     *string
	Slug      *string
}

// ReaderRepository is the tenant contract definition for read-only operation.
type ReaderRepository interface {
	List(ctx context.Context, filter SearchFilter) ([]Tenant, *cursor.PageInfo, error)
	GetByID(ctx context.Context, id ID) (Tenant, error)
	GetByLabel(ctx context.Context, label string) (Tenant, error)
	GetBySlug(ctx context.Context, slug string) (Tenant, error)
}

// WriterRepository describes tenant repository contract for alteration operation.
type WriterRepository interface {
	Save(ctx context.Context, model Tenant) error
	Remove(ctx context.Context, model Tenant) error
}

type IDGenerator interface {
	NextID(ctx context.Context) (ID, error)
}

type Repository interface {
	IDGenerator
	ReaderRepository
	WriterRepository
}
