package main

import (
	"os"

	"github.com/username/watcher/internal"
)

func main() {
	// Set up logger
	logger := internal.NewLogger(os.Stdout)

	// Validate arguments
	if len(os.Args) < 3 {
		logger.Error("Usage: main <config.json> <path1> <path2> ...")
		os.Exit(1)
	}

	// Load configuration
	configFile := os.Args[1]
	config, err := internal.LoadConfig(configFile)
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		os.Exit(1)
	}

	// Initialize watcher
	paths := os.Args[2:]
	watcher, err := internal.NewWatcher(paths, config, logger)
	if err != nil {
		logger.Error("Failed to create watcher: %v", err)
		os.Exit(1)
	}

	// Start monitoring
	if err := watcher.Start(); err != nil {
		logger.Error("Watcher encountered an error: %v", err)
		os.Exit(1)
	}
}
