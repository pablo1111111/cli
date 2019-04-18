// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRequestResetParams creates a new RequestResetParams object
// with the default values initialized.
func NewRequestResetParams() *RequestResetParams {
	var ()
	return &RequestResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestResetParamsWithTimeout creates a new RequestResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestResetParamsWithTimeout(timeout time.Duration) *RequestResetParams {
	var ()
	return &RequestResetParams{

		timeout: timeout,
	}
}

// NewRequestResetParamsWithContext creates a new RequestResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestResetParamsWithContext(ctx context.Context) *RequestResetParams {
	var ()
	return &RequestResetParams{

		Context: ctx,
	}
}

// NewRequestResetParamsWithHTTPClient creates a new RequestResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestResetParamsWithHTTPClient(client *http.Client) *RequestResetParams {
	var ()
	return &RequestResetParams{
		HTTPClient: client,
	}
}

/*RequestResetParams contains all the parameters to send to the API endpoint
for the request reset operation typically these are written to a http.Request
*/
type RequestResetParams struct {

	/*Email
	  User email address

	*/
	Email string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request reset params
func (o *RequestResetParams) WithTimeout(timeout time.Duration) *RequestResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request reset params
func (o *RequestResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request reset params
func (o *RequestResetParams) WithContext(ctx context.Context) *RequestResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request reset params
func (o *RequestResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request reset params
func (o *RequestResetParams) WithHTTPClient(client *http.Client) *RequestResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request reset params
func (o *RequestResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the request reset params
func (o *RequestResetParams) WithEmail(email string) *RequestResetParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the request reset params
func (o *RequestResetParams) SetEmail(email string) {
	o.Email = email
}

// WriteToRequest writes these params to a swagger request
func (o *RequestResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}