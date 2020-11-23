package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rramesh/eatables/data"
)

// Items is a http.Handler
type Items struct {
	l *log.Logger
}

//NewItems creates a items handler with the given logger
func NewItems(l *log.Logger) *Items {
	return &Items{l}
}

//GetItems returns the items from the data store
func (items *Items) GetItems(rw http.ResponseWriter, r *http.Request) {
	items.l.Println("Handle Get Items")
	itemList := data.GetItems()
	err := itemList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to parse list of items as JSON", http.StatusInternalServerError)
	}
}

//AddItem adds a new item to the data store
func (items *Items) AddItem(rw http.ResponseWriter, r *http.Request) {
	items.l.Println("Handle POST Item")
	it := r.Context().Value(KeyItem{}).(data.Item)
	data.AddItem(&it)
	items.l.Printf("Prod: %#v", it)
}

// UpdateItem updates an item in the data store
func (items *Items) UpdateItem(rw http.ResponseWriter, r *http.Request) {
	items.l.Println("Handle PUT Item")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "ID is not a number", http.StatusBadRequest)
		return
	}

	it := r.Context().Value(KeyItem{}).(data.Item)
	err = data.UpdateItem(id, &it)
	if err == data.ErrItemNotFound {
		http.Error(rw, "Item Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Item Not Found", http.StatusInternalServerError)
		return
	}
	items.l.Printf("Prod: %#v", it)
}

// KeyItem is item request from body, to pass back to request context
// after validating request body through middleware
type KeyItem struct{}

//MiddlewareValidateItem validates JSON from request body before passing back to router
func (items Items) MiddlewareValidateItem(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		it := data.Item{}
		err := it.FromJSON(r.Body)
		if err != nil {
			items.l.Println("[ERROR] Deserializing JSON")
			http.Error(rw, "Unable to unmarshall JSON", http.StatusBadRequest)
			return
		}
		err = it.Validate()
		if err != nil {
			items.l.Println("[ERROR] Validation Failed")
			http.Error(
				rw,
				fmt.Sprintf("Error validation input data, %s", err),
				http.StatusBadRequest,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyItem{}, it)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
