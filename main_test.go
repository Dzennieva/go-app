// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(homePage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", rr.Code, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("unexpected content type: got %s, want %s", contentType, expectedContentType)
	}
}
