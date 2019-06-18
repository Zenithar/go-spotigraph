package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/pkg/config"
	configcmd "go.zenithar.org/pkg/config/cmd"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
	iconfig "go.zenithar.org/spotigraph/cli/spotigraph/internal/config"
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

	mainCmd.AddCommand(versionCmd)
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
