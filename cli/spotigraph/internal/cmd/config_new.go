package cmd

import (
	"fmt"
	"sort"

	defaults "github.com/mcuadros/go-defaults"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"go.zenithar.org/pkg/flags"
	"go.zenithar.org/pkg/log"
)

// -----------------------------------------------------------------------------

var configNewAsEnvFlag bool

// -----------------------------------------------------------------------------

var configNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Initialize a default configuration",
	Run: func(cmd *cobra.Command, args []string) {
		defaults.SetDefaults(conf)

		if !configNewAsEnvFlag {
			btes, err := toml.Marshal(*conf)
			if err != nil {
				log.Bg().Fatal("Error during configuration export")
			}
			fmt.Println(string(btes))
		} else {
			m := flags.AsEnvVariables(conf, "SPFG", true)
			keys := []string{}

			for k := range m {
				keys = append(keys, k)
			}

			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("export %s=\"%s\"\n", k, m[k])
			}
		}
	},
}
