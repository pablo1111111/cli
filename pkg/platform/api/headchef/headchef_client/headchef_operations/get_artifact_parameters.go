// Code generated by go-swagger; DO NOT EDIT.

package headchef_operations

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

// NewGetArtifactParams creates a new GetArtifactParams object
// with the default values initialized.
func NewGetArtifactParams() *GetArtifactParams {
	var ()
	return &GetArtifactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactParamsWithTimeout creates a new GetArtifactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArtifactParamsWithTimeout(timeout time.Duration) *GetArtifactParams {
	var ()
	return &GetArtifactParams{

		timeout: timeout,
	}
}

// NewGetArtifactParamsWithContext creates a new GetArtifactParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArtifactParamsWithContext(ctx context.Context) *GetArtifactParams {
	var ()
	return &GetArtifactParams{

		Context: ctx,
	}
}

// NewGetArtifactParamsWithHTTPClient creates a new GetArtifactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArtifactParamsWithHTTPClient(client *http.Client) *GetArtifactParams {
	var ()
	return &GetArtifactParams{
		HTTPClient: client,
	}
}

/*GetArtifactParams contains all the parameters to send to the API endpoint
for the get artifact operation typically these are written to a http.Request
*/
type GetArtifactParams struct {

	/*ArtifactID*/
	ArtifactID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get artifact params
func (o *GetArtifactParams) WithTimeout(timeout time.Duration) *GetArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact params
func (o *GetArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact params
func (o *GetArtifactParams) WithContext(ctx context.Context) *GetArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact params
func (o *GetArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact params
func (o *GetArtifactParams) WithHTTPClient(client *http.Client) *GetArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact params
func (o *GetArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the get artifact params
func (o *GetArtifactParams) WithArtifactID(artifactID strfmt.UUID) *GetArtifactParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the get artifact params
func (o *GetArtifactParams) SetArtifactID(artifactID strfmt.UUID) {
	o.ArtifactID = artifactID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifact_id
	if err := r.SetPathParam("artifact_id", o.ArtifactID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}