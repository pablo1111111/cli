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

// V2BuildRequestRequester Requester
//
// Identifying information about who is placing the build request
// swagger:model v2BuildRequestRequester
type V2BuildRequestRequester struct {

	// The UUID of the platform organization that owns the project being built
	// Required: true
	// Format: uuid
	OrganizationID *strfmt.UUID `json:"organization_id"`

	// The UUID of the platform project being built by this build request
	// Required: true
	// Format: uuid
	ProjectID *strfmt.UUID `json:"project_id"`

	// An optional string describing the service for non user initiated requests
	Service string `json:"service,omitempty"`

	// The UUID of the platform user who initiated the build request
	// Format: uuid
	UserID strfmt.UUID `json:"user_id,omitempty"`
}

// Validate validates this v2 build request requester
func (m *V2BuildRequestRequester) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2BuildRequestRequester) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organization_id", "body", m.OrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organization_id", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2BuildRequestRequester) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("project_id", "body", m.ProjectID); err != nil {
		return err
	}

	if err := validate.FormatOf("project_id", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2BuildRequestRequester) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2BuildRequestRequester) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2BuildRequestRequester) UnmarshalBinary(b []byte) error {
	var res V2BuildRequestRequester
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}