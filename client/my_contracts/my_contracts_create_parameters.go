// Code generated by go-swagger; DO NOT EDIT.

package my_contracts

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

// NewMyContractsCreateParams creates a new MyContractsCreateParams object
// with the default values initialized.
func NewMyContractsCreateParams() *MyContractsCreateParams {
	var ()
	return &MyContractsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyContractsCreateParamsWithTimeout creates a new MyContractsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyContractsCreateParamsWithTimeout(timeout time.Duration) *MyContractsCreateParams {
	var ()
	return &MyContractsCreateParams{

		timeout: timeout,
	}
}

// NewMyContractsCreateParamsWithContext creates a new MyContractsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyContractsCreateParamsWithContext(ctx context.Context) *MyContractsCreateParams {
	var ()
	return &MyContractsCreateParams{

		Context: ctx,
	}
}

// NewMyContractsCreateParamsWithHTTPClient creates a new MyContractsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyContractsCreateParamsWithHTTPClient(client *http.Client) *MyContractsCreateParams {
	var ()
	return &MyContractsCreateParams{
		HTTPClient: client,
	}
}

/*MyContractsCreateParams contains all the parameters to send to the API endpoint
for the my contracts create operation typically these are written to a http.Request
*/
type MyContractsCreateParams struct {

	/*Data*/
	Data MyContractsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my contracts create params
func (o *MyContractsCreateParams) WithTimeout(timeout time.Duration) *MyContractsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my contracts create params
func (o *MyContractsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my contracts create params
func (o *MyContractsCreateParams) WithContext(ctx context.Context) *MyContractsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my contracts create params
func (o *MyContractsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my contracts create params
func (o *MyContractsCreateParams) WithHTTPClient(client *http.Client) *MyContractsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my contracts create params
func (o *MyContractsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the my contracts create params
func (o *MyContractsCreateParams) WithData(data MyContractsCreateBody) *MyContractsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the my contracts create params
func (o *MyContractsCreateParams) SetData(data MyContractsCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *MyContractsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
