// Code generated by go-swagger; DO NOT EDIT.

package my_leaves

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new my leaves API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for my leaves API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
MyLeavesCreate APIs endpoint that allows leaves for the currently authenticated user to be viewed or edited

API endpoint that allows leaves for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyLeavesCreate(params *MyLeavesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*MyLeavesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyLeavesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_leaves_create",
		Method:             "POST",
		PathPattern:        "/api/v1/my_leaves/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyLeavesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyLeavesCreateCreated), nil

}

/*
MyLeavesDelete APIs endpoint that allows leaves for the currently authenticated user to be viewed or edited

API endpoint that allows leaves for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyLeavesDelete(params *MyLeavesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*MyLeavesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyLeavesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_leaves_delete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/my_leaves/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyLeavesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyLeavesDeleteNoContent), nil

}

/*
MyLeavesList APIs endpoint that allows leaves for the currently authenticated user to be viewed or edited

API endpoint that allows leaves for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyLeavesList(params *MyLeavesListParams, authInfo runtime.ClientAuthInfoWriter) (*MyLeavesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyLeavesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_leaves_list",
		Method:             "GET",
		PathPattern:        "/api/v1/my_leaves/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyLeavesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyLeavesListOK), nil

}

/*
MyLeavesPartialUpdate APIs endpoint that allows leaves for the currently authenticated user to be viewed or edited

API endpoint that allows leaves for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyLeavesPartialUpdate(params *MyLeavesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*MyLeavesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyLeavesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_leaves_partial_update",
		Method:             "PATCH",
		PathPattern:        "/api/v1/my_leaves/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyLeavesPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyLeavesPartialUpdateOK), nil

}

/*
MyLeavesRead APIs endpoint that allows leaves for the currently authenticated user to be viewed or edited

API endpoint that allows leaves for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyLeavesRead(params *MyLeavesReadParams, authInfo runtime.ClientAuthInfoWriter) (*MyLeavesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyLeavesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_leaves_read",
		Method:             "GET",
		PathPattern:        "/api/v1/my_leaves/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyLeavesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyLeavesReadOK), nil

}

/*
MyLeavesUpdate APIs endpoint that allows leaves for the currently authenticated user to be viewed or edited

API endpoint that allows leaves for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyLeavesUpdate(params *MyLeavesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*MyLeavesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyLeavesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_leaves_update",
		Method:             "PUT",
		PathPattern:        "/api/v1/my_leaves/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyLeavesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyLeavesUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}