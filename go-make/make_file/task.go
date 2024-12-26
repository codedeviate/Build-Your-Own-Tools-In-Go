package make_file

import "fmt"

// ResolveDependencies resolves all dependencies for a task recursively.
func ResolveDependencies(taskName string, tasks map[string]*Task, resolved map[string]bool) ([]string, error) {
	if resolved[taskName] {
		return nil, nil
	}

	task, exists := tasks[taskName]
	if !exists {
		return nil, fmt.Errorf("task '%s' not found", taskName)
	}

	var dependencies []string
	for _, dep := range task.Dependencies {
		subDeps, err := ResolveDependencies(dep, tasks, resolved)
		if err != nil {
			return nil, err
		}
		dependencies = append(dependencies, subDeps...)
	}
	resolved[taskName] = true
	dependencies = append(dependencies, taskName)
	return dependencies, nil
}
