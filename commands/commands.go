package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/impact-hosting/impact-daemon/util"
	"github.com/spf13/cobra"
)

func Run() {
	if err := defaultCommand.Execute(); err != nil {
		log.Fatalf("Command failed to run:\n%s", err)
	}
}

func VersionCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Impact Daemon %s", util.Version)
}

var defaultCommand = &cobra.Command{
	Use:   "impact-daemon",
	Short: "Run the Impact Daemon API",
	Run: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			versionCommand.Execute()
			os.Exit(0)
		}
	},
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Output the version of the ImpactDaemon",
	Run:   VersionCommand,
}
