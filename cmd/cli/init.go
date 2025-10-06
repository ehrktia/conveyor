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
			init = []string{"config"}
		}
		if len(init) == 0 {
			init = []string{"config"}
		}
	},
}
