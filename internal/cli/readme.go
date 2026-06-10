package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Generate project documentation",
	Long:  "Generate a README file based on the detected project structure, technologies and dependencies.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating README...")
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)
}