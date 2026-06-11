package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"opal/internal/scanner"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run initial project analysis",
	Long:  "Run a full project analysis, including technology detection, dependency validation and environment inspection.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Initializing project...")

		root := "."

		// 1. Detect languages
		languages, err := scanner.DetectLanguages(root)
		if err != nil {
			fmt.Println("Error detecting languages:", err)
			return
		}

		fmt.Println("\nDetected languages:")
		for lang := range languages {
			fmt.Println("-", lang)
		}

		// 2. Check if is frontend
		isFrontend, err := scanner.IsFrontend(root)
		if err != nil {
			fmt.Println("Error checking frontend:", err)
			return
		}

		if isFrontend {
			fmt.Println("\nFrontend project detected")

			hasModules, err := scanner.HasNodeModules(root)
			if err != nil {
				fmt.Println("Error checking node_modules:", err)
				return
			}

			if hasModules {
				fmt.Println("node_modules: found")
			} else {
				fmt.Println("node_modules: missing")
			}
		}

		fmt.Println("\nAnalysis complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}