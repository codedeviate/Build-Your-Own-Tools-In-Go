// reverseproxy.go
package lib

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewProxy(target string) *Proxy {
	url, _ := url.Parse(target)
	return &Proxy{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = p.target.Host
	r.URL.Scheme = p.target.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = p.target.Host
	p.proxy.ServeHTTP(w, r)
}

func Start() {
	target := "http://localhost:8081" // Target server
	proxy := NewProxy(target)

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(proxy.handle)

	log.Println("Starting reverse proxy on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
