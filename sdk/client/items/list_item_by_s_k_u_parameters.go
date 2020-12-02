// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListItemBySKUParams creates a new ListItemBySKUParams object
// with the default values initialized.
func NewListItemBySKUParams() *ListItemBySKUParams {
	var ()
	return &ListItemBySKUParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListItemBySKUParamsWithTimeout creates a new ListItemBySKUParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListItemBySKUParamsWithTimeout(timeout time.Duration) *ListItemBySKUParams {
	var ()
	return &ListItemBySKUParams{

		timeout: timeout,
	}
}

// NewListItemBySKUParamsWithContext creates a new ListItemBySKUParams object
// with the default values initialized, and the ability to set a context for a request
func NewListItemBySKUParamsWithContext(ctx context.Context) *ListItemBySKUParams {
	var ()
	return &ListItemBySKUParams{

		Context: ctx,
	}
}

// NewListItemBySKUParamsWithHTTPClient creates a new ListItemBySKUParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListItemBySKUParamsWithHTTPClient(client *http.Client) *ListItemBySKUParams {
	var ()
	return &ListItemBySKUParams{
		HTTPClient: client,
	}
}

/*ListItemBySKUParams contains all the parameters to send to the API endpoint
for the list item by s k u operation typically these are written to a http.Request
*/
type ListItemBySKUParams struct {

	/*UUID
	  The UUID of the item for which the operation relates

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list item by s k u params
func (o *ListItemBySKUParams) WithTimeout(timeout time.Duration) *ListItemBySKUParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list item by s k u params
func (o *ListItemBySKUParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list item by s k u params
func (o *ListItemBySKUParams) WithContext(ctx context.Context) *ListItemBySKUParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list item by s k u params
func (o *ListItemBySKUParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list item by s k u params
func (o *ListItemBySKUParams) WithHTTPClient(client *http.Client) *ListItemBySKUParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list item by s k u params
func (o *ListItemBySKUParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the list item by s k u params
func (o *ListItemBySKUParams) WithUUID(uuid string) *ListItemBySKUParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the list item by s k u params
func (o *ListItemBySKUParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ListItemBySKUParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
