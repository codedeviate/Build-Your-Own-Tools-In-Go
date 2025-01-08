// loadbalancer.go
package lib

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type ServerPool struct {
	servers []*url.URL
	current uint64
}

func (p *ServerPool) AddServer(server *url.URL) {
	p.servers = append(p.servers, server)
}

func (p *ServerPool) NextServer() *url.URL {
	next := atomic.AddUint64(&p.current, uint64(1))
	return p.servers[next%uint64(len(p.servers))]
}

func (p *ServerPool) LoadBalance(w http.ResponseWriter, r *http.Request) {
	server := p.NextServer()
	proxy := httputil.NewSingleHostReverseProxy(server)
	proxy.ServeHTTP(w, r)
}

var pool ServerPool

func Run() {
	servers := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	for _, server := range servers {
		url, err := url.Parse(server)
		if err != nil {
			log.Fatalf("Failed to parse server URL: %v", err)
		}
		pool.AddServer(url)
	}

	http.HandleFunc("/", pool.LoadBalance)
	fmt.Println("Load balancer started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
