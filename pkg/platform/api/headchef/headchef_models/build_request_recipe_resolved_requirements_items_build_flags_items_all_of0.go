// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0 Build Flag Sub Schema
//
// A build flag is a flag that can be passed when building an ingredient version.
// swagger:model buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0
type BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0 struct {

	// Some flags may add dependencies to a build depending on their value. This generator defines how that is done.
	DependencyGenerator []*BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0DependencyGeneratorItems `json:"dependency_generator"`

	// A description of what this build flag does
	// Required: true
	Description *string `json:"description"`

	// Allowed values for enum-type flags.
	EnumValues []string `json:"enum_values"`

	// The actual flag.
	// Required: true
	Flag *string `json:"flag"`

	// The mechanism for setting this flag.
	// Required: true
	// Enum: [configure environment make]
	FlagType *string `json:"flag_type"`

	// inheritance strategy
	InheritanceStrategy *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0InheritanceStrategy `json:"inheritance_strategy,omitempty"`

	// Can this flag be specified more than once? This should never be true for boolean flags.
	// Required: true
	IsRepeatable *bool `json:"is_repeatable"`

	// The type of the build flag's value.
	// Required: true
	// Enum: [boolean enum integer string]
	ValueType *string `json:"value_type"`
}

// Validate validates this build request recipe resolved requirements items build flags items all of0
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencyGenerator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritanceStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRepeatable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateDependencyGenerator(formats strfmt.Registry) error {

	if swag.IsZero(m.DependencyGenerator) { // not required
		return nil
	}

	for i := 0; i < len(m.DependencyGenerator); i++ {
		if swag.IsZero(m.DependencyGenerator[i]) { // not required
			continue
		}

		if m.DependencyGenerator[i] != nil {
			if err := m.DependencyGenerator[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependency_generator" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

var buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeFlagTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["configure","environment","make"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeFlagTypePropEnum = append(buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeFlagTypePropEnum, v)
	}
}

const (

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0FlagTypeConfigure captures enum value "configure"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0FlagTypeConfigure string = "configure"

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0FlagTypeEnvironment captures enum value "environment"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0FlagTypeEnvironment string = "environment"

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0FlagTypeMake captures enum value "make"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0FlagTypeMake string = "make"
)

// prop value enum
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateFlagTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeFlagTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateFlagType(formats strfmt.Registry) error {

	if err := validate.Required("flag_type", "body", m.FlagType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFlagTypeEnum("flag_type", "body", *m.FlagType); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateInheritanceStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.InheritanceStrategy) { // not required
		return nil
	}

	if m.InheritanceStrategy != nil {
		if err := m.InheritanceStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_strategy")
			}
			return err
		}
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateIsRepeatable(formats strfmt.Registry) error {

	if err := validate.Required("is_repeatable", "body", m.IsRepeatable); err != nil {
		return err
	}

	return nil
}

var buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["boolean","enum","integer","string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeValueTypePropEnum = append(buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeValueTypePropEnum, v)
	}
}

const (

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeBoolean captures enum value "boolean"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeBoolean string = "boolean"

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeEnum captures enum value "enum"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeEnum string = "enum"

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeInteger captures enum value "integer"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeInteger string = "integer"

	// BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeString captures enum value "string"
	BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0ValueTypeString string = "string"
)

// prop value enum
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0TypeValueTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) validateValueType(formats strfmt.Registry) error {

	if err := validate.Required("value_type", "body", m.ValueType); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueTypeEnum("value_type", "body", *m.ValueType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0) UnmarshalBinary(b []byte) error {
	var res BuildRequestRecipeResolvedRequirementsItemsBuildFlagsItemsAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}