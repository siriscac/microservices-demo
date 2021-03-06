// Code generated by go-swagger; DO NOT EDIT.

package checkout_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new checkout service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for checkout service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PlaceOrder place order API
*/
func (a *Client) PlaceOrder(params *PlaceOrderParams) (*PlaceOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlaceOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PlaceOrder",
		Method:             "POST",
		PathPattern:        "/orders/checkout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PlaceOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PlaceOrderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
