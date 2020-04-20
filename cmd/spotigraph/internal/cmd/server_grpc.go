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

package cmd

import (
	"context"
	"net"

	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
	"go.zenithar.org/pkg/platform/actors"
	"go.zenithar.org/spotigraph/build/version"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/dispatchers/grpc"
)

// -----------------------------------------------------------------------------

func serverGRPCCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "grpc",
		Short: "Starts the spotigraph gRPC server",
		Run:   runServerGRPC,
	}

	return c
}

func runServerGRPC(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize config
	initConfig()

	// Starting banner
	log.For(ctx).Info("Starting spotigraph gRPC server ...")

	// Start goroutine group
	err := platform.Serve(ctx, platform.Server{
		Debug:           conf.Debug.Enable,
		Name:            "spotigraph-grpc",
		Version:         version.Version,
		Revision:        version.Revision,
		Instrumentation: conf.Instrumentation,
		Builder: func(ln net.Listener, group run.Group) {
			// Initialize gRPC server
			server, err := grpc.New(ctx, conf)
			if err != nil {
				log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
			}

			// Register gRPC actor
			actors.GRPC(server, ln)(ctx, &group)
		},
	})
	log.CheckErrCtx(ctx, "Unable to run application", err)
}
