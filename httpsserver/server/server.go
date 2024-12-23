// server.go
package server

import (
	"fmt"
	"net/http"
)

// handler is a simple HTTP handler function.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, HTTPS World!")
}

// StartServer starts the HTTPS server on the specified port.
func StartServer(port, certFile, keyFile string) error {
    http.HandleFunc("/", handler)
    return http.ListenAndServeTLS(":"+port, certFile, keyFile, nil)
}
