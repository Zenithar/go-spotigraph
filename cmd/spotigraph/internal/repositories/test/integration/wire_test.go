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

// +build integration

package integration

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	pg "go.zenithar.org/pkg/testing/containers/database/postgresql"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories/pkg/postgresql"
)

func postgreSQLConnection(ctx context.Context) (func(), error) {
	// Initialize connection and/or container
	conn, _, err := pg.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize database server")
	}

	// Try to contact server
	if err = conn.Ping(); err != nil {
		return nil, errors.Wrap(err, "unable to contact database")
	}

	// Migrate schema
	n, err := postgresql.CreateSchemas(conn)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize database schema")
	}

	// Log migration
	log.For(ctx).Info("Applyied migrations to database", zap.Int("level", n))

	// Build repositories
	userRepositories["postgresql"] = postgresql.NewUserRepository(nil, conn)
	squadRepositories["postgresql"] = postgresql.NewSquadRepository(nil, conn)
	chapterRepositories["postgresql"] = postgresql.NewChapterRepository(nil, conn)
	tribeRepositories["postgresql"] = postgresql.NewTribeRepository(nil, conn)
	guildRepositories["postgresql"] = postgresql.NewGuildRepository(nil, conn)

	// Return result
	return func() {
		log.SafeClose(conn, "unable to close connection")
	}, nil
}
