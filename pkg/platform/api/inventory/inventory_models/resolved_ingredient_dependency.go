// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResolvedIngredientDependency Resolved Ingredient Dependency
//
// One dependency for a resolved ingredient in a recipe.
//
// swagger:model resolvedIngredientDependency
type ResolvedIngredientDependency struct {

	// The type or types of dependencies on this ingredient version.
	// Required: true
	// Min Items: 1
	// Unique: true
	DependencyTypes []DependencyType `json:"dependency_types"`

	// ingredient version id
	// Required: true
	// Format: uuid
	IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`
}

// Validate validates this resolved ingredient dependency
func (m *ResolvedIngredientDependency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencyTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResolvedIngredientDependency) validateDependencyTypes(formats strfmt.Registry) error {

	if err := validate.Required("dependency_types", "body", m.DependencyTypes); err != nil {
		return err
	}

	iDependencyTypesSize := int64(len(m.DependencyTypes))

	if err := validate.MinItems("dependency_types", "body", iDependencyTypesSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("dependency_types", "body", m.DependencyTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.DependencyTypes); i++ {

		if err := m.DependencyTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dependency_types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ResolvedIngredientDependency) validateIngredientVersionID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version_id", "body", m.IngredientVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resolved ingredient dependency based on the context it is used
func (m *ResolvedIngredientDependency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDependencyTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResolvedIngredientDependency) contextValidateDependencyTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DependencyTypes); i++ {

		if err := m.DependencyTypes[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dependency_types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResolvedIngredientDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolvedIngredientDependency) UnmarshalBinary(b []byte) error {
	var res ResolvedIngredientDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}