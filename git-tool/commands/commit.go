// commands/commit.go
package commands

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func Commit(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: git-tool commit <message>")
		return
	}

	message := args[0]
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Printf("Error getting worktree: %v\n", err)
		return
	}

	_, err = w.Add(".")
	if err != nil {
		fmt.Printf("Error adding files: %v\n", err)
		return
	}

	_, err = w.Commit(message, &git.CommitOptions{})
	if err != nil {
		fmt.Printf("Error committing changes: %v\n", err)
	}
}
