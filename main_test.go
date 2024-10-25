// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// Initialize the routes for testing
	web.Router("/", &HomeController{})
	web.Router("/files/*file", &FileController{})
}

// Test for the home route
func TestHomeRoute(t *testing.T) {
	// Create a new HTTP request for the root route
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request
	web.BeeApp.Handlers.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// Check if the response body matches the expected output
	expected := "Welcome to the Home Page!"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %q, got %q", expected, rr.Body.String())
	}
}

// Test for the file route
func TestFileRoute(t *testing.T) {
	// Create a new HTTP request for the file route
	req, err := http.NewRequest("GET", "/files/test.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request
	web.BeeApp.Handlers.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// Check if the response body matches the expected output
	expected := "Requested file: test.txt"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %q, got %q", expected, rr.Body.String())
	}
}
