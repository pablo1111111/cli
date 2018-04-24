// Code generated by go-swagger; DO NOT EDIT.

package licenses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLicenseParams creates a new GetLicenseParams object
// with the default values initialized.
func NewGetLicenseParams() *GetLicenseParams {
	var ()
	return &GetLicenseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseParamsWithTimeout creates a new GetLicenseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLicenseParamsWithTimeout(timeout time.Duration) *GetLicenseParams {
	var ()
	return &GetLicenseParams{

		timeout: timeout,
	}
}

// NewGetLicenseParamsWithContext creates a new GetLicenseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLicenseParamsWithContext(ctx context.Context) *GetLicenseParams {
	var ()
	return &GetLicenseParams{

		Context: ctx,
	}
}

// NewGetLicenseParamsWithHTTPClient creates a new GetLicenseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLicenseParamsWithHTTPClient(client *http.Client) *GetLicenseParams {
	var ()
	return &GetLicenseParams{
		HTTPClient: client,
	}
}

/*GetLicenseParams contains all the parameters to send to the API endpoint
for the get license operation typically these are written to a http.Request
*/
type GetLicenseParams struct {

	/*LicenseID
	  licenseID of desired license

	*/
	LicenseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get license params
func (o *GetLicenseParams) WithTimeout(timeout time.Duration) *GetLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license params
func (o *GetLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license params
func (o *GetLicenseParams) WithContext(ctx context.Context) *GetLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license params
func (o *GetLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license params
func (o *GetLicenseParams) WithHTTPClient(client *http.Client) *GetLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license params
func (o *GetLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLicenseID adds the licenseID to the get license params
func (o *GetLicenseParams) WithLicenseID(licenseID strfmt.UUID) *GetLicenseParams {
	o.SetLicenseID(licenseID)
	return o
}

// SetLicenseID adds the licenseId to the get license params
func (o *GetLicenseParams) SetLicenseID(licenseID strfmt.UUID) {
	o.LicenseID = licenseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param licenseID
	if err := r.SetPathParam("licenseID", o.LicenseID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}