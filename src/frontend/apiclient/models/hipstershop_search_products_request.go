// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HipstershopSearchProductsRequest hipstershop search products request
// swagger:model hipstershopSearchProductsRequest
type HipstershopSearchProductsRequest struct {

	// query
	Query string `json:"query,omitempty"`
}

// Validate validates this hipstershop search products request
func (m *HipstershopSearchProductsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HipstershopSearchProductsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HipstershopSearchProductsRequest) UnmarshalBinary(b []byte) error {
	var res HipstershopSearchProductsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
