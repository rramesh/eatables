package handlers

import (
	"net/http"

	"github.com/google/uuid"
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
func (items *Items) Create(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)
	it.SKU = uuid.New().String()
	items.l.Debug("Inserting item", "item", it)
	err := items.itemDB.AddNewItem(it)
	rw.Header().Add("Content-Type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		items.l.Error("Error Adding Item", "error", err)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	data.ToJSON(&GenericMessage{Message: "Item Successfully Added"}, rw)
}
