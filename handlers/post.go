package handlers

import (
	"net/http"

	"github.com/rramesh/eatables/data"
)

//Create creates a new item to the data store
func (items *Items) Create(rw http.ResponseWriter, r *http.Request) {
	it := r.Context().Value(KeyItem{}).(data.Item)

	items.l.Printf("[DEBUG] Inserting item: %#v\n", it)
	data.AddNewItem(it)
}
