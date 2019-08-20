package core

import (
	"github.com/google/wire"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"

	"go.zenithar.org/spotigraph/cli/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/internal/services/pkg/person"
)

// -----------------------------------------------------------------------------

// PosgreSQLConfig declares a Database configuration provider for Wire
func PosgreSQLConfig(cfg *config.Configuration) *pgdb.Configuration {
	return &pgdb.Configuration{
		AutoMigrate:      cfg.Core.Local.AutoMigrate,
		ConnectionString: cfg.Core.Local.Hosts,
		Username:         cfg.Core.Local.Username,
		Password:         cfg.Core.Local.Password,
	}
}

// -----------------------------------------------------------------------------

var localServiceSet = wire.NewSet(
	chapter.New,
	person.New,
)

// -----------------------------------------------------------------------------

// PostgreSQLSet initialize the PGSQL Core context
var PostgreSQLSet = wire.NewSet(
	PosgreSQLConfig,
	postgresql.RepositorySet,
	localServiceSet,
)
