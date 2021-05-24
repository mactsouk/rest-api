// Package main For the RESTful Server
//
// Documentation for REST API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.5
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// @termsOfService http://swagger.io/terms/

// User defines the structure for a Full User Record
//
// swagger:model user
type User struct {

	// The ID for the user
	// in: body
	//
	// required: false
	// min: 1
	ID int `json:"id"`
	// The Username of the user
	// in: body
	//
	// required: true
	Username string `json:"username"`
	// The Password of the user
	//
	// required: true
	Password string `json:"password"`
	// The Last Login time of the User
	//
	// required: true
	// min: 0
	LastLogin int64 `json:"lastlogin"`
	// Is the User Admin or not
	//
	// required: true
	Admin int `json:"admin"`
	// Is the User Logged In or Not
	//
	// required: true
	Active int `json:"active"`
}

// SliceToJSON encodes a slice with JSON records
func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(rw, r)
}

// swagger:route DELETE / Anything EMPTY
// Default Handler for everything that is not a match
//
// responses:
// 404: ErrorMessage

// DefaultHandler is for handling everything that is not a match
func DefaultHandler(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /*
// Default Handler for endpoints used with incorrect HTTP request method
//
// responses:
//	404: ErrorMessage

// MethodNotAllowedHandler is executed when the HTTP method is incorrect
func MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /time time
// Return current time
//
// responses:
//	200: OK

// TimeHandler is for handling /time – it works with plain text
func TimeHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route POST /add User createUser
// Create a new User
//
// responses:
//	200: OK
//  400: BadRequest

// AddHandler is for adding a new User
func AddHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route GET /getid User getUserInfo
// Returns the ID of a User given their username
//
// responses:
//	200: OK
//  400: BadRequest

// GetIDHandler returns the ID of an existing user
func GetIDHandler(rw http.ResponseWriter, r *http.Request) {
}

// swagger:route POST /login User getLoginInfo
// Login an existing user
//
// responses:
//	200: OK
//  400: BadRequest

// LoginHandler is for updating the LastLogin time of a user
// And changing the Active field to true
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
}

// Generic OK message returned as an HTTP Status Code
// swagger:response OK
type OK struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic error message returned as an HTTP Status Code
// swagger:response ErrorMessage
type ErrorMessage struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic BadRequest message returned as an HTTP Status Code
// swagger:response BadRequest
type BadRequest struct {
	// Description of the situation
	// in: body
	Body int
}

var PORT = ":4321"

func main() {
	mux := mux.NewRouter()

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	mux.Handle("/docs", s)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
