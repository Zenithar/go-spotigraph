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

package core

import (
	"github.com/google/wire"
	pgdb "go.zenithar.org/pkg/db/adapter/postgresql"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories/pkg/postgresql"
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

var localServiceSet = wire.NewSet()

// -----------------------------------------------------------------------------

// PostgreSQLSet initialize the PGSQL Core context
var PostgreSQLSet = wire.NewSet(
	PosgreSQLConfig,
	postgresql.RepositorySet,
	localServiceSet,
)
