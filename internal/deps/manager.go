package deps

import (
	"fmt"
	"os/exec"
	"opal/internal/scanner"
)

func AnalyzeNode(root string) error {

	projects, err := scanner.FindNodeProjects(root)
	if err != nil {
		return err
	}

	if len(projects) == 0 {
		fmt.Println("No Node.js projects found")
		return nil
	}

	for _, p := range projects {

		fmt.Println("\nProject:", p)

		cmd := exec.Command("npm", "outdated")
		cmd.Dir = p

		output, _ := cmd.CombinedOutput()

		if len(output) == 0 || string(output) == "{}" {
			fmt.Println("No outdated dependencies found")
			continue
		}

		fmt.Println(string(output))
	}

	return nil
}