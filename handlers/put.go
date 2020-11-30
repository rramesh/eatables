package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// swagger:route PUT /items items updateItem
// Update an eatable item's details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update updates an item in the data store
func (items *Items) Update(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)
	items.l.Println("[Debug] Updating Item with ID ", it.ID)

	err := data.UpdateItem(it)
	if err == data.ErrItemNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		items.l.Println("[Error] Error Updating Item")
		items.l.Println("[Error]", err)
		return
	}

	items.l.Printf("[Debug] Updated Item: %#v", it)
}
