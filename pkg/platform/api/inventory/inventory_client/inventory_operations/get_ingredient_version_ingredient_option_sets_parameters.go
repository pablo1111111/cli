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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetIngredientVersionIngredientOptionSetsParams creates a new GetIngredientVersionIngredientOptionSetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIngredientVersionIngredientOptionSetsParams() *GetIngredientVersionIngredientOptionSetsParams {
	return &GetIngredientVersionIngredientOptionSetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIngredientVersionIngredientOptionSetsParamsWithTimeout creates a new GetIngredientVersionIngredientOptionSetsParams object
// with the ability to set a timeout on a request.
func NewGetIngredientVersionIngredientOptionSetsParamsWithTimeout(timeout time.Duration) *GetIngredientVersionIngredientOptionSetsParams {
	return &GetIngredientVersionIngredientOptionSetsParams{
		timeout: timeout,
	}
}

// NewGetIngredientVersionIngredientOptionSetsParamsWithContext creates a new GetIngredientVersionIngredientOptionSetsParams object
// with the ability to set a context for a request.
func NewGetIngredientVersionIngredientOptionSetsParamsWithContext(ctx context.Context) *GetIngredientVersionIngredientOptionSetsParams {
	return &GetIngredientVersionIngredientOptionSetsParams{
		Context: ctx,
	}
}

// NewGetIngredientVersionIngredientOptionSetsParamsWithHTTPClient creates a new GetIngredientVersionIngredientOptionSetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIngredientVersionIngredientOptionSetsParamsWithHTTPClient(client *http.Client) *GetIngredientVersionIngredientOptionSetsParams {
	return &GetIngredientVersionIngredientOptionSetsParams{
		HTTPClient: client,
	}
}

/* GetIngredientVersionIngredientOptionSetsParams contains all the parameters to send to the API endpoint
   for the get ingredient version ingredient option sets operation.

   Typically these are written to a http.Request.
*/
type GetIngredientVersionIngredientOptionSetsParams struct {

	/* AllowUnstable.

	   Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version
	*/
	AllowUnstable *bool

	// IngredientID.
	//
	// Format: uuid
	IngredientID strfmt.UUID

	// IngredientVersionID.
	//
	// Format: uuid
	IngredientVersionID strfmt.UUID

	/* Limit.

	   The maximum number of items returned per page

	   Default: 50
	*/
	Limit *int64

	/* Page.

	   The page number returned

	   Default: 1
	*/
	Page *int64

	/* StateAt.

	   Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	   Format: date-time
	*/
	StateAt *strfmt.DateTime

	/* UsageType.

	   Filter to just ingredient option sets used by the ingredient version as a default or just those used as an override
	*/
	UsageType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ingredient version ingredient option sets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIngredientVersionIngredientOptionSetsParams) WithDefaults() *GetIngredientVersionIngredientOptionSetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ingredient version ingredient option sets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIngredientVersionIngredientOptionSetsParams) SetDefaults() {
	var (
		allowUnstableDefault = bool(false)

		limitDefault = int64(50)

		pageDefault = int64(1)
	)

	val := GetIngredientVersionIngredientOptionSetsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithTimeout(timeout time.Duration) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithContext(ctx context.Context) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithHTTPClient(client *http.Client) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithAllowUnstable(allowUnstable *bool) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithIngredientID adds the ingredientID to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithIngredientID(ingredientID strfmt.UUID) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithIngredientVersionID adds the ingredientVersionID to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithIngredientVersionID(ingredientVersionID strfmt.UUID) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetIngredientVersionID(ingredientVersionID)
	return o
}

// SetIngredientVersionID adds the ingredientVersionId to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetIngredientVersionID(ingredientVersionID strfmt.UUID) {
	o.IngredientVersionID = ingredientVersionID
}

// WithLimit adds the limit to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithLimit(limit *int64) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithPage(page *int64) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithStateAt(stateAt *strfmt.DateTime) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WithUsageType adds the usageType to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) WithUsageType(usageType *string) *GetIngredientVersionIngredientOptionSetsParams {
	o.SetUsageType(usageType)
	return o
}

// SetUsageType adds the usageType to the get ingredient version ingredient option sets params
func (o *GetIngredientVersionIngredientOptionSetsParams) SetUsageType(usageType *string) {
	o.UsageType = usageType
}

// WriteToRequest writes these params to a swagger request
func (o *GetIngredientVersionIngredientOptionSetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool

		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {

			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}
	}

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
	}

	// path param ingredient_version_id
	if err := r.SetPathParam("ingredient_version_id", o.IngredientVersionID.String()); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime

		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {

			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}
	}

	if o.UsageType != nil {

		// query param usage_type
		var qrUsageType string

		if o.UsageType != nil {
			qrUsageType = *o.UsageType
		}
		qUsageType := qrUsageType
		if qUsageType != "" {

			if err := r.SetQueryParam("usage_type", qUsageType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}