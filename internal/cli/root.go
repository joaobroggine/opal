package cli

import (
	"fmt"
	"os"
)

type Command func()

var commands = map[string]Command{
	"init": Init,
	"help": Help,
	"version": Version,
	"dep": Dependencies,
	"readme": Readme,
}

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use: opal help for a list of commands.")
		return
	}

	command := os.Args[1]

	if cmd, exists := commands[command]; exists {
		cmd()
	} else {
		fmt.Printf("Unknown command: %s\n", command)
	}
}