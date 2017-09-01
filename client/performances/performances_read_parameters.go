// Code generated by go-swagger; DO NOT EDIT.

package performances

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

// NewPerformancesReadParams creates a new PerformancesReadParams object
// with the default values initialized.
func NewPerformancesReadParams() *PerformancesReadParams {
	var ()
	return &PerformancesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancesReadParamsWithTimeout creates a new PerformancesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPerformancesReadParamsWithTimeout(timeout time.Duration) *PerformancesReadParams {
	var ()
	return &PerformancesReadParams{

		timeout: timeout,
	}
}

// NewPerformancesReadParamsWithContext creates a new PerformancesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPerformancesReadParamsWithContext(ctx context.Context) *PerformancesReadParams {
	var ()
	return &PerformancesReadParams{

		Context: ctx,
	}
}

// NewPerformancesReadParamsWithHTTPClient creates a new PerformancesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPerformancesReadParamsWithHTTPClient(client *http.Client) *PerformancesReadParams {
	var ()
	return &PerformancesReadParams{
		HTTPClient: client,
	}
}

/*PerformancesReadParams contains all the parameters to send to the API endpoint
for the performances read operation typically these are written to a http.Request
*/
type PerformancesReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the performances read params
func (o *PerformancesReadParams) WithTimeout(timeout time.Duration) *PerformancesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performances read params
func (o *PerformancesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performances read params
func (o *PerformancesReadParams) WithContext(ctx context.Context) *PerformancesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performances read params
func (o *PerformancesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performances read params
func (o *PerformancesReadParams) WithHTTPClient(client *http.Client) *PerformancesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performances read params
func (o *PerformancesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the performances read params
func (o *PerformancesReadParams) WithID(id string) *PerformancesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the performances read params
func (o *PerformancesReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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