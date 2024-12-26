package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[GoMake] ", log.LstdFlags)

// Info logs an informational message.
func Info(msg string) {
	logger.Println(msg)
}

// Error logs an error message.
func Error(msg string) {
	logger.Println(msg)
}
