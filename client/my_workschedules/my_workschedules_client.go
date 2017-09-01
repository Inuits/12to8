// Code generated by go-swagger; DO NOT EDIT.

package my_workschedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new my workschedules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for my workschedules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
MyWorkschedulesCreate APIs endpoint that allows workschedules for the currently authenticated user to be viewed or edited

API endpoint that allows workschedules for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyWorkschedulesCreate(params *MyWorkschedulesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*MyWorkschedulesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyWorkschedulesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_workschedules_create",
		Method:             "POST",
		PathPattern:        "/api/v1/my_workschedules/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyWorkschedulesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyWorkschedulesCreateCreated), nil

}

/*
MyWorkschedulesDelete APIs endpoint that allows workschedules for the currently authenticated user to be viewed or edited

API endpoint that allows workschedules for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyWorkschedulesDelete(params *MyWorkschedulesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*MyWorkschedulesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyWorkschedulesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_workschedules_delete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/my_workschedules/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyWorkschedulesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyWorkschedulesDeleteNoContent), nil

}

/*
MyWorkschedulesList APIs endpoint that allows workschedules for the currently authenticated user to be viewed or edited

API endpoint that allows workschedules for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyWorkschedulesList(params *MyWorkschedulesListParams, authInfo runtime.ClientAuthInfoWriter) (*MyWorkschedulesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyWorkschedulesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_workschedules_list",
		Method:             "GET",
		PathPattern:        "/api/v1/my_workschedules/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyWorkschedulesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyWorkschedulesListOK), nil

}

/*
MyWorkschedulesPartialUpdate APIs endpoint that allows workschedules for the currently authenticated user to be viewed or edited

API endpoint that allows workschedules for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyWorkschedulesPartialUpdate(params *MyWorkschedulesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*MyWorkschedulesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyWorkschedulesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_workschedules_partial_update",
		Method:             "PATCH",
		PathPattern:        "/api/v1/my_workschedules/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyWorkschedulesPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyWorkschedulesPartialUpdateOK), nil

}

/*
MyWorkschedulesRead APIs endpoint that allows workschedules for the currently authenticated user to be viewed or edited

API endpoint that allows workschedules for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyWorkschedulesRead(params *MyWorkschedulesReadParams, authInfo runtime.ClientAuthInfoWriter) (*MyWorkschedulesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyWorkschedulesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_workschedules_read",
		Method:             "GET",
		PathPattern:        "/api/v1/my_workschedules/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyWorkschedulesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyWorkschedulesReadOK), nil

}

/*
MyWorkschedulesUpdate APIs endpoint that allows workschedules for the currently authenticated user to be viewed or edited

API endpoint that allows workschedules for the currently authenticated user to be viewed or edited.
*/
func (a *Client) MyWorkschedulesUpdate(params *MyWorkschedulesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*MyWorkschedulesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMyWorkschedulesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "my_workschedules_update",
		Method:             "PUT",
		PathPattern:        "/api/v1/my_workschedules/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MyWorkschedulesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MyWorkschedulesUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}