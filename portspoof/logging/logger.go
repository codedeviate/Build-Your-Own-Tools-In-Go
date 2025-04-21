// logging/logger.go
package logging

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	loggers = make(map[string]*log.Logger)
	mu      sync.Mutex
)

// getLogger returns (or creates) a logger for a given service.
func getLogger(service string) *log.Logger {
	mu.Lock()
	defer mu.Unlock()

	if logger, exists := loggers[service]; exists {
		return logger
	}

	// Create or append to a log file named after the service.
	fileName := fmt.Sprintf("%s.log", service)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to open log file %s: %v", fileName, err)
	}
	logger := log.New(file, fmt.Sprintf("[%s] ", service), log.LstdFlags)
	loggers[service] = logger
	return logger
}

// LogEvent writes an event string to the appropriate log file.
func LogEvent(service, event string) {
	logger := getLogger(service)
	logger.Printf("%s - %s", time.Now().Format(time.RFC3339), event)
}
