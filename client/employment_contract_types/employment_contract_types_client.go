// Code generated by go-swagger; DO NOT EDIT.

package employment_contract_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new employment contract types API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for employment contract types API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EmploymentContractTypesList APIs endpoint that allows employment contract types to be viewed or edited

API endpoint that allows employment contract types to be viewed or edited.
*/
func (a *Client) EmploymentContractTypesList(params *EmploymentContractTypesListParams, authInfo runtime.ClientAuthInfoWriter) (*EmploymentContractTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmploymentContractTypesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "employment_contract_types_list",
		Method:             "GET",
		PathPattern:        "/api/v1/employment_contract_types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmploymentContractTypesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmploymentContractTypesListOK), nil

}

/*
EmploymentContractTypesRead APIs endpoint that allows employment contract types to be viewed or edited

API endpoint that allows employment contract types to be viewed or edited.
*/
func (a *Client) EmploymentContractTypesRead(params *EmploymentContractTypesReadParams, authInfo runtime.ClientAuthInfoWriter) (*EmploymentContractTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmploymentContractTypesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "employment_contract_types_read",
		Method:             "GET",
		PathPattern:        "/api/v1/employment_contract_types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmploymentContractTypesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmploymentContractTypesReadOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}