package core

import (
	"strings"

	"github.com/google/wire"
	mdb "go.zenithar.org/pkg/db/adapter/mongodb"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"
	rdb "go.zenithar.org/pkg/db/adapter/rethinkdb"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/mongodb"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql"
	"go.zenithar.org/spotigraph/internal/repositories/pkg/rethinkdb"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/internal/services/pkg/graph"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe"
	"go.zenithar.org/spotigraph/internal/services/pkg/user"
)

// -----------------------------------------------------------------------------

// RethinkDBConfig declares a Database configuration provider for Wire
func RethinkDBConfig(cfg *config.Configuration) *rdb.Configuration {
	return &rdb.Configuration{
		AutoMigrate: cfg.Core.Local.AutoMigrate,
		Addresses:   strings.Split(cfg.Core.Local.Hosts, ","),
		Database:    cfg.Core.Local.Database,
		AuthKey:     cfg.Core.Local.Password,
	}
}

var rdbRepositorySet = wire.NewSet(
	RethinkDBConfig,
	rethinkdb.RepositorySet,
)

// MongoDBConfig declares a Database configuration provider for Wire
func MongoDBConfig(cfg *config.Configuration) *mdb.Configuration {
	return &mdb.Configuration{
		AutoMigrate:      cfg.Core.Local.AutoMigrate,
		ConnectionString: cfg.Core.Local.Hosts,
		DatabaseName:     cfg.Core.Local.Database,
		Username:         cfg.Core.Local.Username,
		Password:         cfg.Core.Local.Password,
	}
}

var mgoRepositorySet = wire.NewSet(
	MongoDBConfig,
	mongodb.RepositorySet,
)

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
	user.New,
	chapter.New,
	squad.New,
	guild.New,
	tribe.New,
	graph.New,
)

// -----------------------------------------------------------------------------

// LocalPostgreSQLSet initialize the PGSQL Core context
var LocalPostgreSQLSet = wire.NewSet(
	pgRepositorySet,
	localServiceSet,
)

// LocalRethinkDBSet initialize the RDB Core context
var LocalRethinkDBSet = wire.NewSet(
	rdbRepositorySet,
	localServiceSet,
)

// LocalMongoDBSet initialize the MGO Core context
var LocalMongoDBSet = wire.NewSet(
	mgoRepositorySet,
	localServiceSet,
)
