package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

// ListAll returns all the items from the data store
//
// swagger:route GET /items items listItems
// Returns a list of Eatable Items
// responses:
//	200: itemResponse
func (items *Items) ListAll(rw http.ResponseWriter, r *http.Request) {
	items.l.Debug("Fetching Item List")
	rw.Header().Add("Content-Type", "application/json")
	itemList := items.itemDB.GetItems()
	err := data.ToJSON(itemsResponse(itemList), rw)
	if err != nil {
		items.l.Error("Seralizing item list to JSON", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
	}
}

// ListSingle returns a specific item with ID passed in the URL
//
// swagger:route GET /items/{id} items listSingleItem
// Return a specific Eatable item from the database
// responses:
//	200: itemResponse
//	404: errorResponse
func (items *Items) ListSingle(rw http.ResponseWriter, r *http.Request) {
	items.l.Debug("Fetching Item List")
	id := getItemID(r)
	rw.Header().Add("Content-Type", "application/json")
	item, err := items.itemDB.GetItemByID(id)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Info("Could not find", "ID", id)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	default:
		items.l.Error("Error Fetching Item", "ID", id, "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(singleItemResponse(item), rw)
	if err != nil {
		items.l.Error("Seralizing item to JSON", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
	}
}

// ListItemBySKU returns an item by SKU (UUID)
//
// swagger:route GET /items/sku/{uuid} items listItemBySKU
// Returns an item by SKU (UUID)
// responses:
//	200: itemResponse
//	404: errorResponse
func (items *Items) ListItemBySKU(rw http.ResponseWriter, r *http.Request) {
	items.l.Debug("Fetching Item List by SKU")
	uuid := getUUID(r)
	rw.Header().Add("Content-Type", "application/json")
	item, err := items.itemDB.GetItemBySKU(uuid)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Info("Item not found", "SKU", uuid)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	default:
		items.l.Error("Error Fetching Item", "SKU", uuid, "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(singleItemResponse(item), rw)
	if err != nil {
		items.l.Error("Seralizing items to JSON", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
	}
}

// ListItemsByVendor returns a list of items by Vendor Code (UUID)
//
// swagger:route GET /items/vendor/{uuid} items listItemsByVendor
// Returns a list of items by Vendor Code (UUID)
// responses:
//	200: itemResponse
//	404: errorResponse
func (items *Items) ListItemsByVendor(rw http.ResponseWriter, r *http.Request) {
	items.l.Debug("Fetching Item List by Vendor Code")
	uuid := getUUID(r)
	rw.Header().Add("Content-Type", "application/json")
	itemList, err := items.itemDB.GetItemByVendorCode(uuid)
	switch err {
	case nil:
	case data.ErrItemNotFound:
		items.l.Info("Item not found", "Vendor Code", uuid)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	default:
		items.l.Error("Error Fetching Items", "Vendor Code", uuid, "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(itemsResponse(itemList), rw)
	if err != nil {
		items.l.Error("Seralizing items to JSON", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericMessage{Message: err.Error()}, rw)
	}
}

func singleItemResponse(item *data.Item) map[string]data.Items {
	var itemResp data.Items = make(data.Items, 1)
	itemResp[0] = item
	return map[string]data.Items{"items": itemResp}
}

func itemsResponse(items data.Items) map[string]data.Items {
	return map[string]data.Items{"items": items}
}
