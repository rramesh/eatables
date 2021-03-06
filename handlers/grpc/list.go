package handlers

import (
	"context"

	"github.com/rramesh/eatables/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protos "github.com/rramesh/eatables/protos/items"
)

// ListAll returns all the item list
func (it *ItemsGRPC) ListAll(ctx context.Context, rr *protos.ListAllRequest) (*protos.ItemsListResponse, error) {
	it.l.Debug("Fetching Item List (GRPC)")
	itemList := it.itemDB.GetItems()
	return &protos.ItemsListResponse{Items: data.FromItems(itemList)}, nil
}

//ListByID returns specific item identified by ID
func (it *ItemsGRPC) ListByID(ctx context.Context, req *protos.IDRequest) (*protos.ItemsListResponse, error) {
	it.l.Debug("Fetch Item (GRPC)", "ID", req.Id)
	item, err := it.itemDB.GetItemByID(int(req.Id))
	if err == data.ErrItemNotFound {
		errg := status.Newf(
			codes.NotFound,
			"Item with ID %d not found",
			req.Id,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Info("Item not found", "ID", req.Id)
		return nil, errg.Err()
	}
	if err != nil {
		errg := status.Newf(
			codes.Internal,
			"An internal error occured, try again later",
			req.Id,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("An internal error occured", "error", err)
		return nil, errg.Err()
	}
	return &protos.ItemsListResponse{Items: data.FromItems(data.Items{item})}, nil
}

// ListBySKU returns item identifyed by SKU
func (it *ItemsGRPC) ListBySKU(ctx context.Context, req *protos.SKURequest) (*protos.ItemsListResponse, error) {
	sku := req.Sku
	it.l.Debug("Fetch Item (GRPC)", "SKU", sku)
	item, err := it.itemDB.GetItemBySKU(sku)
	if err == data.ErrItemNotFound {
		errg := status.Newf(
			codes.NotFound,
			"Item with SKU %s not found",
			sku,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Info("Item not found ", "SKU", sku)
		return nil, errg.Err()
	}
	if err == data.ErrInvalidUUID {
		errg := status.Newf(
			codes.InvalidArgument,
			"Invalid SKU. Should be a valid UUID format or value",
			sku,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		return nil, errg.Err()
	}
	if err != nil {
		errg := status.Newf(
			codes.Internal,
			"An internal error occured, try again later",
			sku,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("An internal error occured", "error", err)
		return nil, errg.Err()
	}
	return &protos.ItemsListResponse{Items: data.FromItems(data.Items{item})}, nil
}

// ListByVendorCode returns list of items by Vendor Code
func (it *ItemsGRPC) ListByVendorCode(ctx context.Context, req *protos.VendorCodeRequest) (*protos.ItemsListResponse, error) {
	vc := req.VendorCode
	it.l.Debug("Fetch Item (GRPC)", "Vendor Code", vc)
	items, err := it.itemDB.GetItemByVendorCode(vc)
	if err == data.ErrInvalidUUID {
		errg := status.Newf(
			codes.InvalidArgument,
			"Invalid Vendor Code. Should be a valid UUID format or value",
			vc,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		return nil, errg.Err()
	}
	if err != nil {
		errg := status.Newf(
			codes.Internal,
			"An internal error occured, try again later",
			vc,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Error("An internal error occured", "error", err)
		return nil, errg.Err()
	}
	return &protos.ItemsListResponse{Items: data.FromItems(items)}, nil
}
