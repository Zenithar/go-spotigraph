// +build integration

package integration

import (
	"context"

	"github.com/pkg/errors"

	"go.zenithar.org/pkg/testing/containers/database"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
)

func postgreSQLConnection(ctx context.Context) (func(), error) {
	// Initialize connection and/or container
	conn, cfg, err := database.ConnectToPostgreSQL(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to initialize database server")
	}

	// Try to contact server
	if err = conn.Ping(); err != nil {
		return nil, nil, errors.Wrap(err, "unable to contact database")
	}

	// Migrate schema
	if _, err = postgresql.CreateSchemas(conn); err != nil {
		return nil, nil, errors.Wrap(err, "unable to initialize database schema")
	}

	// Build repositories
	userRepositories["postgresql"] = postgresql.NewUserRepository(nil, conn)
	squadRepositories["postgresql"] = postgresql.NewSquadRepository(nil, conn)
	chapterRepositories["postgresql"] = postgresql.NewChapterRepository(nil, conn)
	tribeRepositories["postgresql"] = postgresql.NewTribeRepository(nil, conn)
	guildRepositories["postgresql"] = postgresql.NewGuildRepository(nil, conn)

	// Return result
	return nil
}
