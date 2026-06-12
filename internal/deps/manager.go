package deps

import (
	"fmt"
	"opal/internal/scanner"
	"os"
	"os/exec"
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

		fmt.Println("Running npm install...")
		install := exec.Command("npm", "install")
		install.Dir = p
		install.Stdout = os.Stdout
		install.Stderr = os.Stderr

		if err := install.Run(); err != nil {
			fmt.Println("Error running npm install:", err)
			continue
		}

		cmd := exec.Command("npm", "outdated")
		cmd.Dir = p

		output, _ := cmd.CombinedOutput()

		if len(output) == 0 || string(output) == "{}" {
			fmt.Println("No outdated dependencies found")
			continue
		}

		fmt.Println(string(output))

		var response string
		fmt.Println("Do you want update your dependencies? (y/n)")
		fmt.Scanln(&response)

		if response == "y" {
			update := exec.Command("npm", "update")
			update.Dir = p
			update.Stdout = os.Stdout
			update.Stderr = os.Stderr

			if err := update.Run(); err != nil {
				fmt.Println("Error running npm update:", err)
				continue
			}
			fmt.Println("")
		} else {
			fmt.Println("Skipping npm update")
		}

	}

	return nil
}
