// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

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

// IngredientVersion Ingredient Version
//
// An ingredient version is a specific version of an ingredient.
// swagger:model ingredientVersion
type IngredientVersion struct {

	// One or more optional build flags that can be applied to this ingredient when it is built.
	BuildFlags []*IngredientVersionBuildFlagsItems0 `json:"build_flags"`

	// description
	// Required: true
	Description *string `json:"description"`

	// ingredient id
	// Required: true
	// Format: uuid
	IngredientID *strfmt.UUID `json:"ingredient_id"`

	// ingredient version id
	// Required: true
	// Format: uuid
	IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`

	// A release is not stable if it is an alpha, beta, developer test release, etc.
	// Required: true
	IsStableRelease *bool `json:"is_stable_release"`

	// The packages provided by this ingredient version.
	Provides map[string]IngredientVersionProvidesAnon `json:"provides,omitempty"`

	// The release date for this version of the ingredient.
	// Required: true
	// Format: date-time
	ReleaseDate *strfmt.DateTime `json:"release_date"`

	// An internal version that starts at 1 and is incremented if any change is made to how this ingredient version is built.
	// Required: true
	// Minimum: 1
	Revision *int32 `json:"revision"`

	// A link to this ingredient's source code on the public Internet.
	// Required: true
	// Format: uri
	SourceURI *strfmt.URI `json:"source_uri"`

	// version
	// Required: true
	Version *string `json:"version"`

	// Ingredients which this ingredient depends on, referenced by the package names and package versions that they provide.
	Dependencies []*IngredientVersionDependenciesItems0 `json:"dependencies"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IngredientVersion) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		BuildFlags []*IngredientVersionBuildFlagsItems0 `json:"build_flags"`

		Description *string `json:"description"`

		IngredientID *strfmt.UUID `json:"ingredient_id"`

		IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`

		IsStableRelease *bool `json:"is_stable_release"`

		Provides map[string]IngredientVersionProvidesAnon `json:"provides,omitempty"`

		ReleaseDate *strfmt.DateTime `json:"release_date"`

		Revision *int32 `json:"revision"`

		SourceURI *strfmt.URI `json:"source_uri"`

		Version *string `json:"version"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.BuildFlags = dataAO0.BuildFlags

	m.Description = dataAO0.Description

	m.IngredientID = dataAO0.IngredientID

	m.IngredientVersionID = dataAO0.IngredientVersionID

	m.IsStableRelease = dataAO0.IsStableRelease

	m.Provides = dataAO0.Provides

	m.ReleaseDate = dataAO0.ReleaseDate

	m.Revision = dataAO0.Revision

	m.SourceURI = dataAO0.SourceURI

	m.Version = dataAO0.Version

	// AO1
	var dataAO1 struct {
		Dependencies []*IngredientVersionDependenciesItems0 `json:"dependencies"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Dependencies = dataAO1.Dependencies

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IngredientVersion) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		BuildFlags []*IngredientVersionBuildFlagsItems0 `json:"build_flags"`

		Description *string `json:"description"`

		IngredientID *strfmt.UUID `json:"ingredient_id"`

		IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`

		IsStableRelease *bool `json:"is_stable_release"`

		Provides map[string]IngredientVersionProvidesAnon `json:"provides,omitempty"`

		ReleaseDate *strfmt.DateTime `json:"release_date"`

		Revision *int32 `json:"revision"`

		SourceURI *strfmt.URI `json:"source_uri"`

		Version *string `json:"version"`
	}

	dataAO0.BuildFlags = m.BuildFlags

	dataAO0.Description = m.Description

	dataAO0.IngredientID = m.IngredientID

	dataAO0.IngredientVersionID = m.IngredientVersionID

	dataAO0.IsStableRelease = m.IsStableRelease

	dataAO0.Provides = m.Provides

	dataAO0.ReleaseDate = m.ReleaseDate

	dataAO0.Revision = m.Revision

	dataAO0.SourceURI = m.SourceURI

	dataAO0.Version = m.Version

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		Dependencies []*IngredientVersionDependenciesItems0 `json:"dependencies"`
	}

	dataAO1.Dependencies = m.Dependencies

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ingredient version
func (m *IngredientVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsStableRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngredientVersion) validateBuildFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildFlags) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildFlags); i++ {
		if swag.IsZero(m.BuildFlags[i]) { // not required
			continue
		}

		if m.BuildFlags[i] != nil {
			if err := m.BuildFlags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_flags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersion) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateIngredientID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_id", "body", m.IngredientID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_id", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateIngredientVersionID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version_id", "body", m.IngredientVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateIsStableRelease(formats strfmt.Registry) error {

	if err := validate.Required("is_stable_release", "body", m.IsStableRelease); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateProvides(formats strfmt.Registry) error {

	if swag.IsZero(m.Provides) { // not required
		return nil
	}

	for k := range m.Provides {

		if swag.IsZero(m.Provides[k]) { // not required
			continue
		}
		if val, ok := m.Provides[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersion) validateReleaseDate(formats strfmt.Registry) error {

	if err := validate.Required("release_date", "body", m.ReleaseDate); err != nil {
		return err
	}

	if err := validate.FormatOf("release_date", "body", "date-time", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	if err := validate.MinimumInt("revision", "body", int64(*m.Revision), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateSourceURI(formats strfmt.Registry) error {

	if err := validate.Required("source_uri", "body", m.SourceURI); err != nil {
		return err
	}

	if err := validate.FormatOf("source_uri", "body", "uri", m.SourceURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersion) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersion) UnmarshalBinary(b []byte) error {
	var res IngredientVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IngredientVersionBuildFlagsItems0 Build Flag Sub Schema
//
// A build flag is a flag that can be passed when building an ingredient version.
// swagger:model IngredientVersionBuildFlagsItems0
type IngredientVersionBuildFlagsItems0 struct {

	// Some flags may add dependencies to a build depending on their value. This generator defines how that is done.
	DependencyGenerator []*IngredientVersionBuildFlagsItems0DependencyGeneratorItems0 `json:"dependency_generator"`

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
	InheritanceStrategy *IngredientVersionBuildFlagsItems0InheritanceStrategy `json:"inheritance_strategy,omitempty"`

	// Can this flag be specified more than once? This should never be true for boolean flags.
	// Required: true
	IsRepeatable *bool `json:"is_repeatable"`

	// The type of the build flag's value.
	// Required: true
	// Enum: [boolean enum integer string]
	ValueType *string `json:"value_type"`
}

// Validate validates this ingredient version build flags items0
func (m *IngredientVersionBuildFlagsItems0) Validate(formats strfmt.Registry) error {
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

func (m *IngredientVersionBuildFlagsItems0) validateDependencyGenerator(formats strfmt.Registry) error {

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

func (m *IngredientVersionBuildFlagsItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersionBuildFlagsItems0) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

var ingredientVersionBuildFlagsItems0TypeFlagTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["configure","environment","make"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionBuildFlagsItems0TypeFlagTypePropEnum = append(ingredientVersionBuildFlagsItems0TypeFlagTypePropEnum, v)
	}
}

const (

	// IngredientVersionBuildFlagsItems0FlagTypeConfigure captures enum value "configure"
	IngredientVersionBuildFlagsItems0FlagTypeConfigure string = "configure"

	// IngredientVersionBuildFlagsItems0FlagTypeEnvironment captures enum value "environment"
	IngredientVersionBuildFlagsItems0FlagTypeEnvironment string = "environment"

	// IngredientVersionBuildFlagsItems0FlagTypeMake captures enum value "make"
	IngredientVersionBuildFlagsItems0FlagTypeMake string = "make"
)

// prop value enum
func (m *IngredientVersionBuildFlagsItems0) validateFlagTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionBuildFlagsItems0TypeFlagTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionBuildFlagsItems0) validateFlagType(formats strfmt.Registry) error {

	if err := validate.Required("flag_type", "body", m.FlagType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFlagTypeEnum("flag_type", "body", *m.FlagType); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersionBuildFlagsItems0) validateInheritanceStrategy(formats strfmt.Registry) error {

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

func (m *IngredientVersionBuildFlagsItems0) validateIsRepeatable(formats strfmt.Registry) error {

	if err := validate.Required("is_repeatable", "body", m.IsRepeatable); err != nil {
		return err
	}

	return nil
}

var ingredientVersionBuildFlagsItems0TypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["boolean","enum","integer","string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionBuildFlagsItems0TypeValueTypePropEnum = append(ingredientVersionBuildFlagsItems0TypeValueTypePropEnum, v)
	}
}

const (

	// IngredientVersionBuildFlagsItems0ValueTypeBoolean captures enum value "boolean"
	IngredientVersionBuildFlagsItems0ValueTypeBoolean string = "boolean"

	// IngredientVersionBuildFlagsItems0ValueTypeEnum captures enum value "enum"
	IngredientVersionBuildFlagsItems0ValueTypeEnum string = "enum"

	// IngredientVersionBuildFlagsItems0ValueTypeInteger captures enum value "integer"
	IngredientVersionBuildFlagsItems0ValueTypeInteger string = "integer"

	// IngredientVersionBuildFlagsItems0ValueTypeString captures enum value "string"
	IngredientVersionBuildFlagsItems0ValueTypeString string = "string"
)

// prop value enum
func (m *IngredientVersionBuildFlagsItems0) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionBuildFlagsItems0TypeValueTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionBuildFlagsItems0) validateValueType(formats strfmt.Registry) error {

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
func (m *IngredientVersionBuildFlagsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionBuildFlagsItems0) UnmarshalBinary(b []byte) error {
	var res IngredientVersionBuildFlagsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IngredientVersionBuildFlagsItems0DependencyGeneratorItems0 ingredient version build flags items0 dependency generator items0
// swagger:model IngredientVersionBuildFlagsItems0DependencyGeneratorItems0
type IngredientVersionBuildFlagsItems0DependencyGeneratorItems0 struct {

	// When this is true the dependency is added.
	// Required: true
	Condition *string `json:"condition"`

	// DSL code to generate the dependency that is needed. This is expected to return a list of paired names and version specs.
	// Required: true
	Dependencies *string `json:"dependencies"`
}

// Validate validates this ingredient version build flags items0 dependency generator items0
func (m *IngredientVersionBuildFlagsItems0DependencyGeneratorItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngredientVersionBuildFlagsItems0DependencyGeneratorItems0) validateCondition(formats strfmt.Registry) error {

	if err := validate.Required("condition", "body", m.Condition); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersionBuildFlagsItems0DependencyGeneratorItems0) validateDependencies(formats strfmt.Registry) error {

	if err := validate.Required("dependencies", "body", m.Dependencies); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionBuildFlagsItems0DependencyGeneratorItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionBuildFlagsItems0DependencyGeneratorItems0) UnmarshalBinary(b []byte) error {
	var res IngredientVersionBuildFlagsItems0DependencyGeneratorItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IngredientVersionBuildFlagsItems0InheritanceStrategy The strategy describes how this flag is inherited from other build ingredients. If this is not set then the flag is entirely independent.
// swagger:model IngredientVersionBuildFlagsItems0InheritanceStrategy
type IngredientVersionBuildFlagsItems0InheritanceStrategy struct {

	// The name of the dependency from which to get the flag value when the strategy is 'dependency'.
	DependencyName string `json:"dependency_name,omitempty"`

	// If the strategy is 'language_core' then it is derived by looking for a flag of the same name in the language core build flags. If the strategy is 'dependency' then it looks at the value of a dependency matching a certain name in the build. If this is omitted then it is set unconditionally based on the recipe's settings for this ingredient version
	// Required: true
	// Enum: [dependency language_core]
	Strategy *string `json:"strategy"`
}

// Validate validates this ingredient version build flags items0 inheritance strategy
func (m *IngredientVersionBuildFlagsItems0InheritanceStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ingredientVersionBuildFlagsItems0InheritanceStrategyTypeStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dependency","language_core"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionBuildFlagsItems0InheritanceStrategyTypeStrategyPropEnum = append(ingredientVersionBuildFlagsItems0InheritanceStrategyTypeStrategyPropEnum, v)
	}
}

const (

	// IngredientVersionBuildFlagsItems0InheritanceStrategyStrategyDependency captures enum value "dependency"
	IngredientVersionBuildFlagsItems0InheritanceStrategyStrategyDependency string = "dependency"

	// IngredientVersionBuildFlagsItems0InheritanceStrategyStrategyLanguageCore captures enum value "language_core"
	IngredientVersionBuildFlagsItems0InheritanceStrategyStrategyLanguageCore string = "language_core"
)

// prop value enum
func (m *IngredientVersionBuildFlagsItems0InheritanceStrategy) validateStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionBuildFlagsItems0InheritanceStrategyTypeStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionBuildFlagsItems0InheritanceStrategy) validateStrategy(formats strfmt.Registry) error {

	if err := validate.Required("inheritance_strategy"+"."+"strategy", "body", m.Strategy); err != nil {
		return err
	}

	// value enum
	if err := m.validateStrategyEnum("inheritance_strategy"+"."+"strategy", "body", *m.Strategy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionBuildFlagsItems0InheritanceStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionBuildFlagsItems0InheritanceStrategy) UnmarshalBinary(b []byte) error {
	var res IngredientVersionBuildFlagsItems0InheritanceStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IngredientVersionDependenciesItems0 ingredient version dependencies items0
// swagger:model IngredientVersionDependenciesItems0
type IngredientVersionDependenciesItems0 struct {

	// If this attribute is set, the dependency is only needed on platforms with the named bit width. If unset, the dependency is needed regardless of bit width.
	// Enum: [32 64]
	BitWidth string `json:"bit_width,omitempty"`

	// Whether this dependency is needed when this ingredient built or just when it's run.
	// Required: true
	// Enum: [build runtime]
	DependencyType *string `json:"dependency_type"`

	// If true, this dependency doesn't declare a new dependency but rather invalidates an existing one for the named package. Only used to invalidate dependencies for specific platforms, will always be false if 'os' and 'bit_width' are both unset.
	Negation bool `json:"negation,omitempty"`

	// If this attribute is set, the dependency is only needed on the named operating system. If unset, the dependency is needed on all operating systems.
	// Enum: [AIX HP-UX Linux macOS Solaris Windows]
	Os string `json:"os,omitempty"`

	// package name
	// Required: true
	PackageName *string `json:"package_name"`

	// The version of the dependency needed. If this version exactly matches an available version, the exact match will be used, otherwise this is treated as a minimum version. For dependency negations, this string will be empty.
	VersionRequirements string `json:"version_requirements,omitempty"`
}

// Validate validates this ingredient version dependencies items0
func (m *IngredientVersionDependenciesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBitWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ingredientVersionDependenciesItems0TypeBitWidthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["32","64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionDependenciesItems0TypeBitWidthPropEnum = append(ingredientVersionDependenciesItems0TypeBitWidthPropEnum, v)
	}
}

const (

	// IngredientVersionDependenciesItems0BitWidthNr32 captures enum value "32"
	IngredientVersionDependenciesItems0BitWidthNr32 string = "32"

	// IngredientVersionDependenciesItems0BitWidthNr64 captures enum value "64"
	IngredientVersionDependenciesItems0BitWidthNr64 string = "64"
)

// prop value enum
func (m *IngredientVersionDependenciesItems0) validateBitWidthEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionDependenciesItems0TypeBitWidthPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionDependenciesItems0) validateBitWidth(formats strfmt.Registry) error {

	if swag.IsZero(m.BitWidth) { // not required
		return nil
	}

	// value enum
	if err := m.validateBitWidthEnum("bit_width", "body", m.BitWidth); err != nil {
		return err
	}

	return nil
}

var ingredientVersionDependenciesItems0TypeDependencyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["build","runtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionDependenciesItems0TypeDependencyTypePropEnum = append(ingredientVersionDependenciesItems0TypeDependencyTypePropEnum, v)
	}
}

const (

	// IngredientVersionDependenciesItems0DependencyTypeBuild captures enum value "build"
	IngredientVersionDependenciesItems0DependencyTypeBuild string = "build"

	// IngredientVersionDependenciesItems0DependencyTypeRuntime captures enum value "runtime"
	IngredientVersionDependenciesItems0DependencyTypeRuntime string = "runtime"
)

// prop value enum
func (m *IngredientVersionDependenciesItems0) validateDependencyTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionDependenciesItems0TypeDependencyTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionDependenciesItems0) validateDependencyType(formats strfmt.Registry) error {

	if err := validate.Required("dependency_type", "body", m.DependencyType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDependencyTypeEnum("dependency_type", "body", *m.DependencyType); err != nil {
		return err
	}

	return nil
}

var ingredientVersionDependenciesItems0TypeOsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AIX","HP-UX","Linux","macOS","Solaris","Windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingredientVersionDependenciesItems0TypeOsPropEnum = append(ingredientVersionDependenciesItems0TypeOsPropEnum, v)
	}
}

const (

	// IngredientVersionDependenciesItems0OsAIX captures enum value "AIX"
	IngredientVersionDependenciesItems0OsAIX string = "AIX"

	// IngredientVersionDependenciesItems0OsHPUX captures enum value "HP-UX"
	IngredientVersionDependenciesItems0OsHPUX string = "HP-UX"

	// IngredientVersionDependenciesItems0OsLinux captures enum value "Linux"
	IngredientVersionDependenciesItems0OsLinux string = "Linux"

	// IngredientVersionDependenciesItems0OsMacOS captures enum value "macOS"
	IngredientVersionDependenciesItems0OsMacOS string = "macOS"

	// IngredientVersionDependenciesItems0OsSolaris captures enum value "Solaris"
	IngredientVersionDependenciesItems0OsSolaris string = "Solaris"

	// IngredientVersionDependenciesItems0OsWindows captures enum value "Windows"
	IngredientVersionDependenciesItems0OsWindows string = "Windows"
)

// prop value enum
func (m *IngredientVersionDependenciesItems0) validateOsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ingredientVersionDependenciesItems0TypeOsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IngredientVersionDependenciesItems0) validateOs(formats strfmt.Registry) error {

	if swag.IsZero(m.Os) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsEnum("os", "body", m.Os); err != nil {
		return err
	}

	return nil
}

func (m *IngredientVersionDependenciesItems0) validatePackageName(formats strfmt.Registry) error {

	if err := validate.Required("package_name", "body", m.PackageName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionDependenciesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionDependenciesItems0) UnmarshalBinary(b []byte) error {
	var res IngredientVersionDependenciesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IngredientVersionProvidesAnon Metadata about the provided packages, keyed by package name.
// swagger:model IngredientVersionProvidesAnon
type IngredientVersionProvidesAnon struct {

	// Whether this ingredient version is the default provider of this package.
	IsDefault bool `json:"is_default,omitempty"`

	// The version of this package provided by the ingredient version.
	Version string `json:"version,omitempty"`
}

// Validate validates this ingredient version provides anon
func (m *IngredientVersionProvidesAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionProvidesAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionProvidesAnon) UnmarshalBinary(b []byte) error {
	var res IngredientVersionProvidesAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}