// Command server runs the beginner HTTP and metrics demo.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

// application holds the state shared by the HTTP handlers.
type application struct {
	requests atomic.Uint64
}

func main() {
	app := &application{}
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	server := &http.Server{
		Addr:    addr,
		Handler: app.handler(),
	}

	log.Printf("listening on http://0.0.0.0%s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func (app *application) handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", app.hello)
	mux.HandleFunc("GET /healthz", app.healthz)
	mux.HandleFunc("GET /metrics", app.metrics)
	return app.countRequests(mux)
}

func (app *application) countRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.requests.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (app *application) hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintln(w, "Hello from hackathon-demo!")
}

func (app *application) healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintln(w, "ok")
}

func (app *application) metrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
	_, _ = fmt.Fprintln(w, "# HELP hackathon_demo_http_requests_total Total HTTP requests handled by this server, including /metrics and 404 responses.")
	_, _ = fmt.Fprintln(w, "# TYPE hackathon_demo_http_requests_total counter")
	_, _ = fmt.Fprintf(w, "hackathon_demo_http_requests_total %d\n", app.requests.Load())
}
