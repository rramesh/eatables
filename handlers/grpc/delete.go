package handlers

import (
	"context"

	"github.com/rramesh/eatables/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protos "github.com/rramesh/eatables/protos/items"
)

// Delete deletes an item in the DB identified by ID by GRPC
func (it *ItemsGRPC) Delete(ctx context.Context, req *protos.SKURequest) (*protos.GenericResponse, error) {
	sku := req.Sku
	it.l.Debug("Deleting Item (GRPC)", "ID", sku)
	err := it.itemDB.DeleteItem(sku)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		errg := status.Newf(
			codes.NotFound,
			"Item with SKU %s not found for Deletion",
			sku,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("Could not find item", "SKU", sku)
		return nil, errg.Err()
	default:
		errg := status.Newf(
			codes.Internal,
			"Internal error deleting Item",
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("Error Deleting Item", "SKU", sku)
		return nil, errg.Err()
	}
	it.l.Debug("Deleted Item", "ID", sku)
	return &protos.GenericResponse{Message: "Item Deleted Successfully"}, nil
}
