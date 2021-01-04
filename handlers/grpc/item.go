package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/rramesh/eatables/data"
)

//ItemsGRPC holds logger
type ItemsGRPC struct {
	l      hclog.Logger
	v      *data.Validation
	itemDB *data.ItemDB
}

//NewItemsGRPC returns a new ItemsGRPC
func NewItemsGRPC(l hclog.Logger, v *data.Validation, idb *data.ItemDB) *ItemsGRPC {
	return &ItemsGRPC{l, v, idb}
}
