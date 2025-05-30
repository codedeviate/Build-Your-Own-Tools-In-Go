// commands/status.go
package commands

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func Status(args []string) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Printf("Error getting worktree: %v\n", err)
		return
	}

	status, err := w.Status()
	if err != nil {
		fmt.Printf("Error checking status: %v\n", err)
		return
	}

	fmt.Println(status)
}
