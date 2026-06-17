package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"
var Commit = "none"
var BuildDate = "unknown"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Opal version",
	Long:  "Show version information for the installed Opal binary.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opal CLI " + Version)
		fmt.Println("Commit:", Commit)
		fmt.Println("Build date:", BuildDate)
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(versionCmd)
}
