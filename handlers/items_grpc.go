package handlers

import (
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/rramesh/eatables/data"
	protos "github.com/rramesh/eatables/protos/items"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return &protos.ItemsListResponse{Items: data.FromItems(data.Items{item})}, nil
}

// ListBySKU returns item identifyed by SKU
func (it *ItemsGRPC) ListBySKU(ctx context.Context, req *protos.UUIDRequest) (*protos.ItemsListResponse, error) {
	it.l.Debug("Fetch Item (GRPC)", "SKU", req.Uuid)
	item, err := it.itemDB.GetItemBySKU(req.Uuid)
	if err == data.ErrItemNotFound {
		errg := status.Newf(
			codes.NotFound,
			"Item with SKU %s not found",
			req.Uuid,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Info("Item not found ", "SKU", req.Uuid)
		return nil, errg.Err()
	}
	return &protos.ItemsListResponse{Items: data.FromItems(data.Items{item})}, nil
}

// ListByVendorCode returns list of items by Vendor Code
func (it *ItemsGRPC) ListByVendorCode(ctx context.Context, req *protos.UUIDRequest) (*protos.ItemsListResponse, error) {
	it.l.Debug("Fetch Item (GRPC)", "Vendor Code", req.Uuid)
	items, err := it.itemDB.GetItemByVendorCode(req.Uuid)
	if err == data.ErrItemNotFound {
		errg := status.Newf(
			codes.NotFound,
			"Item with Vendor Code %s not found",
			req.Uuid,
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		it.l.Info("Item not found", "Vendor Code", req.Uuid)
		return nil, errg.Err()
	}
	return &protos.ItemsListResponse{Items: data.FromItems(items)}, nil
}

// Add adds a new item to the DB
func (it *ItemsGRPC) Add(ctx context.Context, req *protos.CreateOrUpdateRequest) (*protos.GenericResponse, error) {
	it.l.Debug("Add new item (GRPC)")
	item := data.ToItem(req)
	item.SKU = uuid.New().String()
	resp := &protos.GenericResponse{}
	errs := it.v.Validate(item)
	if len(errs) != 0 {
		errg := status.Newf(
			codes.FailedPrecondition,
			"Validation Error",
		)
		errg, cpe := errg.WithDetails(req)
		if cpe != nil {
			return nil, cpe
		}
		resp.Message = strings.Join(errs.Errors(), ", ")
		return resp, errg.Err()
	}
	it.itemDB.AddNewItem(*item)
	resp.Message = "Item Successfully Added"
	return resp, nil
}

// Update updates existing item using SKU
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

// Delete deletes an item in the list identified by ID
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
