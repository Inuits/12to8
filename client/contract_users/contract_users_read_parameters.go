// Code generated by go-swagger; DO NOT EDIT.

package contract_users

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

// NewContractUsersReadParams creates a new ContractUsersReadParams object
// with the default values initialized.
func NewContractUsersReadParams() *ContractUsersReadParams {
	var ()
	return &ContractUsersReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContractUsersReadParamsWithTimeout creates a new ContractUsersReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContractUsersReadParamsWithTimeout(timeout time.Duration) *ContractUsersReadParams {
	var ()
	return &ContractUsersReadParams{

		timeout: timeout,
	}
}

// NewContractUsersReadParamsWithContext creates a new ContractUsersReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewContractUsersReadParamsWithContext(ctx context.Context) *ContractUsersReadParams {
	var ()
	return &ContractUsersReadParams{

		Context: ctx,
	}
}

// NewContractUsersReadParamsWithHTTPClient creates a new ContractUsersReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContractUsersReadParamsWithHTTPClient(client *http.Client) *ContractUsersReadParams {
	var ()
	return &ContractUsersReadParams{
		HTTPClient: client,
	}
}

/*ContractUsersReadParams contains all the parameters to send to the API endpoint
for the contract users read operation typically these are written to a http.Request
*/
type ContractUsersReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contract users read params
func (o *ContractUsersReadParams) WithTimeout(timeout time.Duration) *ContractUsersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contract users read params
func (o *ContractUsersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contract users read params
func (o *ContractUsersReadParams) WithContext(ctx context.Context) *ContractUsersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contract users read params
func (o *ContractUsersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contract users read params
func (o *ContractUsersReadParams) WithHTTPClient(client *http.Client) *ContractUsersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contract users read params
func (o *ContractUsersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contract users read params
func (o *ContractUsersReadParams) WithID(id string) *ContractUsersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contract users read params
func (o *ContractUsersReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContractUsersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
