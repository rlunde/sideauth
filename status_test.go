package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(status)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var s Status
	err = json.NewDecoder(rr.Body).Decode(&s)
	if err != nil {
		t.Errorf("handler returned unexpected body: wanted a Status struct, but got error %s with content  %v",
			err.Error(), rr.Body.String())
	}
	// Check the response body is what we expect.
	expected := Status{
		Database:       "unknown",
		Uptime:         "100h",
		Duration:       "6µs",
		ServiceVersion: "0.1.0",
	}
	if s.Database != expected.Database {
		t.Errorf("handler returned unexpected database status: got %v want %v",
			s.Database, expected.Database)
	}
	if s.ServiceVersion != expected.ServiceVersion {
		t.Errorf("handler returned unexpected service version: got %v want %v",
			s.ServiceVersion, expected.ServiceVersion)
	}
	// TODO: figure out what to do to verify uptime and duration
}
