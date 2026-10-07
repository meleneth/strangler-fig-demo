package app

import (
    "net/http/httptest"
    "testing"
)

func TestHealth(t *testing.T) {
    response := httptest.NewRecorder()
    Handler(Dependencies{}).ServeHTTP(response, httptest.NewRequest("GET", "/health", nil))
    if response.Code != 200 || response.Body.String() != `{"status":"ok"}` || response.Header().Get("Content-Type") != "application/json" {
        t.Fatalf("health response: %d %s %v", response.Code, response.Body, response.Header())
    }
}
