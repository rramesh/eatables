package server

import (
	"github.com/hashicorp/go-hclog"
	"github.com/rramesh/eatables/data"
	"github.com/rramesh/eatables/handlers"
	protos "github.com/rramesh/eatables/protos/items"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer contains logger
type GRPCServer struct {
	l      hclog.Logger
	itemDB *data.ItemDB
}

// NewGRPCServer creates a new GRPC server
func NewGRPCServer(l hclog.Logger, idb *data.ItemDB) *GRPCServer {
	return &GRPCServer{l, idb}
}

// Server returns a GRPC server instance
func (g *GRPCServer) Server(v *data.Validation) *grpc.Server {
	gs := grpc.NewServer()
	its := handlers.NewItemsGRPC(g.l, v, g.itemDB)
	protos.RegisterItemsServer(gs, its)
	reflection.Register(gs)
	return gs
}
