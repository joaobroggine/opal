package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"
var commit = "none"
var buildDate = "unknown"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Opal version",
	Long:  "Show version information for the installed Opal binary.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opal CLI " + version)
		fmt.Println("Commit:", commit)
		fmt.Println("Build date:", buildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
