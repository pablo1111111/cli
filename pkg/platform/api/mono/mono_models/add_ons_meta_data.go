// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddOnsMetaData add ons meta data
//
// swagger:model AddOnsMetaData
type AddOnsMetaData struct {

	// add ons
	AddOns []*AddOn `json:"addOns"`

	// tier
	Tier string `json:"tier,omitempty"`
}

// Validate validates this add ons meta data
func (m *AddOnsMetaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddOnsMetaData) validateAddOns(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOns) { // not required
		return nil
	}

	for i := 0; i < len(m.AddOns); i++ {
		if swag.IsZero(m.AddOns[i]) { // not required
			continue
		}

		if m.AddOns[i] != nil {
			if err := m.AddOns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addOns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add ons meta data based on the context it is used
func (m *AddOnsMetaData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddOnsMetaData) contextValidateAddOns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AddOns); i++ {

		if m.AddOns[i] != nil {
			if err := m.AddOns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addOns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddOnsMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOnsMetaData) UnmarshalBinary(b []byte) error {
	var res AddOnsMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}