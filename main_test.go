// Test the main function

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	testPaths := []string{"/home", "/"}

	for _, path := range testPaths {
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(homePage)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("path %s returned wrong status code: got %v want %v",
				path, status, http.StatusOK)
		}

		// Verify content type
		expected := "text/html; charset=utf-8"
		if contentType := rr.Header().Get("Content-Type"); contentType != expected {
			t.Errorf("path %s returned unexpected content type: got %v want %v",
				path, contentType, expected)
		}
	}
}
