// Code generated by go-swagger; DO NOT EDIT.

package rest_model_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Environment environment
//
// swagger:model environment
type Environment struct {

	// active
	Active bool `json:"active,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`

	// ziti identity Id
	ZitiIdentityID string `json:"zitiIdentityId,omitempty"`
}

// Validate validates this environment
func (m *Environment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this environment based on context it is used
func (m *Environment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Environment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Environment) UnmarshalBinary(b []byte) error {
	var res Environment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
