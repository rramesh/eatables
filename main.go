package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/rramesh/eatables/handlers"
)

func main() {
	l := log.New(os.Stdout, "hello-api>", log.LstdFlags)

	ih := handlers.NewItems(l)

	sm := http.NewServeMux()

	sm.Handle("/", ih)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	l.Println("Starting servier on port", s.Addr)
	l.Println("Number of CPU Cores:", runtime.NumCPU())

	go func(server *http.Server) {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}(s)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, shutting down gracefully", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
