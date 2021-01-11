package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"

	"github.com/rramesh/eatables/data"
)

// ItemHandler is a http.Handler
type ItemHandler struct {
	l      hclog.Logger
	v      *data.Validation
	itemDB *data.ItemDB
}

// KeyItem is item request from body, to pass back to request context
// after validating request body through middleware
type KeyItem struct{}

// NewItems creates a items handler with the given logger
func NewItems(l hclog.Logger, v *data.Validation, idb *data.ItemDB) *ItemHandler {
	return &ItemHandler{l, v, idb}
}

// GenericMessage holds a message string to be sent as JSON
type GenericMessage struct {
	Message string `json:"message"`
}

// CreateUpdateMessage holds message response and data to be sent as JSON
type CreateUpdateMessage struct {
	// Status Message
	// example: Item Added/Updated Successfully
	Message string `json:"message"`

	// SKU UUID of item added/updated
	// example: dca98ae0-2b9f-441c-9007-ef824f1581fd
	SKU string `json:"sku"`
}

// ValidationError holds a slice of error messages to be sent as JSON
type ValidationError struct {
	Message []string `json:"message"`
}

// getItemID gets the ID from the URL
// Panics if it cannot convert String to integer
// It should not happen as the router already
// validates, but caught as a good practice
func getItemID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Should have never happened
		panic(err)
	}
	return id
}

// getSKUFromRequest gets the SKU UUID from the URL
func getSKUFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	sku := vars["sku"]
	if sku == "" {
		// Should never have happened
		panic("SKU UUID is missing/empty in the Path. Should have been handled by the router")
	}
	return sku
}

// getVCFromRequest gets the Vendor Code UUID from the URL
func getVCFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	vc := vars["vendorCode"]
	if vc == "" {
		// Should never have happened
		panic("SKU UUID is missing/empty in the Path. Should have been handled by the router")
	}
	return vc
}
