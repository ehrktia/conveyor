/*
Copyright © 2024 Conveyor CI Contributors
*/
package cli

import (
	"fmt"

	apiServer "github.com/open-ug/conveyor/cmd/api"
	sampledriver "github.com/open-ug/conveyor/cmd/sample-driver"
	"github.com/spf13/cobra"
)

var APIServerCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the Conveyor Service",
	Long: `Start the Conveyor Service

`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		if port == "" {
			port = "8080"
		}
		apiServer.StartServer(port)
	},
}

var SampleDriverCmd = &cobra.Command{
	Use:   "sampledriver",
	Short: "Start the Sample Driver",
	Long:  `Start the Sample Driver for testing purposes, You can specify the name and resources the driver will manage.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		resources, err := cmd.Flags().GetStringSlice("resources")
		if err != nil {
			fmt.Println("Error getting resources flag: ", err)
			fmt.Println("Defaulting to 'pipe' resource")
			resources = []string{"pipe"}
		}
		if name == "" {
			name = "sampledriver"
		}
		if len(resources) == 0 {
			resources = []string{"pipe"}
		}
		sampledriver.Listen(name, resources)
	},
}

const DEFAULT_CONFIG = "DEFAULT_CONFIG"
const DEFAULT_CA_CERT = "DEFAULT_CA_CERT"
const DEFAULT_PRIVATE_KEY = "DEFAULT_PRIVATE_KEY"
const DEFAULT_CA = "DEFAULT_CA"

var cliConfig string
var apiPort string
var natsPort string
var authEnable bool
var caCert string
var privateKey string
var cert string

func init() {
	APIServerCmd.Flags().StringP("port", "p", "",
		`Port to run the API Server on (default: 3000)`)
	SampleDriverCmd.Flags().StringP("name", "n", "sampledriver",
		`Name of the driver`)
	SampleDriverCmd.Flags().StringSliceP("resources", "r", []string{"pipe"},
		`Resources the driver will manage`)
	initCmd.Flags().StringVarP(&cliConfig, "config", "c",
		"", "config for application")
	initCmd.Flags().StringVarP(&apiPort, "api-port", "a",
		"", "api port for the server")
	initCmd.Flags().StringVarP(&natsPort, "nats-port", "m",
		"", "nats port")
	initCmd.Flags().BoolVarP(&authEnable, "auth-enable", "e",
		false, "auth-enabled")
	initCmd.Flags().StringVarP(&caCert, "ca", "n",
		"", "ca-cert")
	initCmd.Flags().StringVarP(&privateKey, "private-key", "p",
		"", "private-key")
	initCmd.Flags().StringVarP(&cert, "crt", "r",
		"", "cert")
}
