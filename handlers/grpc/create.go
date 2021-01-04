package handlers

import (
	"context"
	"strings"

	"github.com/rramesh/eatables/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protos "github.com/rramesh/eatables/protos/items"
)

// Add creates a new item to the data store via GRPC
func (it *ItemsGRPC) Add(ctx context.Context, req *protos.CreateOrUpdateRequest) (*protos.GenericResponse, error) {
	it.l.Debug("Add new item (GRPC)")
	item := data.ToItem(req)
	resp := &protos.GenericResponse{}
	errs := it.v.Validate(item)
	if len(errs) != 0 {
		return it.validationErrorGRPC(req, resp, errs.Errors())
	}
	err := it.itemDB.AddNewItem(*item)
	if err != nil {
		return it.validationErrorGRPC(req, resp, []string{err.Error()})
	}
	resp.Message = "Item Successfully Added"
	return resp, nil
}

func (it *ItemsGRPC) validationErrorGRPC(
	req *protos.CreateOrUpdateRequest,
	resp *protos.GenericResponse,
	errs []string,
) (*protos.GenericResponse, error) {
	errg := status.Newf(
		codes.FailedPrecondition,
		"Validation Error",
	)
	errg, cpe := errg.WithDetails(req)
	if cpe != nil {
		return nil, cpe
	}
	resp.Message = strings.Join(errs, ", ")
	return resp, errg.Err()
}
