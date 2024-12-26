package make_file

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/username/go-make/vars"
)

type Task struct {
	Name         string
	Dependencies []string
	Command      []string
}

// ParseMakefile reads and parses a Makefile.
func ParseMakefile(filename string) (map[string]*Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tasks := make(map[string]*Task)
	scanner := bufio.NewScanner(file)
	var currentTask *Task

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if !strings.HasPrefix(line, "\t") {
			if strings.Contains(line, "=") {
				parts := strings.SplitN(line, "=", 2)
				vars.VarList[parts[0]] = vars.EvalVar(parts[1])
				continue
			}
			parts := strings.Split(trimmedLine, ":")
			taskName := strings.TrimSuffix(parts[0], ":")
			currentTask = &Task{Name: taskName}
			tasks[taskName] = currentTask
			if len(parts) > 1 && parts[1] != "" {
				currentTask.Dependencies = strings.Fields(parts[1])
			}
		} else if currentTask != nil && strings.HasPrefix(line, "\t") {
			currentTask.Command = append(currentTask.Command, strings.TrimSpace(line))
		} else if currentTask != nil {
			currentTask.Dependencies = append(currentTask.Dependencies, strings.Fields(line)...)
		} else {
			log.Fatalf("Unknown action in line: %s", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
