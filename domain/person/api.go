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

package person

import (
	"context"
	"errors"

	"zntr.io/spotigraph/pkg/cursor"
)

type ID string

// Person describes the Person domain object contract.
type Person interface {
	// GetID returns the domain object identifier created by the repository.
	GetID() ID
	// GetPrincipal returns the person identity.
	GetPrincipal() string
	// IsLocked returns the identity lockage status.
	IsLocked() bool
}

var (
	// ErrPersonNotFound is raised when the person entity lookup returned no result.
	ErrPersonNotFound = errors.New("unable to resolve person entity")
)

// SearchFilter represents person entity collection search criteria
type SearchFilter struct {
	Limit     *uint64
	Cursor    *string
	ObjectIDs []string
	Principal *string
}

// ReaderRepository is the person contract definition for read-only operation.
type ReaderRepository interface {
	List(ctx context.Context, filter SearchFilter) ([]Person, *cursor.PageInfo, error)
	GetByID(ctx context.Context, id ID) (Person, error)
	GetByPrincipal(ctx context.Context, principal string) (Person, error)
}

// WriterRepository describes person repository contract for alteration operation.
type WriterRepository interface {
	Save(ctx context.Context, model Person) error
	Remove(ctx context.Context, model Person) error
}

type IDGenerator interface {
	NextID(ctx context.Context) (ID, error)
}

type Repository interface {
	IDGenerator
	ReaderRepository
	WriterRepository
}
