package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	// Create a request to pass to our handler. TODO: figure out what to use instead of nil for 3rd parameter
	req, err := http.NewRequest("POST", "/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var account Account
	err = json.NewDecoder(rr.Body).Decode(&account)
	if err != nil {
		t.Errorf("handler returned unexpected body: wanted an Account struct, but got error %s with content  %v",
			err.Error(), rr.Body.String())
	}
	// Check the response body is what we expect.
	expected := Account{
		Account: "testAccount",
		Pwhash:  "xyzzy",
		Email:   "test@example.com",
	}
	if account.Account != expected.Account {
		t.Errorf("handler returned unexpected account name: got %v want %v",
			account.Account, expected.Account)
	}

	// TODO: finish this up
}
