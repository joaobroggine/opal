package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var dependenciesCmd = &cobra.Command{
	Use:   "dep",
	Short: "Analyze project dependencies",
	Long:  "Analyze project dependencies, detect unused packages and identify outdated versions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking dependencies...")
	},
}

func init() {
	rootCmd.AddCommand(dependenciesCmd)
}