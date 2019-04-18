// Code generated by go-swagger; DO NOT EDIT.

package ingredients

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

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewAddIngredientParams creates a new AddIngredientParams object
// with the default values initialized.
func NewAddIngredientParams() *AddIngredientParams {
	var ()
	return &AddIngredientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngredientParamsWithTimeout creates a new AddIngredientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddIngredientParamsWithTimeout(timeout time.Duration) *AddIngredientParams {
	var ()
	return &AddIngredientParams{

		timeout: timeout,
	}
}

// NewAddIngredientParamsWithContext creates a new AddIngredientParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddIngredientParamsWithContext(ctx context.Context) *AddIngredientParams {
	var ()
	return &AddIngredientParams{

		Context: ctx,
	}
}

// NewAddIngredientParamsWithHTTPClient creates a new AddIngredientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddIngredientParamsWithHTTPClient(client *http.Client) *AddIngredientParams {
	var ()
	return &AddIngredientParams{
		HTTPClient: client,
	}
}

/*AddIngredientParams contains all the parameters to send to the API endpoint
for the add ingredient operation typically these are written to a http.Request
*/
type AddIngredientParams struct {

	/*Ingredient
	  Ingredient to add

	*/
	Ingredient *mono_models.Ingredient

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add ingredient params
func (o *AddIngredientParams) WithTimeout(timeout time.Duration) *AddIngredientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingredient params
func (o *AddIngredientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingredient params
func (o *AddIngredientParams) WithContext(ctx context.Context) *AddIngredientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingredient params
func (o *AddIngredientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingredient params
func (o *AddIngredientParams) WithHTTPClient(client *http.Client) *AddIngredientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingredient params
func (o *AddIngredientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredient adds the ingredient to the add ingredient params
func (o *AddIngredientParams) WithIngredient(ingredient *mono_models.Ingredient) *AddIngredientParams {
	o.SetIngredient(ingredient)
	return o
}

// SetIngredient adds the ingredient to the add ingredient params
func (o *AddIngredientParams) SetIngredient(ingredient *mono_models.Ingredient) {
	o.Ingredient = ingredient
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngredientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ingredient != nil {
		if err := r.SetBodyParam(o.Ingredient); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}