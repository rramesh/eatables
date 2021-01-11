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

// Generic message returned as a string
// swagger:response messageResponse
type messageResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericMessage
}

// Successfull Create or Update response
// swagger:response createUpdateResponse
type createUpdateResponseWrapper struct {
	// Response message and data
	// in: body
	Body CreateUpdateMessage
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericMessage
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of eatable items returned in the response
// swagger:response itemResponse
type itemResponseWrapper struct {
	// All items in the DB
	// in: body
	Body struct {
		Items []data.Item `json:"items"`
	}
}

// swagger:parameters
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

// swagger:parameters listSingleItem
type itemIDParamsWrapper struct {
	// The id of the item for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters listItemBySKU deleteItem
type itemSKUParamsWrapper struct {
	// The UUID of the item for which the operation relates
	// in: path
	// required: true
	SKU string `json:"sku"`
}

// swagger:parameters listItemsByVendor
type itemVendorCodeParamsWrapper struct {
	// The UUID of the item for which the operation relates
	// in: path
	// required: true
	VendorCode string `json:"vendorCode"`
}
