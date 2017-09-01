// Code generated by go-swagger; DO NOT EDIT.

package my_performances

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

// NewMyPerformancesActivityCreateParams creates a new MyPerformancesActivityCreateParams object
// with the default values initialized.
func NewMyPerformancesActivityCreateParams() *MyPerformancesActivityCreateParams {
	var ()
	return &MyPerformancesActivityCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyPerformancesActivityCreateParamsWithTimeout creates a new MyPerformancesActivityCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyPerformancesActivityCreateParamsWithTimeout(timeout time.Duration) *MyPerformancesActivityCreateParams {
	var ()
	return &MyPerformancesActivityCreateParams{

		timeout: timeout,
	}
}

// NewMyPerformancesActivityCreateParamsWithContext creates a new MyPerformancesActivityCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyPerformancesActivityCreateParamsWithContext(ctx context.Context) *MyPerformancesActivityCreateParams {
	var ()
	return &MyPerformancesActivityCreateParams{

		Context: ctx,
	}
}

// NewMyPerformancesActivityCreateParamsWithHTTPClient creates a new MyPerformancesActivityCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyPerformancesActivityCreateParamsWithHTTPClient(client *http.Client) *MyPerformancesActivityCreateParams {
	var ()
	return &MyPerformancesActivityCreateParams{
		HTTPClient: client,
	}
}

/*MyPerformancesActivityCreateParams contains all the parameters to send to the API endpoint
for the my performances activity create operation typically these are written to a http.Request
*/
type MyPerformancesActivityCreateParams struct {

	/*Data*/
	Data MyPerformancesActivityCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) WithTimeout(timeout time.Duration) *MyPerformancesActivityCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) WithContext(ctx context.Context) *MyPerformancesActivityCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) WithHTTPClient(client *http.Client) *MyPerformancesActivityCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) WithData(data MyPerformancesActivityCreateBody) *MyPerformancesActivityCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the my performances activity create params
func (o *MyPerformancesActivityCreateParams) SetData(data MyPerformancesActivityCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *MyPerformancesActivityCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
