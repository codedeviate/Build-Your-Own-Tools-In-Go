// netstat.go
package netstat

import (
	"bufio"
	"os/exec"
	"strings"
)

// Netstat retrieves the list of active network connections.
func Netstat() ([]string, error) {
    cmd := exec.Command("netstat", "-an")
    stdout, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(strings.NewReader(string(stdout)))
    var connections []string
    for scanner.Scan() {
        connections = append(connections, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return connections, nil
}
