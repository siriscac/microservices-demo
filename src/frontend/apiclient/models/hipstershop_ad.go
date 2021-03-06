// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HipstershopAd hipstershop ad
// swagger:model hipstershopAd
type HipstershopAd struct {

	// url to redirect to when an ad is clicked.
	RedirectURL string `json:"redirect_url,omitempty"`

	// short advertisement text to display.
	Text string `json:"text,omitempty"`
}

// Validate validates this hipstershop ad
func (m *HipstershopAd) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HipstershopAd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HipstershopAd) UnmarshalBinary(b []byte) error {
	var res HipstershopAd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
