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
//	200: itemResponse
//  422: errorValidation
//  501: errorResponse
func (items *Items) Create(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)
	rw.Header().Add("Content-Type", "application/json")
	items.l.Debug("Inserting item", "item", it)
	items.itemDB.AddNewItem(it)
}
