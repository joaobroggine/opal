package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Opal version",
	Long:  "Show version information for the installed Opal binary.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opal CLI v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}