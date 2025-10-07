package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initalise using config",
	Long: `init command initialize with provided or default config and other
sub commands required for auth-enabled,api-port,nats-port,ca cert`,
	Run: func(cmd *cobra.Command, args []string) {
		init, err := cmd.Flags().GetStringSlice("init")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error initializing startup command:%v", err)
			fmt.Fprintf(os.Stderr, "%s", "using default config to startup")
			os.Exit(1)
		}
		if len(init) == 0 {
			// load default config
			config, err := loadDefaultConfig()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error loading default config:%v\n", err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "config used:%#v\n", config)

			init = []string{"config"}
		}
	},
}
