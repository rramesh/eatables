package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the Server")

func main() {
	env.Parse()

	l := hclog.New(&hclog.LoggerOptions{
		Name:  "eatables",
		Level: hclog.LevelFromString("DEBUG"),
	})
	v := data.NewValidation()
	db := data.NewItemDB(l)

	l.Debug("Starting servier on port", "address", *bindAddress)
	l.Debug("Number of CPU Cores", "cores", runtime.NumCPU())

	ln, err := net.Listen("tcp", *bindAddress)

	if err != nil {
		l.Error("Error Starting Server", "error", err)
	}
	defer ln.Close()
	m := cmux.New(ln)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := server.NewGRPCServer(l, db)
	restS := server.NewRESTServer(l, db)

	g := grpcS.Server(v)
	h := restS.Server(v)

	go func(gServer *grpc.Server) {
		err := gServer.Serve(grpcL)
		if err != nil {
			l.Error("Error Starting Server", "error", err)
		}
	}(g)
	go func(hServer *http.Server) {
		err = hServer.Serve(httpL)
		if err != nil {
			l.Error("Error Starting Server", "error", err)
		}
	}(h)

	go m.Serve()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	l.Debug("Recieved terminate, shutting down gracefully", "Signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	g.GracefulStop()
	h.Shutdown(tc)
}
