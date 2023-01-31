// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Configuration configuration
//
// swagger:model configuration
type Configuration struct {

	// tou link
	TouLink string `json:"touLink,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this configuration
func (m *Configuration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this configuration based on context it is used
func (m *Configuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configuration) UnmarshalBinary(b []byte) error {
	var res Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
