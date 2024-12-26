package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/username/go-make/make_file"
)

func RunTask(taskName string, tasks map[string]*make_file.Task) error {
	resolved := make(map[string]bool)
	executionOrder, err := make_file.ResolveDependencies(taskName, tasks, resolved)
	if err != nil {
		return err
	}

	for _, task := range executionOrder {
		if len(tasks[task].Command) > 0 {
			for _, command := range tasks[task].Command {
				cmd := exec.Command("sh", "-c", strings.TrimPrefix(command, "@"))
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("task '%s' failed: %s", task, err.Error())
				}
			}
		}
	}

	return nil
}
