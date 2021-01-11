package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// Delete deletes an item in the DB identified by SKU UUID
//
// swagger:route DELETE /items/{sku} items deleteItem
// Deleta an eatable Item
// responses:
//	200: messageResponse
//  404: errorResponse
//  501: errorResponse
func (items *ItemHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	sku := getSKUFromRequest(r)
	items.l.Debug("Deleting Item", "SKU", sku)
	err := items.itemDB.DeleteItem(sku)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Error("Item not found", "ID", sku)
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Add("Content-Type", "application/json")
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	default:
		items.l.Error("Error Deleting", "ID", sku, "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	data.ToJSON(&GenericMessage{Message: "Item Deleted Successfully"}, rw)
	items.l.Debug("Deleted Item", "ID", sku)
}
