// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewServicesMonthInfoListParams creates a new ServicesMonthInfoListParams object
// with the default values initialized.
func NewServicesMonthInfoListParams() *ServicesMonthInfoListParams {

	return &ServicesMonthInfoListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServicesMonthInfoListParamsWithTimeout creates a new ServicesMonthInfoListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServicesMonthInfoListParamsWithTimeout(timeout time.Duration) *ServicesMonthInfoListParams {

	return &ServicesMonthInfoListParams{

		timeout: timeout,
	}
}

// NewServicesMonthInfoListParamsWithContext creates a new ServicesMonthInfoListParams object
// with the default values initialized, and the ability to set a context for a request
func NewServicesMonthInfoListParamsWithContext(ctx context.Context) *ServicesMonthInfoListParams {

	return &ServicesMonthInfoListParams{

		Context: ctx,
	}
}

// NewServicesMonthInfoListParamsWithHTTPClient creates a new ServicesMonthInfoListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServicesMonthInfoListParamsWithHTTPClient(client *http.Client) *ServicesMonthInfoListParams {

	return &ServicesMonthInfoListParams{
		HTTPClient: client,
	}
}

/*ServicesMonthInfoListParams contains all the parameters to send to the API endpoint
for the services month info list operation typically these are written to a http.Request
*/
type ServicesMonthInfoListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the services month info list params
func (o *ServicesMonthInfoListParams) WithTimeout(timeout time.Duration) *ServicesMonthInfoListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the services month info list params
func (o *ServicesMonthInfoListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the services month info list params
func (o *ServicesMonthInfoListParams) WithContext(ctx context.Context) *ServicesMonthInfoListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the services month info list params
func (o *ServicesMonthInfoListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the services month info list params
func (o *ServicesMonthInfoListParams) WithHTTPClient(client *http.Client) *ServicesMonthInfoListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the services month info list params
func (o *ServicesMonthInfoListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServicesMonthInfoListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}