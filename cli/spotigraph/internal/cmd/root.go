package cmd

import (
	"fmt"
	"os"
	"strings"

	"go.zenithar.org/spotigraph/cli/spotigraph/internal/config"

	defaults "github.com/mcuadros/go-defaults"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/flags"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
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
	mainCmd.AddCommand(configCmd)
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

var (
	cfgFile string
	conf    = &config.Configuration{}
)

// -----------------------------------------------------------------------------

func initConfig() {
	for k := range flags.AsEnvVariables(conf, "", false) {
		log.CheckErr("Unable to bind environment variable", viper.BindEnv(strings.ToLower(strings.Replace(k, "_", ".", -1)), "SPFG_"+k), zap.String("var", "SPFG_"+k))
	}

	switch {
	case cfgFile != "":
		// If the config file doesn't exists, let's exit
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			log.Bg().Fatal("File doesn't exists", zap.Error(err))
		}
		fmt.Println("Reading configuration file", cfgFile)

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Bg().Fatal("Unable to read config", zap.Error(err))
		}
	default:
		defaults.SetDefaults(conf)
	}

	if err := viper.Unmarshal(conf); err != nil {
		log.Bg().Fatal("Unable to parse config", zap.Error(err))
	}
}
