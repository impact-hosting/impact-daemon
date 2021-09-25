package commands

import (
	"fmt"
	"log"

	"github.com/impact-hosting/impact-daemon/routes"
	"github.com/impact-hosting/impact-daemon/util"
	"github.com/spf13/cobra"
)

var ()

func Run() {
	initialise()
	if err := defaultCommand.Execute(); err != nil {
		log.Fatalf("Command failed to run:\n%s", err)
	}
}

var defaultCommand = &cobra.Command{
	Use:   "impact-daemon",
	Short: "Run the Impact Daemon API",
	Run: func(cmd *cobra.Command, args []string) {
		routes.Launch()
	},
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Output the version of the ImpactDaemon",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Impact Daemon %s", util.Version)
	},
}

func initialise() {
	defaultCommand.AddCommand(versionCommand)
}
