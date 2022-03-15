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

package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/timshannon/badgerhold/v4"

	"zntr.io/spotigraph/cmd/tenant/router"
	"zntr.io/spotigraph/domain/tenant"
	"zntr.io/spotigraph/infrastructure/publisher"
	"zntr.io/spotigraph/infrastructure/repositories/badger"
	"zntr.io/spotigraph/pkg/web/respond"
)

func main() {
	// Logger
	logger := zerolog.New(os.Stdout)

	// Data storage
	options := badgerhold.DefaultOptions
	options.Dir = "tenants-data"
	options.ValueDir = "tenants-data"
	store, err := badgerhold.Open(options)
	if err != nil {
		logger.Fatal().Err(err)
		return
	}
	defer store.Close()

	// Initialize repositories
	tenants := badger.Tenants(store, &logger)

	// Initialize services
	createHandler := tenant.CreateHandler(tenants, publisher.Discard())

	// REST API
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// Homepage
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		respond.WithJSON(w, http.StatusOK, map[string]string{
			"service":   "tenant",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	})

	r.Route("/api/v1", func(api chi.Router) {
		api.Route("/tenants", func(tenantRouter chi.Router) {
			tenantRouter.Post("/", router.CreateTenant(&logger, createHandler))
		})
	})

	// Start and listen
	if err := http.ListenAndServe(":3333", r); err != nil {
		logger.Fatal().Err(err)
	}
}
