package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// Update updates an item in the data store
//
// swagger:route PUT /items items updateItem
// Update an eatable item's details
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation
func (items *Items) Update(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)
	items.l.Debug("Updating Item", "SKU", it.SKU)

	err := items.itemDB.UpdateItem(it)
	if err == data.ErrItemNotFound {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		items.l.Error("Error Updating Item", "error", err)
		return
	}

	items.l.Debug("Item updated successfully", "SKU", it.SKU)
}
