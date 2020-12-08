package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/rramesh/eatables/data"
	protos "github.com/rramesh/eatables/protos/items"
	"golang.org/x/net/context"
)

//ItemsGRPC holds logger
type ItemsGRPC struct {
	l *log.Logger
	v *data.Validation
}

//NewItemsGRPC returns a new ItemsGRPC
func NewItemsGRPC(l *log.Logger, v *data.Validation) *ItemsGRPC {
	return &ItemsGRPC{l, v}
}

// ListAll returns all the item list
func (it *ItemsGRPC) ListAll(ctx context.Context, rr *protos.ListAllRequest) (*protos.ItemsListResponse, error) {
	it.l.Println("[Debug] Fetching Item List (GRPC)")
	itemList := data.GetItems()
	return &protos.ItemsListResponse{Items: data.FromItems(itemList)}, nil
}

//ListByID returns specific item identified by ID
func (it *ItemsGRPC) ListByID(ctx context.Context, req *protos.IDRequest) (*protos.ItemsListResponse, error) {
	it.l.Println("[Debug] Fetch Item (GRPC) by ID", req.Id)
	item, err := data.GetItemByID(int(req.Id))
	if err == data.ErrItemNotFound {
		it.l.Println("[Info] Item with ID ", req.Id, " not found")
		return nil, err
	}
	return &protos.ItemsListResponse{Items: data.FromItems(data.Items{item})}, nil
}

// ListBySKU returns item identifyed by SKU
func (it *ItemsGRPC) ListBySKU(ctx context.Context, req *protos.UUIDRequest) (*protos.ItemsListResponse, error) {
	it.l.Println("[Debug] Fetch Item (GRPC) by SKU", req.Uuid)
	item, err := data.GetItemBySKU(req.Uuid)
	if err == data.ErrItemNotFound {
		it.l.Println("[Info] Item with SKU ", req.Uuid, " not found")
		return nil, err
	}
	return &protos.ItemsListResponse{Items: data.FromItems(data.Items{item})}, nil
}

// ListByVendorCode returns list of items by Vendor Code
func (it *ItemsGRPC) ListByVendorCode(ctx context.Context, req *protos.UUIDRequest) (*protos.ItemsListResponse, error) {
	it.l.Println("[Debug] Fetch Item (GRPC) by Vendor Code", req.Uuid)
	items, err := data.GetItemByVendorCode(req.Uuid)
	if err == data.ErrItemNotFound {
		it.l.Println("[Info] Item with Vendor Code ", req.Uuid, " not found")
		return nil, err
	}
	return &protos.ItemsListResponse{Items: data.FromItems(items)}, nil
}

// Add adds a new item to the DB
func (it *ItemsGRPC) Add(ctx context.Context, req *protos.CreateOrUpdateRequest) (*protos.GenericResponse, error) {
	it.l.Println("[Debug] Add new item (GRPC)")
	item := data.ToItem(req)
	resp := &protos.GenericResponse{}
	errs := it.v.Validate(item)
	if len(errs) != 0 {
		resp.Message = strings.Join(errs.Errors(), ", ")
		return resp, fmt.Errorf("Validationn Error")
	}
	data.AddNewItem(*item)
	resp.Message = "Item Successfully Added"
	return resp, nil
}

// Update updates existing item using SKU
func (it *ItemsGRPC) Update(ctx context.Context, req *protos.CreateOrUpdateRequest) (*protos.GenericResponse, error) {
	item := data.ToItem(req)
	it.l.Println("[Debug] Updating Item with SKU (GRPC)", item.SKU)
	err := data.UpdateItem(*item)
	if err == data.ErrItemNotFound {
		it.l.Println("[Info] Item with SKU ", item.SKU, " not found")
		return nil, fmt.Errorf("Item with given SKU not found for Update")
	}
	if err != nil {
		it.l.Println("[Error] Error Updating Item", err)
		return nil, fmt.Errorf("Error Updating Item")
	}
	return &protos.GenericResponse{Message: "Item Updated Successfully"}, nil
}

// Delete deletes an item in the list identified by ID
func (it *ItemsGRPC) Delete(ctx context.Context, req *protos.IDRequest) (*protos.GenericResponse, error) {
	id := req.Id
	it.l.Println("[Debug] Deleting Item (GRPC) with ID", id)
	err := data.DeleteItem(int(id))
	switch err {
	case nil:
	case data.ErrItemNotFound:
		it.l.Println("[Error] Could not find item by ID ", id)
		return nil, fmt.Errorf("Item with given ID not found for Delete")
	default:
		it.l.Println("[Error] Error Deleting Item with ID ", id)
		return nil, fmt.Errorf("Error Deleting Item")
	}
	it.l.Println("[Debug] Deleted Item with ID ", id)
	return &protos.GenericResponse{Message: "Item Deleted Successfully"}, nil
}
