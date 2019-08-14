package core

import (
	"github.com/google/wire"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"

	"go.zenithar.org/spotigraph/cli/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe"
	"go.zenithar.org/spotigraph/internal/services/pkg/user"
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

// Squad builder provider for wire
func Squad(squads repositories.Squad, users repositories.User, memberships repositories.Membership) services.Squad {
	return squad.NewWithDecorators(squads, users, memberships, squad.WithTracer())
}

var localServiceSet = wire.NewSet(
	user.New,
	chapter.New,
	Squad,
	guild.New,
	tribe.New,
)

// -----------------------------------------------------------------------------

// LocalPostgreSQLSet initialize the PGSQL Core context
var LocalPostgreSQLSet = wire.NewSet(
	pgRepositorySet,
	localServiceSet,
)
