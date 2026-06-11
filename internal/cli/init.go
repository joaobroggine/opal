package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"opal/internal/deps"
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

		// 2. Node.js / frontend dependency analysis (delegated)
		err = deps.AnalyzeNode(root)
		if err != nil {
			fmt.Println("Error analyzing Node projects:", err)
			return
		}

		fmt.Println("Analysis complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}