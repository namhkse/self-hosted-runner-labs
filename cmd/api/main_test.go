package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		statusCode int
		body       string
	}{
		{name: "home", path: "/", statusCode: http.StatusOK, body: "Hello, world!\n"},
		{name: "health", path: "/healthz", statusCode: http.StatusOK, body: "ok\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, tt.path, nil)

			newMux().ServeHTTP(response, request)

			if response.Code != tt.statusCode {
				t.Errorf("status code = %d, want %d", response.Code, tt.statusCode)
			}
			if response.Body.String() != tt.body {
				t.Errorf("body = %q, want %q", response.Body.String(), tt.body)
			}
		})
	}
}
