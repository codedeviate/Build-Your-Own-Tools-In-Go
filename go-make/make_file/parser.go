package make_file

import (
	"bufio"
	"fmt"
	"github.com/username/go-make/utils"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/username/go-make/vars"
)

type Task struct {
	Name         string
	Dependencies []string
	Command      []string
}

func ParseLogic(line string) bool {
	// Remove leading stuff until the first parenthesis
	pos := strings.Index(line, "(")
	if pos == -1 {
		return false
	}
	stmt := strings.Trim(line[:pos], " ")
	rawLogic := line[pos+1:]
	logic := ""

	idx := 0
	pCount := 0
	for idx < len(rawLogic) {
		if rawLogic[idx] == '(' {
			pCount++
		} else if rawLogic[idx] == ')' {
			pCount--
			if pCount < 0 {
				logic = strings.Trim(rawLogic[:idx], " ")
				break
			}
		}
		idx++
	}

	re := regexp.MustCompile(`\$\([a-zA-Z0-9_-]+\)`)

	for re.MatchString(logic) {
		logic = re.ReplaceAllStringFunc(logic, func(match string) string {
			// Replace the match with some value, for example, the key itself without $()
			key := match[2 : len(match)-1]
			value, ok := vars.VarList[key]
			if !ok {
				log.Fatalf("Variable %s not found", key)
			}
			return fmt.Sprintf("%s", value) // Example replacement
		})
	}

	// Evaluate the logic
	left := strings.Trim(strings.Split(logic, ",")[0], " ")
	right := strings.Trim(strings.Split(logic, ",")[1], " ")
	if stmt == "ifeq" {
		return left == right
	} else if stmt == "ifneq" {
		return left != right
	} else if stmt == "ifdef" {
		_, ok := vars.VarList[left]
		return ok
	} else if stmt == "ifndef" {
		_, ok := vars.VarList[left]
		return !ok
	}
	return false
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

	logicStack := utils.Stack{}

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "ifeq") {
			logicStack.Push(ParseLogic(trimmedLine))
		} else if strings.HasPrefix(trimmedLine, "ifneq") {
			logicStack.Push(ParseLogic(trimmedLine))
		} else if strings.HasPrefix(trimmedLine, "ifdef") {
			logicStack.Push(ParseLogic(trimmedLine))
		} else if strings.HasPrefix(trimmedLine, "ifndef") {
			logicStack.Push(ParseLogic(trimmedLine))
		} else if strings.HasPrefix(trimmedLine, "else") {
			logicStack.InvertCurrent()
		} else if strings.HasPrefix(trimmedLine, "endif") {
			if logicStack.IsEmpty() {
				log.Fatalf("Unexpected endif in line: %s", line)
			}
			logicStack.Pop()
		}
		if !logicStack.IsEmpty() {
			if value, ok := logicStack.Peek(); ok && !value {
				continue
			}
		}
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
