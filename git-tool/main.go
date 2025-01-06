// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/git-tool/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git-tool <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "clone":
		commands.Clone(args)
	case "commit":
		commands.Commit(args)
	case "push":
		commands.Push(args)
	case "pull":
		commands.Pull(args)
	case "status":
		commands.Status(args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
