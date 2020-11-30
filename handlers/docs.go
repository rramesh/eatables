// Package handlers - classification of eatables API
//
// Documentation for Eatables API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// 	- application/json
//
// Produces:
// 	- application/json
// swagger:meta
package handlers

import "github.com/rramesh/eatables/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// swagger:response itemResponse
// A list of eatable items returned in the response
type itemResponseWrapper struct {
	// All Products in the DB
	// in: body
	Body []data.Item
}

// swagger:parameters deleteItem
type itemIDParameterWrapper struct {
	// The ID of the Item to be deleted from the DB
	//in: path
	//required: true
	ID int `json:"id"`
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct{}

// swagger:parameters updateItem createItem
type itemParamsWrapper struct {
	// Item data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Item
}

// swagger:parameters updateItem
type itemIDParamsWrapper struct {
	// The id of the item for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
