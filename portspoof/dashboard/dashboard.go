// dashboard/dashboard.go
package dashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// DashboardEvent represents a single event record.
type DashboardEvent struct {
	Time     time.Time `json:"time"`
	Service  string    `json:"service"`
	Message  string    `json:"message"`
	RemoteIP string    `json:"remote_ip"`
}

var (
	events []DashboardEvent
	mu     sync.Mutex
)

// AddEvent appends an event to the dashboard log.
func AddEvent(event DashboardEvent) {
	mu.Lock()
	defer mu.Unlock()
	events = append(events, event)
	if len(events) > 100 {
		events = events[len(events)-100:]
	}
}

// StartDashboard runs an HTTP server that returns recent events as JSON.
func StartDashboard(port int) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	})
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Dashboard available at http://localhost%s", addr)
	return http.ListenAndServe(addr, nil)
}
