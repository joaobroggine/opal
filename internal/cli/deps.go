package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"opal/internal/deps"
)

var dependenciesCmd = &cobra.Command{
	Use:   "dep",
	Short: "Analyze project dependencies",
	Long:  "Analyze project dependencies, detect unused packages and identify outdated versions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking dependencies...")
		err := deps.AnalyzeNode(".")
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(dependenciesCmd)
}