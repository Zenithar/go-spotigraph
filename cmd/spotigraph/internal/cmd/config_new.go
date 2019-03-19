package cmd

import (
	"fmt"
	"sort"

	"go.zenithar.org/spotigraph/pkg/flag"

	defaults "github.com/mcuadros/go-defaults"
	toml "github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// -----------------------------------------------------------------------------

var (
	configNewAsEnvFlag bool
)

// -----------------------------------------------------------------------------

var configNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Initialize a default configuration",
	Run: func(cmd *cobra.Command, args []string) {
		defaults.SetDefaults(conf)

		if !configNewAsEnvFlag {
			btes, err := toml.Marshal(*conf)
			if err != nil {
				logrus.WithError(err).Fatalln("Error during configuration export")
			}
			fmt.Println(string(btes))
		} else {
			m := flag.AsEnvVariables(conf, "SPFG", true)
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
