// Code generated by go-swagger; DO NOT EDIT.

package cart_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cart service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cart service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddItem add item API
*/
func (a *Client) AddItem(params *AddItemParams) (*AddItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddItem",
		Method:             "POST",
		PathPattern:        "/carts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddItemReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddItemOK), nil

}

/*
EmptyCart empty cart API
*/
func (a *Client) EmptyCart(params *EmptyCartParams) (*EmptyCartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmptyCartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmptyCart",
		Method:             "DELETE",
		PathPattern:        "/carts/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EmptyCartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmptyCartOK), nil

}

/*
GetCart get cart API
*/
func (a *Client) GetCart(params *GetCartParams) (*GetCartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCart",
		Method:             "GET",
		PathPattern:        "/carts/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCartOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}