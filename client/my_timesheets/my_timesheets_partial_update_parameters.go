// Code generated by go-swagger; DO NOT EDIT.

package my_timesheets

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

// NewMyTimesheetsPartialUpdateParams creates a new MyTimesheetsPartialUpdateParams object
// with the default values initialized.
func NewMyTimesheetsPartialUpdateParams() *MyTimesheetsPartialUpdateParams {
	var ()
	return &MyTimesheetsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMyTimesheetsPartialUpdateParamsWithTimeout creates a new MyTimesheetsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMyTimesheetsPartialUpdateParamsWithTimeout(timeout time.Duration) *MyTimesheetsPartialUpdateParams {
	var ()
	return &MyTimesheetsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewMyTimesheetsPartialUpdateParamsWithContext creates a new MyTimesheetsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMyTimesheetsPartialUpdateParamsWithContext(ctx context.Context) *MyTimesheetsPartialUpdateParams {
	var ()
	return &MyTimesheetsPartialUpdateParams{

		Context: ctx,
	}
}

// NewMyTimesheetsPartialUpdateParamsWithHTTPClient creates a new MyTimesheetsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMyTimesheetsPartialUpdateParamsWithHTTPClient(client *http.Client) *MyTimesheetsPartialUpdateParams {
	var ()
	return &MyTimesheetsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*MyTimesheetsPartialUpdateParams contains all the parameters to send to the API endpoint
for the my timesheets partial update operation typically these are written to a http.Request
*/
type MyTimesheetsPartialUpdateParams struct {

	/*Data*/
	Data MyTimesheetsPartialUpdateBody
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) WithTimeout(timeout time.Duration) *MyTimesheetsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) WithContext(ctx context.Context) *MyTimesheetsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) WithHTTPClient(client *http.Client) *MyTimesheetsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) WithData(data MyTimesheetsPartialUpdateBody) *MyTimesheetsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) SetData(data MyTimesheetsPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) WithID(id string) *MyTimesheetsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the my timesheets partial update params
func (o *MyTimesheetsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MyTimesheetsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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