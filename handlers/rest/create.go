package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// Create creates a new item to the data store
//
// swagger:route POST /items items createItem
// Create a new Eatable item
// responses:
//	200: messageResponse
//  422: errorValidation
//  501: errorResponse
func (items *ItemHandler) Create(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)
	items.l.Debug("Inserting item", "item", it)
	err := items.itemDB.AddNewItem(it)
	rw.Header().Add("Content-Type", "application/json")
	if err != nil {
		items.validationErrorResponse(rw, err)
		return
	}
	rw.WriteHeader(http.StatusOK)
	data.ToJSON(&GenericMessage{Message: "Item Successfully Added"}, rw)
}

func (items *ItemHandler) validationErrorResponse(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusUnprocessableEntity)
	items.l.Error("Error Adding Item", "error", err)
	data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
}
