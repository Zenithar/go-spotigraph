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

package pgsql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"

	"zntr.io/spotigraph/pkg/types"
)

func makeTx(ctx context.Context, conn DBConn, log *zerolog.Logger) (pgx.Tx, func(error) error, error) {
	// Check arguments
	if types.IsNil(conn) {
		return nil, nil, fmt.Errorf("unable to create transaction with nil connection")
	}
	if log == nil {
		return nil, nil, fmt.Errorf("unable to create transaction with nil logger")
	}

	// Start transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Fail begin transaction")
		return nil, nil, err
	}

	// Close function
	closeFn := func(txErr error) error {
		var err error

		// Try to recover
		p := recover()

		switch {
		case p != nil:
			log.Debug().Msg("Transaction panic, try to rollback")

			if errRecover := tx.Rollback(ctx); errRecover != nil {
				log.Warn().Err(errRecover).Msg("error occurred during transaction rollback")
			}
			panic(p) // re-throw panic after Rollback
		case txErr != nil:
			log.Debug().Err(txErr).Msg("Transaction error, try to rollback")

			// err is non-nil; don't change it
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				log.Warn().Err(errRollback).Msg("error occurred during transaction rollback")
			}
		default:
			log.Debug().Msg("Commit transaction")
			err = tx.Commit(ctx) // if Commit returns error update err with commit err
		}

		return err
	}
	return tx, closeFn, nil
}

func withTransaction(ctx context.Context, conn DBConn, log *zerolog.Logger, queryFunc func(tx DBConn) error) (txErr error) {
	tx, closeFn, err := makeTx(ctx, conn, log)
	if err != nil {
		return fmt.Errorf("unable to initialize pgsql transaction: %w", err)
	}
	defer func() {
		if err := closeFn(txErr); err != nil {
			log.Err(err).Msg("unable to close transaction")
		}
	}()

	// Delegate to query function
	txErr = queryFunc(tx)
	return
}
