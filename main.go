package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/rramesh/eatables/config"
	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		os.Exit(1)
	}
	l := conf.Logger
	dbh := data.NewDBHandle(l)
	err = dbh.Connect(conf)
	if err != nil {
		l.Error("Error connecting to Database", "error", err)
		os.Exit(1)
	}
	err = dbh.Init()
	if err != nil {
		l.Error("Error migrating DB", "error", err)
		os.Exit(1)
	}
	v := data.NewValidation()
	item := data.NewItemDB(l, dbh.DB)

	l.Info("Starting servier on port", "address", conf.BindAddress)
	l.Debug("Number of CPU Cores", "cores", runtime.NumCPU())

	ln, err := net.Listen("tcp", conf.BindAddress)

	if err != nil {
		l.Error("Error Starting Server", "error", err)
		os.Exit(1)
	}
	defer ln.Close()
	m := cmux.New(ln)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := server.NewGRPCServer(l, item)
	restS := server.NewRESTServer(l, item)

	g := grpcS.Server(v)
	h := restS.Server(v)

	go func(gServer *grpc.Server) {
		err := gServer.Serve(grpcL)
		if err != nil {
			l.Error("Error Starting Server", "error", err)
			os.Exit(1)
		}
	}(g)
	go func(hServer *http.Server) {
		err = hServer.Serve(httpL)
		if err != nil {
			l.Error("Error Starting Server", "error", err)
			os.Exit(1)
		}
	}(h)

	go m.Serve()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	l.Info("Recieved terminate, shutting down gracefully", "Signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	g.GracefulStop()
	h.Shutdown(tc)
}
