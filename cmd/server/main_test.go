package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestHandlerEndpoints(t *testing.T) {
	app := &application{}
	handler := app.handler()

	tests := []struct {
		name         string
		path         string
		wantStatus   int
		wantBodyPart string
	}{
		{name: "hello", path: "/hello", wantStatus: http.StatusOK, wantBodyPart: "Hello from hackathon-demo!"},
		{name: "health", path: "/healthz", wantStatus: http.StatusOK, wantBodyPart: "ok"},
		{name: "unknown route", path: "/missing", wantStatus: http.StatusNotFound, wantBodyPart: "404 page not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			recorder := httptest.NewRecorder()

			handler.ServeHTTP(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", recorder.Code, tt.wantStatus)
			}
			if !strings.Contains(recorder.Body.String(), tt.wantBodyPart) {
				t.Fatalf("body = %q, want it to contain %q", recorder.Body.String(), tt.wantBodyPart)
			}
		})
	}
}

func TestMetricsCountEveryRequest(t *testing.T) {
	app := &application{}
	handler := app.handler()

	for _, path := range []string{"/hello", "/missing", "/metrics"} {
		handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, path, nil))
	}

	metrics := httptest.NewRecorder()
	handler.ServeHTTP(metrics, httptest.NewRequest(http.MethodGet, "/metrics", nil))
	if metrics.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", metrics.Code, http.StatusOK)
	}
	if !strings.Contains(metrics.Body.String(), "# HELP hackathon_demo_http_requests_total") ||
		!strings.Contains(metrics.Body.String(), "# TYPE hackathon_demo_http_requests_total counter") {
		t.Fatalf("metrics exposition missing HELP/TYPE lines:\n%s", metrics.Body.String())
	}
	if !strings.Contains(metrics.Body.String(), "hackathon_demo_http_requests_total 4\n") {
		t.Fatalf("metrics = %q, want counter value 4", metrics.Body.String())
	}
}

func TestRequestCounterIsConcurrentSafe(t *testing.T) {
	app := &application{}
	handler := app.handler()
	const requests = 100

	var group sync.WaitGroup
	group.Add(requests)
	for range requests {
		go func() {
			defer group.Done()
			handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/hello", nil))
		}()
	}
	group.Wait()

	metrics := httptest.NewRecorder()
	handler.ServeHTTP(metrics, httptest.NewRequest(http.MethodGet, "/metrics", nil))
	want := "hackathon_demo_http_requests_total " + strconv.Itoa(requests+1) + "\n"
	if !strings.Contains(metrics.Body.String(), want) {
		t.Fatalf("metrics = %q, want %q", metrics.Body.String(), want)
	}
}
