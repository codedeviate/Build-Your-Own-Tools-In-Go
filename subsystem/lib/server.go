package lib

import (
	"fmt"
	"net/http"
	"strings"
)

// GetServerInfo fetches the headers of the given URL and tries to identify the server and its version.
func GetServerInfo(url string) {
	// Ensure the URL starts with http:// or https://
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Headers for %s:\n", url)
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	// Try to extract specific information
	server := resp.Header.Get("Server")
	if server != "" {
		fmt.Printf("\nDetected Server: %s\n", server)
	} else {
		fmt.Println("\nServer information not found in headers.")
	}

	xPoweredBy := resp.Header.Get("X-Powered-By")
	if xPoweredBy != "" {
		fmt.Printf("Detected Subsystem (e.g., PHP, Python): %s\n", xPoweredBy)
	} else {
		fmt.Println("Subsystem information not found in headers.")
	}
}
