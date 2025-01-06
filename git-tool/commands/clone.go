// commands/clone.go
package commands

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: git-tool clone <repository-url>")
		return
	}

	repoURL := args[0]
	_, err := git.PlainClone(".", false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
	}
}
