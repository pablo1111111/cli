// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1SearchIngredientsResponseIngredientsItemsLatestVersion Ingredient Version
//
// The full ingredient version data model. Returned from all read requests.
// swagger:model v1SearchIngredientsResponseIngredientsItemsLatestVersion
type V1SearchIngredientsResponseIngredientsItemsLatestVersion struct {
	V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf0

	V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf1

	V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf2

	V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf3
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1SearchIngredientsResponseIngredientsItemsLatestVersion) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf0 = aO0

	// AO1
	var aO1 V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf1 = aO1

	// AO2
	var aO2 V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf2 = aO2

	// AO3
	var aO3 V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf3
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf3 = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1SearchIngredientsResponseIngredientsItemsLatestVersion) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 search ingredients response ingredients items latest version
func (m *V1SearchIngredientsResponseIngredientsItemsLatestVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf0
	if err := m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf1
	if err := m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf2
	if err := m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf3
	if err := m.V1SearchIngredientsResponseIngredientsItemsLatestVersionAllOf3.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchIngredientsResponseIngredientsItemsLatestVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchIngredientsResponseIngredientsItemsLatestVersion) UnmarshalBinary(b []byte) error {
	var res V1SearchIngredientsResponseIngredientsItemsLatestVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}