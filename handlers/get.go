package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// ListAll returns all the items from the data store
func (items *Items) ListAll(rw http.ResponseWriter, r *http.Request) {
	items.l.Println("[Debug] Fetchingn Item List")

	itemList := data.GetItems()

	err := data.ToJSON(itemList, rw)
	if err != nil {
		items.l.Println("[Error] seralizing item list to JSON")
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericErrorMessage{Message: err.Error()}, rw)
	}
}

// ListSingle returns a specific item with ID passed in the URL
func (items *Items) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getItemID(r)

	item, err := data.GetItemByID(id)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Println("[Error] Could not find item by ID ", id)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericErrorMessage{Message: err.Error()}, rw)
		return
	default:
		items.l.Println("[Error] Error Fetching Item with ID ", id)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericErrorMessage{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(item, rw)
	if err != nil {
		items.l.Println("[Error] seralizing item to JSON")
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericErrorMessage{Message: err.Error()}, rw)
	}
}
