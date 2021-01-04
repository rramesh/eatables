package handlers

import (
	"context"

	"github.com/rramesh/eatables/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protos "github.com/rramesh/eatables/protos/items"
)

// Update updates an item in the data store via GRPC
func (it *ItemsGRPC) Update(ctx context.Context, req *protos.CreateOrUpdateRequest) (*protos.GenericResponse, error) {
	item := data.ToItem(req)
	it.l.Debug("Updating Item (GRPC)", "SKU", item.SKU)
	err := it.itemDB.UpdateItem(*item)
	if err == data.ErrItemNotFound {
		errg := status.Newf(
			codes.NotFound,
			"Item with SKU %s not found for Updating",
			item.SKU,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Info("Item not found", "SKU", item.SKU)
		return nil, errg.Err()
	}
	if err != nil {
		errg := status.Newf(
			codes.Internal,
			"Internal error updating Item",
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("Error Updating Item", "error", err)
		return nil, errg.Err()
	}
	return &protos.GenericResponse{Message: "Item Updated Successfully"}, nil
}
