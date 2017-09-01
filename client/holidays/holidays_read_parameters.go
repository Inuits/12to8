// Code generated by go-swagger; DO NOT EDIT.

package holidays

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

// NewHolidaysReadParams creates a new HolidaysReadParams object
// with the default values initialized.
func NewHolidaysReadParams() *HolidaysReadParams {
	var ()
	return &HolidaysReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHolidaysReadParamsWithTimeout creates a new HolidaysReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHolidaysReadParamsWithTimeout(timeout time.Duration) *HolidaysReadParams {
	var ()
	return &HolidaysReadParams{

		timeout: timeout,
	}
}

// NewHolidaysReadParamsWithContext creates a new HolidaysReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewHolidaysReadParamsWithContext(ctx context.Context) *HolidaysReadParams {
	var ()
	return &HolidaysReadParams{

		Context: ctx,
	}
}

// NewHolidaysReadParamsWithHTTPClient creates a new HolidaysReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHolidaysReadParamsWithHTTPClient(client *http.Client) *HolidaysReadParams {
	var ()
	return &HolidaysReadParams{
		HTTPClient: client,
	}
}

/*HolidaysReadParams contains all the parameters to send to the API endpoint
for the holidays read operation typically these are written to a http.Request
*/
type HolidaysReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the holidays read params
func (o *HolidaysReadParams) WithTimeout(timeout time.Duration) *HolidaysReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the holidays read params
func (o *HolidaysReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the holidays read params
func (o *HolidaysReadParams) WithContext(ctx context.Context) *HolidaysReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the holidays read params
func (o *HolidaysReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the holidays read params
func (o *HolidaysReadParams) WithHTTPClient(client *http.Client) *HolidaysReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the holidays read params
func (o *HolidaysReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the holidays read params
func (o *HolidaysReadParams) WithID(id string) *HolidaysReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the holidays read params
func (o *HolidaysReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HolidaysReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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