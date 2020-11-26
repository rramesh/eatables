package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gorilla/mux"

	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/handlers"
)

func main() {
	l := log.New(os.Stdout, "hello-api>", log.LstdFlags)
	v := data.NewValidation()
	ih := handlers.NewItems(l, v)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/items", ih.ListAll)
	getRouter.HandleFunc("/items/{id:[0-9]+}", ih.ListSingle)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/items", ih.Create)
	postRouter.Use(ih.MiddlewareValidateItem)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/items", ih.UpdateItem)
	putRouter.Use(ih.MiddlewareValidateItem)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/items/{id:[0-9]+}", ih.Delete)

	// create a new server
	s := &http.Server{
		Addr:         ":9090",           //configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	l.Println("Starting servier on port", s.Addr)
	l.Println("Number of CPU Cores:", runtime.NumCPU())

	// start the server
	go func(server *http.Server) {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}(s)

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Recieved terminate, shutting down gracefully", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
