package internal

import (
	"io"
	"log"
)

// Logger is a simple wrapper around log for structured logging.
type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

// NewLogger creates a new Logger.
func NewLogger(output io.Writer) *Logger {
	return &Logger{
		infoLog:  log.New(output, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(output, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs informational messages.
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLog.Printf(format, v...)
}

// Error logs error messages.
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLog.Printf(format, v...)
}
