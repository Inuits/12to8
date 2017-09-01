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

// NewMyLeavesDeleteParams creates a new MyLeavesDeleteParams object
// with the default values initialized.
func NewMyLeavesDeleteParams() *MyLeavesDeleteParams {
	var ()
	return &MyLeavesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyLeavesDeleteParamsWithTimeout creates a new MyLeavesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyLeavesDeleteParamsWithTimeout(timeout time.Duration) *MyLeavesDeleteParams {
	var ()
	return &MyLeavesDeleteParams{

		timeout: timeout,
	}
}

// NewMyLeavesDeleteParamsWithContext creates a new MyLeavesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyLeavesDeleteParamsWithContext(ctx context.Context) *MyLeavesDeleteParams {
	var ()
	return &MyLeavesDeleteParams{

		Context: ctx,
	}
}

// NewMyLeavesDeleteParamsWithHTTPClient creates a new MyLeavesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyLeavesDeleteParamsWithHTTPClient(client *http.Client) *MyLeavesDeleteParams {
	var ()
	return &MyLeavesDeleteParams{
		HTTPClient: client,
	}
}

/*MyLeavesDeleteParams contains all the parameters to send to the API endpoint
for the my leaves delete operation typically these are written to a http.Request
*/
type MyLeavesDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my leaves delete params
func (o *MyLeavesDeleteParams) WithTimeout(timeout time.Duration) *MyLeavesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my leaves delete params
func (o *MyLeavesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my leaves delete params
func (o *MyLeavesDeleteParams) WithContext(ctx context.Context) *MyLeavesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my leaves delete params
func (o *MyLeavesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my leaves delete params
func (o *MyLeavesDeleteParams) WithHTTPClient(client *http.Client) *MyLeavesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my leaves delete params
func (o *MyLeavesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the my leaves delete params
func (o *MyLeavesDeleteParams) WithID(id string) *MyLeavesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the my leaves delete params
func (o *MyLeavesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MyLeavesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
