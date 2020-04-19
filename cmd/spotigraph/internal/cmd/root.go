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
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/config"
	configcmd "go.zenithar.org/pkg/config/cmd"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/build/version"
	iconfig "go.zenithar.org/spotigraph/cmd/spotigraph/internal/config"
)

// -----------------------------------------------------------------------------

var (
	cfgFile string
	conf    = &iconfig.Configuration{}
)

// -----------------------------------------------------------------------------

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "spotigraph",
	Short: "Spotify Agile methodology graph microservice",
}

func init() {
	mainCmd.Flags().StringVar(&cfgFile, "config", "", "config file")

	mainCmd.AddCommand(version.Command())
	mainCmd.AddCommand(configcmd.NewConfigCommand(conf, "SPFG"))
	mainCmd.AddCommand(serverCmd)
	mainCmd.AddCommand(clientCmd)
}

// -----------------------------------------------------------------------------

// Execute main command
func Execute() error {
	feature.DefaultMutableGate.AddFlag(mainCmd.Flags())
	return mainCmd.Execute()
}

// -----------------------------------------------------------------------------

func initConfig() {
	if err := config.Load(conf, "SPFG", cfgFile); err != nil {
		log.Bg().Fatal("Unable to load settings", zap.Error(err))
	}
}
