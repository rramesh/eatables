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
//	200: messageResponse
//  404: errorResponse
//  501: errorResponse
func (items *Items) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getItemID(r)
	items.l.Debug("Deleting Item", "ID", id)
	err := items.itemDB.DeleteItem(id)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Error("Item not found", "ID", id)
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Add("Content-Type", "application/json")
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	default:
		items.l.Error("Error Deleting", "ID", id, "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	data.ToJSON(&GenericMessage{Message: "Item Deleted Successfully"}, rw)
	items.l.Debug("Deleted Item", "ID", id)
}
