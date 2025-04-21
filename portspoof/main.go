// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"portspoof/config"
	"portspoof/dashboard"
	"portspoof/network"
	"syscall"
)

func main() {
	// Load settings from the external configuration file.
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	log.Println("Starting Portspoof Service...")

	// Start the real-time dashboard on a separate goroutine.
	go func() {
		if err := dashboard.StartDashboard(cfg.DashboardPort); err != nil {
			log.Fatalf("Dashboard failed: %v", err)
		}
	}()

	// Launch network listeners based on external configuration.
	go network.StartListeners(cfg)

	// Wait for termination signal (e.g., Ctrl+C).
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Shutting down Portspoof...")
}
