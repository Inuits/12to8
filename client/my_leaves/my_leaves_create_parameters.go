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

// NewMyLeavesCreateParams creates a new MyLeavesCreateParams object
// with the default values initialized.
func NewMyLeavesCreateParams() *MyLeavesCreateParams {
	var ()
	return &MyLeavesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyLeavesCreateParamsWithTimeout creates a new MyLeavesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyLeavesCreateParamsWithTimeout(timeout time.Duration) *MyLeavesCreateParams {
	var ()
	return &MyLeavesCreateParams{

		timeout: timeout,
	}
}

// NewMyLeavesCreateParamsWithContext creates a new MyLeavesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyLeavesCreateParamsWithContext(ctx context.Context) *MyLeavesCreateParams {
	var ()
	return &MyLeavesCreateParams{

		Context: ctx,
	}
}

// NewMyLeavesCreateParamsWithHTTPClient creates a new MyLeavesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyLeavesCreateParamsWithHTTPClient(client *http.Client) *MyLeavesCreateParams {
	var ()
	return &MyLeavesCreateParams{
		HTTPClient: client,
	}
}

/*MyLeavesCreateParams contains all the parameters to send to the API endpoint
for the my leaves create operation typically these are written to a http.Request
*/
type MyLeavesCreateParams struct {

	/*Data*/
	Data MyLeavesCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my leaves create params
func (o *MyLeavesCreateParams) WithTimeout(timeout time.Duration) *MyLeavesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my leaves create params
func (o *MyLeavesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my leaves create params
func (o *MyLeavesCreateParams) WithContext(ctx context.Context) *MyLeavesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my leaves create params
func (o *MyLeavesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my leaves create params
func (o *MyLeavesCreateParams) WithHTTPClient(client *http.Client) *MyLeavesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my leaves create params
func (o *MyLeavesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the my leaves create params
func (o *MyLeavesCreateParams) WithData(data MyLeavesCreateBody) *MyLeavesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the my leaves create params
func (o *MyLeavesCreateParams) SetData(data MyLeavesCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *MyLeavesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}