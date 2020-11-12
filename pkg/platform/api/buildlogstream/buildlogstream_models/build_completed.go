// Code generated by go-swagger; DO NOT EDIT.

package buildlogstream_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildCompleted Build Completed
//
// A message indicating that a requested build has been completed.
//
// swagger:model BuildCompleted
type BuildCompleted struct {

	// An array containing all the artifacts that make up this build.
	// Required: true
	Artifacts []*Artifact `json:"artifacts"`

	// An S3 URI containing the log for this build.
	// Format: uri
	LogURI strfmt.URI `json:"log_uri,omitempty"`

	// A user-facing message describing the build results.
	Message string `json:"message,omitempty"`
}

// Validate validates this build completed
func (m *BuildCompleted) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildCompleted) validateArtifacts(formats strfmt.Registry) error {

	if err := validate.Required("artifacts", "body", m.Artifacts); err != nil {
		return err
	}

	for i := 0; i < len(m.Artifacts); i++ {
		if swag.IsZero(m.Artifacts[i]) { // not required
			continue
		}

		if m.Artifacts[i] != nil {
			if err := m.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildCompleted) validateLogURI(formats strfmt.Registry) error {

	if swag.IsZero(m.LogURI) { // not required
		return nil
	}

	if err := validate.FormatOf("log_uri", "body", "uri", m.LogURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildCompleted) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildCompleted) UnmarshalBinary(b []byte) error {
	var res BuildCompleted
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}