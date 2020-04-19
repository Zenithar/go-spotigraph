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

package grpc

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
)

type application struct {
	cfg    *config.Configuration
	server *grpc.Server
}

var (
	app  *application
	once sync.Once
)

// -----------------------------------------------------------------------------

// New initialize the application
func New(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	var err error

	once.Do(func() {
		// Initialize application
		app = &application{}

		// Apply configuration
		if err := app.ApplyConfiguration(cfg); err != nil {
			log.For(ctx).Fatal("Unable to initialize server settings", zap.Error(err))
		}

		// Initialize Core components
		switch cfg.Core.Mode {
		case "local":
			app.server, err = setupLocalPostgreSQL(ctx, cfg)
		default:
			log.For(ctx).Fatal("Invalid core mode, use 'local' only.")
		}
	})

	// Return server
	return app.server, err
}

// -----------------------------------------------------------------------------

// ApplyConfiguration apply the configuration after checking it
func (s *application) ApplyConfiguration(cfg interface{}) error {
	// Check configuration validity
	if err := s.checkConfiguration(cfg); err != nil {
		return err
	}

	// Apply to current component (type assertion done if check)
	s.cfg, _ = cfg.(*config.Configuration)

	// No error
	return nil
}

// -----------------------------------------------------------------------------

func (s *application) checkConfiguration(obj interface{}) error {
	// Check via type assertion
	cfg, ok := obj.(*config.Configuration)
	if !ok {
		return fmt.Errorf("server: invalid configuration")
	}

	// nolint
	switch cfg.Core.Mode {
	case "local":
		if cfg.Core.Local.Hosts == "" {
			return fmt.Errorf("server: database hosts list is mandatory")
		}
	case "remote":
	default:
		return fmt.Errorf("server: invalid core mode, grpc only support 'local'")
	}

	// No error
	return nil
}
