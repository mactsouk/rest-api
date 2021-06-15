package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTimeHanlderV1(t *testing.T) {
	req, err := http.NewRequest("GET", "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TimeHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestMethodNotAllowedV1(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MethodNotAllowedHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
