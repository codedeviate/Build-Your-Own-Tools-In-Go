// analyzer.go
package lib

import (
	"bufio"
	"os"
	"regexp"
)

// LogEntry represents a parsed log entry.
type LogEntry struct {
	RemoteAddress string
	Client        string
	User          string
	Timestamp     string
	Message       string
}

// ProcessApacheAccessLog processes Apache2 access log entries.
func ProcessApacheAccessLog(line string) LogEntry {
	re := regexp.MustCompile(`([0-9a-f\.\:]+)\s(.*)\s(.*)\s\[(.*?)\] (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 6 {
		return LogEntry{
			RemoteAddress: matches[1],
			Client:        matches[2],
			User:          matches[3],
			Timestamp:     matches[4],
			Message:       matches[5],
		}
	}
	return LogEntry{}
}

// ProcessApacheErrorLog processes Apache2 error log entries.
func ProcessApacheErrorLog(line string) LogEntry {
	re := regexp.MustCompile(`\[(.*?)\] \[(.*?)\] \[(.*?)\] (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 5 {
		return LogEntry{Timestamp: matches[1], Message: matches[4]}
	}
	return LogEntry{}
}

// ProcessNginxAccessLog processes Nginx access log entries.
func ProcessNginxAccessLog(line string) LogEntry {
	re := regexp.MustCompile(`\[(.*?)\] (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 3 {
		return LogEntry{Timestamp: matches[1], Message: matches[2]}
	}
	return LogEntry{}
}

// ProcessNginxErrorLog processes Nginx error log entries.
func ProcessNginxErrorLog(line string) LogEntry {
	re := regexp.MustCompile(`\[(.*?)\] \[(.*?)\] (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 4 {
		return LogEntry{Timestamp: matches[1], Message: matches[3]}
	}
	return LogEntry{}
}

// ProcessMySQLLog processes MySQL/MariaDB log entries.
func ProcessMySQLLog(line string) LogEntry {
	re := regexp.MustCompile(`(\d{6} \d{2}:\d{2}:\d{2}) (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 3 {
		return LogEntry{Timestamp: matches[1], Message: matches[2]}
	}
	return LogEntry{}
}

// ReadFile reads the content of a file line by line.
func ReadFile(filePath string, processFunc func(string) LogEntry) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry := processFunc(line)
		if entry.Timestamp != "" {
			entries = append(entries, entry)
		}
	}
	return entries, scanner.Err()
}
