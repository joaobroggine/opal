package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run initial project analysis",
	Long:  "Run a full project analysis, including technology detection, dependency validation and environment inspection.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing project...")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}