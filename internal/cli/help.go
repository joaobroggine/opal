package cli

import "fmt"

func Help() {
	fmt.Println("Usage:")
	fmt.Println("  opal <command>")
	fmt.Println()

	fmt.Println("Available commands:")
	fmt.Println("  init      Initialize project")
	fmt.Println("  help      Show help menu")
	fmt.Println("  version   Show CLI version")
	fmt.Println("  dep       Check dependencies")
	fmt.Println("  readme    Generate README")
	fmt.Println()
}