// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NormalizedName A single mapping from an unnormalized name to its normalized form.
//
// swagger:model normalizedName
type NormalizedName struct {

	// The normalized version of 'requested_as'.
	// Required: true
	Normalized *string `json:"normalized"`

	// The original, unnormalized name.
	// Required: true
	RequestedAs *string `json:"requested_as"`
}

// Validate validates this normalized name
func (m *NormalizedName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNormalized(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedAs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NormalizedName) validateNormalized(formats strfmt.Registry) error {

	if err := validate.Required("normalized", "body", m.Normalized); err != nil {
		return err
	}

	return nil
}

func (m *NormalizedName) validateRequestedAs(formats strfmt.Registry) error {

	if err := validate.Required("requested_as", "body", m.RequestedAs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this normalized name based on context it is used
func (m *NormalizedName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NormalizedName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NormalizedName) UnmarshalBinary(b []byte) error {
	var res NormalizedName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}