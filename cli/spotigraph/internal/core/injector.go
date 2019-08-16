package core

import (
	"github.com/google/wire"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"

	"go.zenithar.org/spotigraph/cli/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
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

var pgRepositorySet = wire.NewSet(
	PosgreSQLConfig,
	postgresql.RepositorySet,
)

// -----------------------------------------------------------------------------

var localServiceSet = wire.NewSet(
	chapter.New,
)

// -----------------------------------------------------------------------------

// LocalPostgreSQLSet initialize the PGSQL Core context
var LocalPostgreSQLSet = wire.NewSet(
	pgRepositorySet,
	localServiceSet,
)
