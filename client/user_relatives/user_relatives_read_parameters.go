// Code generated by go-swagger; DO NOT EDIT.

package user_relatives

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

// NewUserRelativesReadParams creates a new UserRelativesReadParams object
// with the default values initialized.
func NewUserRelativesReadParams() *UserRelativesReadParams {
	var ()
	return &UserRelativesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserRelativesReadParamsWithTimeout creates a new UserRelativesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserRelativesReadParamsWithTimeout(timeout time.Duration) *UserRelativesReadParams {
	var ()
	return &UserRelativesReadParams{

		timeout: timeout,
	}
}

// NewUserRelativesReadParamsWithContext creates a new UserRelativesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserRelativesReadParamsWithContext(ctx context.Context) *UserRelativesReadParams {
	var ()
	return &UserRelativesReadParams{

		Context: ctx,
	}
}

// NewUserRelativesReadParamsWithHTTPClient creates a new UserRelativesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserRelativesReadParamsWithHTTPClient(client *http.Client) *UserRelativesReadParams {
	var ()
	return &UserRelativesReadParams{
		HTTPClient: client,
	}
}

/*UserRelativesReadParams contains all the parameters to send to the API endpoint
for the user relatives read operation typically these are written to a http.Request
*/
type UserRelativesReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user relatives read params
func (o *UserRelativesReadParams) WithTimeout(timeout time.Duration) *UserRelativesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user relatives read params
func (o *UserRelativesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user relatives read params
func (o *UserRelativesReadParams) WithContext(ctx context.Context) *UserRelativesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user relatives read params
func (o *UserRelativesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user relatives read params
func (o *UserRelativesReadParams) WithHTTPClient(client *http.Client) *UserRelativesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user relatives read params
func (o *UserRelativesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user relatives read params
func (o *UserRelativesReadParams) WithID(id string) *UserRelativesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user relatives read params
func (o *UserRelativesReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserRelativesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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