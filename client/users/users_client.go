// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UsersList APIs endpoint that allows users to be viewed or edited

API endpoint that allows users to be viewed or edited.
*/
func (a *Client) UsersList(params *UsersListParams, authInfo runtime.ClientAuthInfoWriter) (*UsersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "users_list",
		Method:             "GET",
		PathPattern:        "/api/v1/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersListOK), nil

}

/*
UsersRead APIs endpoint that allows users to be viewed or edited

API endpoint that allows users to be viewed or edited.
*/
func (a *Client) UsersRead(params *UsersReadParams, authInfo runtime.ClientAuthInfoWriter) (*UsersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "users_read",
		Method:             "GET",
		PathPattern:        "/api/v1/users/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersReadOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}