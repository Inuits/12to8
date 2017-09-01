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

// NewMyLeavesUpdateParams creates a new MyLeavesUpdateParams object
// with the default values initialized.
func NewMyLeavesUpdateParams() *MyLeavesUpdateParams {
	var ()
	return &MyLeavesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyLeavesUpdateParamsWithTimeout creates a new MyLeavesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyLeavesUpdateParamsWithTimeout(timeout time.Duration) *MyLeavesUpdateParams {
	var ()
	return &MyLeavesUpdateParams{

		timeout: timeout,
	}
}

// NewMyLeavesUpdateParamsWithContext creates a new MyLeavesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyLeavesUpdateParamsWithContext(ctx context.Context) *MyLeavesUpdateParams {
	var ()
	return &MyLeavesUpdateParams{

		Context: ctx,
	}
}

// NewMyLeavesUpdateParamsWithHTTPClient creates a new MyLeavesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyLeavesUpdateParamsWithHTTPClient(client *http.Client) *MyLeavesUpdateParams {
	var ()
	return &MyLeavesUpdateParams{
		HTTPClient: client,
	}
}

/*MyLeavesUpdateParams contains all the parameters to send to the API endpoint
for the my leaves update operation typically these are written to a http.Request
*/
type MyLeavesUpdateParams struct {

	/*Data*/
	Data MyLeavesUpdateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my leaves update params
func (o *MyLeavesUpdateParams) WithTimeout(timeout time.Duration) *MyLeavesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my leaves update params
func (o *MyLeavesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my leaves update params
func (o *MyLeavesUpdateParams) WithContext(ctx context.Context) *MyLeavesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my leaves update params
func (o *MyLeavesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my leaves update params
func (o *MyLeavesUpdateParams) WithHTTPClient(client *http.Client) *MyLeavesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my leaves update params
func (o *MyLeavesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the my leaves update params
func (o *MyLeavesUpdateParams) WithData(data MyLeavesUpdateBody) *MyLeavesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the my leaves update params
func (o *MyLeavesUpdateParams) SetData(data MyLeavesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the my leaves update params
func (o *MyLeavesUpdateParams) WithID(id string) *MyLeavesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the my leaves update params
func (o *MyLeavesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MyLeavesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
