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

// NewMyPerformancesStandbyReadParams creates a new MyPerformancesStandbyReadParams object
// with the default values initialized.
func NewMyPerformancesStandbyReadParams() *MyPerformancesStandbyReadParams {
	var ()
	return &MyPerformancesStandbyReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyPerformancesStandbyReadParamsWithTimeout creates a new MyPerformancesStandbyReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyPerformancesStandbyReadParamsWithTimeout(timeout time.Duration) *MyPerformancesStandbyReadParams {
	var ()
	return &MyPerformancesStandbyReadParams{

		timeout: timeout,
	}
}

// NewMyPerformancesStandbyReadParamsWithContext creates a new MyPerformancesStandbyReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyPerformancesStandbyReadParamsWithContext(ctx context.Context) *MyPerformancesStandbyReadParams {
	var ()
	return &MyPerformancesStandbyReadParams{

		Context: ctx,
	}
}

// NewMyPerformancesStandbyReadParamsWithHTTPClient creates a new MyPerformancesStandbyReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyPerformancesStandbyReadParamsWithHTTPClient(client *http.Client) *MyPerformancesStandbyReadParams {
	var ()
	return &MyPerformancesStandbyReadParams{
		HTTPClient: client,
	}
}

/*MyPerformancesStandbyReadParams contains all the parameters to send to the API endpoint
for the my performances standby read operation typically these are written to a http.Request
*/
type MyPerformancesStandbyReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) WithTimeout(timeout time.Duration) *MyPerformancesStandbyReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) WithContext(ctx context.Context) *MyPerformancesStandbyReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) WithHTTPClient(client *http.Client) *MyPerformancesStandbyReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) WithID(id string) *MyPerformancesStandbyReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the my performances standby read params
func (o *MyPerformancesStandbyReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MyPerformancesStandbyReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
