package handlers

import (
	"context"

	"github.com/rramesh/eatables/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protos "github.com/rramesh/eatables/protos/items"
)

// Delete deletes an item in the DB identified by ID by GRPC
func (it *ItemsGRPC) Delete(ctx context.Context, req *protos.IDRequest) (*protos.GenericResponse, error) {
	id := req.Id
	it.l.Debug("Deleting Item (GRPC)", "ID", id)
	err := it.itemDB.DeleteItem(int(id))
	switch err {
	case nil:
	case data.ErrItemNotFound:
		errg := status.Newf(
			codes.NotFound,
			"Item with ID %d not found for Deletion",
			id,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("Could not find item", "ID", id)
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
		it.l.Error("Error Deleting Item", "ID", id)
		return nil, errg.Err()
	}
	it.l.Debug("Deleted Item", "ID", id)
	return &protos.GenericResponse{Message: "Item Deleted Successfully"}, nil
}
