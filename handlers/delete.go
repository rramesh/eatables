package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// Delete returns a specific item with ID passed in the URL
//
// swagger:route DELETE /items/{id} items deleteItem
// Deleta an eatable Item
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse
func (items *Items) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getItemID(r)
	items.l.Println("[Debug] Deleting Item with ID ", id)
	err := data.DeleteItem(id)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Println("[Error] Could not find item by ID ", id)
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Add("Content-Type", "application/json")
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		items.l.Println("[Error] Error Deleting Item with ID ", id)
		rw.WriteHeader(http.StatusInternalServerError)
		items.l.Println("[Error]", err)
		return
	}
	items.l.Println("[Debug] Deleted Item with ID ", id)
}
