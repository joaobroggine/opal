package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opal",
	Short: "Opal CLI",
	Long:  `Opal helps developers analyze, maintain and document projects.

	Features:
	- Project scanning
	- Dependency analysis
	- README generation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Opal CLI! Use 'opal --help' to see available commands.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}