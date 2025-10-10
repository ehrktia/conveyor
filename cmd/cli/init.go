package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// conveyor init # run with defaults
// conveyor init --config /path/to/conveyor.yml # Custom config file, will be copied to config directory
// conveyor init --force                # overwrite existing files
// conveyor init --auth-enabled true                # configure is authentication is enabled or not
// conveyor init --api-port 8080            # configure api port
// conveyor init --nats-port 4222                # configure nats port
// conveyor init --ca /path/to/ca.pem --private-key /path/to/key --crt path/tp/certificate # to use existing certificates
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initalise using config",
	Long: `init command initialize with provided or default config and other
sub commands required for auth-enabled,api-port,nats-port,ca cert`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if cliConfig == "" {
			cliConfig = DEFAULT_CONFIG
		}
		if apiPort == "" {
			apiPort = "8080"
		}
		if natsPort == "" {
			natsPort = "4222"
		}
		if caCert == "" {
			caCert = DEFAULT_CA_CERT
		}
		if cert == "" {
			cert = DEFAULT_CA
		}
		if privateKey == "" {
			privateKey = DEFAULT_PRIVATE_KEY
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stderr, "using api port:%s\n", apiPort)
		fmt.Fprintf(os.Stderr, "using nats port:%s\n", natsPort)
		var conf Config
		var err error
		if cliConfig == DEFAULT_CONFIG {
			conf, err = loadDefaultConfig()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error loading default config:%v\n", err)
				os.Exit(1)
			}
		} else {
			conf, err = loadConfigFromFile(cliConfig)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error loading config:%v\n", err)
				os.Exit(1)
			}
		}
		fmt.Fprintf(os.Stderr, "using config :%#v\n", conf)
		fmt.Fprintf(os.Stderr, "using ca cert:%s\n", caCert)
		fmt.Fprintf(os.Stderr, "using cert:%s\n", cert)
	},
}
