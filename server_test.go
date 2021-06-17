package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTimeHanlder(t *testing.T) {
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
func TestMethodNotAllowed(t *testing.T) {
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

func TestLogin(t *testing.T) {
	UserPass := []byte(`{"user": "admin", "password": "admin"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

func TestLogout(t *testing.T) {
	UserPass := []byte(`{"user": "admin", "password": "admin"}`)
	req, err := http.NewRequest("POST", "/logout", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LogoutHandler)
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

func TestAdd(t *testing.T) {
	UserPass := []byte(`[{"user": "admin", "password": "admin"}, {"user": "test", "password": "myPass"}]`)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}

func TestGetUserDataHandler(t *testing.T) {
	UserPass := []byte(`{"username": "admin", "password": "admin"}`)
	req, err := http.NewRequest("GET", "/username/1", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserDataHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	expected := `[{"id":1,"user":"admin","password":"admin","lastlogin":0,"admin":1,"active":0}]`
	serverResponse := rr.Body.String()
	result := strings.Split(serverResponse, "lastlogin")
	serverResponse = result[0] + `lastlogin":0,"admin":1,"active":0}]`

	if serverResponse != expected {
		t.Errorf("handler returned unexpected body: got %v but wanted %v",
			rr.Body.String(), expected)
	}
}
