// Code generated by go-swagger; DO NOT EDIT.

package my_leaves

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMyLeavesPartialUpdateParams creates a new MyLeavesPartialUpdateParams object
// with the default values initialized.
func NewMyLeavesPartialUpdateParams() *MyLeavesPartialUpdateParams {
	var ()
	return &MyLeavesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyLeavesPartialUpdateParamsWithTimeout creates a new MyLeavesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyLeavesPartialUpdateParamsWithTimeout(timeout time.Duration) *MyLeavesPartialUpdateParams {
	var ()
	return &MyLeavesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewMyLeavesPartialUpdateParamsWithContext creates a new MyLeavesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyLeavesPartialUpdateParamsWithContext(ctx context.Context) *MyLeavesPartialUpdateParams {
	var ()
	return &MyLeavesPartialUpdateParams{

		Context: ctx,
	}
}

// NewMyLeavesPartialUpdateParamsWithHTTPClient creates a new MyLeavesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyLeavesPartialUpdateParamsWithHTTPClient(client *http.Client) *MyLeavesPartialUpdateParams {
	var ()
	return &MyLeavesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*MyLeavesPartialUpdateParams contains all the parameters to send to the API endpoint
for the my leaves partial update operation typically these are written to a http.Request
*/
type MyLeavesPartialUpdateParams struct {

	/*Data*/
	Data MyLeavesPartialUpdateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) WithTimeout(timeout time.Duration) *MyLeavesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) WithContext(ctx context.Context) *MyLeavesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) WithHTTPClient(client *http.Client) *MyLeavesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) WithData(data MyLeavesPartialUpdateBody) *MyLeavesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) SetData(data MyLeavesPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) WithID(id string) *MyLeavesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the my leaves partial update params
func (o *MyLeavesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MyLeavesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
