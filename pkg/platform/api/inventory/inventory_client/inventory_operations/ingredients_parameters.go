// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

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

// NewIngredientsParams creates a new IngredientsParams object
// with the default values initialized.
func NewIngredientsParams() *IngredientsParams {
	var ()
	return &IngredientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIngredientsParamsWithTimeout creates a new IngredientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIngredientsParamsWithTimeout(timeout time.Duration) *IngredientsParams {
	var ()
	return &IngredientsParams{

		timeout: timeout,
	}
}

// NewIngredientsParamsWithContext creates a new IngredientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIngredientsParamsWithContext(ctx context.Context) *IngredientsParams {
	var ()
	return &IngredientsParams{

		Context: ctx,
	}
}

// NewIngredientsParamsWithHTTPClient creates a new IngredientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIngredientsParamsWithHTTPClient(client *http.Client) *IngredientsParams {
	var ()
	return &IngredientsParams{
		HTTPClient: client,
	}
}

/*IngredientsParams contains all the parameters to send to the API endpoint
for the ingredients operation typically these are written to a http.Request
*/
type IngredientsParams struct {

	/*PackageName*/
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ingredients params
func (o *IngredientsParams) WithTimeout(timeout time.Duration) *IngredientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ingredients params
func (o *IngredientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ingredients params
func (o *IngredientsParams) WithContext(ctx context.Context) *IngredientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ingredients params
func (o *IngredientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ingredients params
func (o *IngredientsParams) WithHTTPClient(client *http.Client) *IngredientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ingredients params
func (o *IngredientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the ingredients params
func (o *IngredientsParams) WithPackageName(packageName string) *IngredientsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the ingredients params
func (o *IngredientsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *IngredientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param package_name
	qrPackageName := o.PackageName
	qPackageName := qrPackageName
	if qPackageName != "" {
		if err := r.SetQueryParam("package_name", qPackageName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}