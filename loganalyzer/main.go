// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/username/loganalyzer/lib"
)

func main() {
	file := flag.String("file", "", "Path to the log file")
	logType := flag.String("type", "apache_access", "Log type (apache_access, apache_error, nginx_access, nginx_error, mysql)")
	flag.Parse()

	if *file == "" {
		fmt.Println("Usage: loganalyzer --file <log_file> --type <apache_access|apache_error|nginx_access|nginx_error|mysql>")
		os.Exit(1)
	}

	var processFunc func(string) lib.LogEntry
	switch *logType {
	case "apache_access":
		processFunc = lib.ProcessApacheAccessLog
	case "apache_error":
		processFunc = lib.ProcessApacheErrorLog
	case "nginx_access":
		processFunc = lib.ProcessNginxAccessLog
	case "nginx_error":
		processFunc = lib.ProcessNginxErrorLog
	case "mysql":
		processFunc = lib.ProcessMySQLLog
	default:
		fmt.Println("Unsupported log type. Use 'apache_access', 'apache_error', 'nginx_access', 'nginx_error', or 'mysql'.")
		os.Exit(1)
	}

	entries, err := lib.ReadFile(*file, processFunc)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		fmt.Printf("[%s] %s\n", entry.Timestamp, entry.Message)
		if entry.Client != "" && entry.Client != "-" {
			fmt.Printf("  Client: %s\n", entry.Client)
		}
		if entry.RemoteAddress != "" && entry.RemoteAddress != "-" {
			fmt.Printf("  RemoteAddress: %s\n", entry.RemoteAddress)
		}
		if entry.User != "" && entry.User != "-" {
			fmt.Printf("  User: %s\n", entry.User)
		}
	}
}
