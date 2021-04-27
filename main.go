package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	arguments := os.Args
	if len(arguments) >= 2 {
		PORT = ":" + arguments[1]
	}

	// Create a new ServeMux using Gorilla
	mux := mux.NewRouter()

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	mux.NotFoundHandler = http.HandlerFunc(DefaultHandler)

	notAllowed := notAllowedHandler{}
	mux.MethodNotAllowedHandler = notAllowed

	go func() {
		log.Println("Listening to", PORT)
		err := s.ListenAndServe()
		if err != nil {
		  log.Printf("Error starting server: %s\n", err)
		  return
		}
	  }()
	
	  sigs := make(chan os.Signal, 1)
	  signal.Notify(sigs, os.Interrupt)
	  sig := <-sigs
	  log.Println("Quitting after signal:", sig)
	  time.Sleep(5 * time.Second)
	  s.Shutdown(nil)
	}

}
