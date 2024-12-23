// server.go
package server

import (
	"fmt"
	"net/http"
)

// handler is a simple HTTP handler function.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

// StartServer starts the HTTP server on the specified port.
func StartServer(port string) error {
    http.HandleFunc("/", handler)
    return http.ListenAndServe(":"+port, nil)
}
