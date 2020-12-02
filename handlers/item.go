package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rramesh/eatables/data"
)

// Items is a http.Handler
type Items struct {
	l *log.Logger
	v *data.Validation
}

// KeyItem is item request from body, to pass back to request context
// after validating request body through middleware
type KeyItem struct{}

// NewItems creates a items handler with the given logger
func NewItems(l *log.Logger, v *data.Validation) *Items {
	return &Items{l, v}
}

// GenericError holds a message string to be sent as JSON
type GenericError struct {
	Message string
}

// ValidationError holds a slice of error messages to be sent as JSON
type ValidationError struct {
	Message []string
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

// getUUID gets the UUID from the URL
func getUUID(r *http.Request) string {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	if uuid == "" {
		// Should never have happened
		panic("UUID is missing/empty in the Path. Should have been handled by the router")
	}
	return uuid
}
