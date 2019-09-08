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

package postgresql

import (
	"context"

	"go.opencensus.io/trace"
	db "go.zenithar.org/pkg/db/adapter/postgresql"
	"go.zenithar.org/spotigraph/internal/helpers"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type pgMembershipRepository struct {
	adapter *db.Default
}

// NewMembershipRepository returns an initialized PostgreSQL repository for memberships
func NewMembershipRepository(cfg *db.Configuration, session *sqlx.DB) repositories.Membership {
	// Defines allowed columns
	defaultColumns := []string{
		"id", "person_id", "group_id", "group_type",
	}

	// Sortable columns
	sortableColumns := []string{
		"id", "person_id", "group_id", "group_type",
	}

	return &pgMembershipRepository{
		adapter: db.NewCRUDTable(session, "", MembershipTableName, defaultColumns, sortableColumns),
	}
}

// -----------------------------------------------------------------------------

type sqlMembership struct {
	ID        string `db:"id"`
	PersonID  string `db:"person_id"`
	GroupID   string `db:"group_id"`
	GroupType string `db:"group_type"`
}

// -----------------------------------------------------------------------------

func (r *pgMembershipRepository) Join(ctx context.Context, entity *models.Person, ug models.PersonGroup) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.membership.Join")
	span.AddAttributes(
		trace.StringAttribute("person_id", entity.ID),
		trace.StringAttribute("group_type", ug.GetGroupType()),
		trace.StringAttribute("group_id", ug.GetGroupID()),
	)
	defer span.End()

	return r.adapter.Create(ctx, &sqlMembership{
		ID:        helpers.IDGeneratorFunc(),
		PersonID:  entity.ID,
		GroupID:   ug.GetGroupID(),
		GroupType: ug.GetGroupType(),
	})
}

func (r *pgMembershipRepository) Leave(ctx context.Context, entity *models.Person, ug models.PersonGroup) error {
	ctx, span := trace.StartSpan(ctx, "postgresql.membership.Leave")
	span.AddAttributes(
		trace.StringAttribute("person_id", entity.ID),
		trace.StringAttribute("group_type", ug.GetGroupType()),
		trace.StringAttribute("group_id", ug.GetGroupID()),
	)
	defer span.End()

	return r.adapter.RemoveOne(ctx, &map[string]interface{}{
		"person_id":  entity.ID,
		"group_type": ug.GetGroupType(),
		"group_id":   ug.GetGroupID(),
	})
}
