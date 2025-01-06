// commands/push.go
package commands

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func Push(args []string) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		return
	}

	err = repo.Push(&git.PushOptions{})
	if err != nil {
		fmt.Printf("Error pushing changes: %v\n", err)
	}
}
