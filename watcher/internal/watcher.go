package internal

import (
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gobwas/glob"
)

// Watcher monitors file changes and executes actions based on a configuration.
type Watcher struct {
	paths    []string
	debounce map[string]map[string]*time.Timer
	config   Config
	logger   *Logger
	watcher  *fsnotify.Watcher
}

// NewWatcher initializes a new Watcher.
func NewWatcher(paths []string, config Config, logger *Logger) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		paths:   paths,
		config:  config,
		logger:  logger,
		watcher: watcher,
	}, nil
}

// Start begins monitoring the file system for changes.
func (w *Watcher) Start() error {
	defer w.watcher.Close()

	w.debounce = make(map[string]map[string]*time.Timer)

	// Add paths to the watcher
	for _, path := range w.paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			w.logger.Error("Invalid path: %v", err)
			continue
		}

		err = w.watcher.Add(absPath)
		if err != nil {
			w.logger.Error("Failed to watch %s: %v", absPath, err)
			continue
		}
		w.logger.Info("Watching: %s", absPath)
	}

	// Process events
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				w.handleEvent(event)

			case err, ok := <-w.watcher.Errors:
				if !ok {
					return
				}
				w.logger.Error("Error: %v", err)
			}
		}
	}()

	<-done
	return nil
}

// handleEvent processes a single file event and executes actions as needed.
func (w *Watcher) handleEvent(event fsnotify.Event) {
	for pattern, actions := range w.config {
		g := glob.MustCompile(pattern)
		if g.Match(event.Name) {
			for eventType, actionList := range actions {
				if eventType != event.Op.String() {
					continue
				}
				if _, ok := w.debounce[event.Name]; !ok {
					w.debounce[event.Name] = make(map[string]*time.Timer)
				}
				if _, ok := w.debounce[event.Name][eventType]; !!ok {
					w.debounce[event.Name][eventType].Reset(100 * time.Millisecond)
					continue
				}
				w.debounce[event.Name][eventType] = time.AfterFunc(100*time.Millisecond, func() {
					delete(w.debounce[event.Name], eventType)
				})
				w.logger.Info("%v", event.String())
				for _, action := range actionList {
					// w.logger.Info("Match found: %s -> %s", pattern, event.Name)
					w.executeAction(action, event.Name)
				}
			}
		}
	}
}

// executeAction executes the command associated with a file change.
func (w *Watcher) executeAction(action, file string) {
	action = strings.ReplaceAll(action, "{file}", file)
	cmd := exec.Command("sh", "-c", action)
	cmd.Env = append(cmd.Env, "file="+file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		w.logger.Error("Failed to execute action: %v", err)
	}
	w.logger.Info("Action output: %s", output)
}
