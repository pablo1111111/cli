// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources The minimum resource requirements with which to run the build. Must specify one or both of the fields in this object.
// swagger:model v1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources
type V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources struct {

	// The minimum number of CPU cores to dedicate to the execution of the build. Can be a non-integer number to allow sharing of a core with other builds, but values are rounded to the nearest 0.1 CPU.
	// Maximum: 16
	// Minimum: 0.1
	Cpus float64 `json:"cpus,omitempty"`

	// The minimum amount of memory in megabytes to make available to the build.
	// Maximum: 131072
	// Minimum: 1
	MemoryMb int64 `json:"memory_mb,omitempty"`
}

// Validate validates this v1 build request recipe resolved ingredients items ingredient options items resources
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryMb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources) validateCpus(formats strfmt.Registry) error {

	if swag.IsZero(m.Cpus) { // not required
		return nil
	}

	if err := validate.Minimum("cpus", "body", float64(m.Cpus), 0.1, false); err != nil {
		return err
	}

	if err := validate.Maximum("cpus", "body", float64(m.Cpus), 16, false); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources) validateMemoryMb(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryMb) { // not required
		return nil
	}

	if err := validate.MinimumInt("memory_mb", "body", int64(m.MemoryMb), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("memory_mb", "body", int64(m.MemoryMb), 131072, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipeResolvedIngredientsItemsIngredientOptionsItemsResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
