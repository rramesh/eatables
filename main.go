package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/nicholasjackson/env"
	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the Server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "eatables>", log.LstdFlags)
	v := data.NewValidation()

	l.Println("Starting servier on port", *bindAddress)
	l.Println("Number of CPU Cores:", runtime.NumCPU())

	ln, err := net.Listen("tcp", *bindAddress)

	if err != nil {
		l.Fatal(err)
	}

	m := cmux.New(ln)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := server.NewGRPCServer(l)
	restS := server.NewRESTServer(l)

	g := grpcS.Server(v)
	h := restS.Server(v)

	go func(gServer *grpc.Server) {
		err := gServer.Serve(grpcL)
		if err != nil {
			l.Fatal(err)
		}
	}(g)
	go func(hServer *http.Server) {
		err = hServer.Serve(httpL)
		if err != nil {
			l.Fatal(err)
		}
	}(h)

	m.Serve()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Recieved terminate, shutting down gracefully", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	g.Stop()
	h.Shutdown(tc)
}
