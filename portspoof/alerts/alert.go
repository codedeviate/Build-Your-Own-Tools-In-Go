// alerts/alert.go
package alerts

import (
	"fmt"
	"portspoof/logging"
	"sync"
)

var (
	ipActivity = make(map[string]int)
	mu         sync.Mutex
	threshold  = 5 // Default; can be overridden from config.
)

// SetThreshold configures the threshold for suspicious activity.
func SetThreshold(t int) {
	mu.Lock()
	defer mu.Unlock()
	threshold = t
}

// RecordActivity increments the connection count for an IP.
func RecordActivity(ip string) {
	mu.Lock()
	defer mu.Unlock()

	ipActivity[ip]++
	if ipActivity[ip] >= threshold {
		triggerAlert(ip, ipActivity[ip])
		// Reset after issuing an alert.
		ipActivity[ip] = 0
	}
}

// triggerAlert logs and prints an alert message.
func triggerAlert(ip string, count int) {
	alertMsg := fmt.Sprintf("Suspicious activity from IP %s: %d attempts", ip, count)
	logging.LogEvent("alerts", alertMsg)
	fmt.Println("ALERT:", alertMsg)
}
