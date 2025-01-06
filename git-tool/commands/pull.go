// commands/pull.go
package commands

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func Pull(args []string) {
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

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		fmt.Printf("Error pulling changes: %v\n", err)
	}
}
