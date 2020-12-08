package server

import (
	"log"

	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/handlers"
	protos "github.com/rramesh/eatables/protos/items"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer contains logger
type GRPCServer struct {
	l *log.Logger
}

// NewGRPCServer creates a new GRPC server
func NewGRPCServer(l *log.Logger) *GRPCServer {
	return &GRPCServer{l}
}

// Server returns a GRPC server instance
func (g *GRPCServer) Server(v *data.Validation) *grpc.Server {
	gs := grpc.NewServer()
	its := handlers.NewItemsGRPC(g.l, v)
	protos.RegisterItemsServer(gs, its)
	reflection.Register(gs)
	return gs
}
