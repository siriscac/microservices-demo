// Code generated by go-swagger; DO NOT EDIT.

package product_catalog_service

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

// NewGetProductParams creates a new GetProductParams object
// with the default values initialized.
func NewGetProductParams() *GetProductParams {
	var ()
	return &GetProductParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductParamsWithTimeout creates a new GetProductParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductParamsWithTimeout(timeout time.Duration) *GetProductParams {
	var ()
	return &GetProductParams{

		timeout: timeout,
	}
}

// NewGetProductParamsWithContext creates a new GetProductParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductParamsWithContext(ctx context.Context) *GetProductParams {
	var ()
	return &GetProductParams{

		Context: ctx,
	}
}

// NewGetProductParamsWithHTTPClient creates a new GetProductParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductParamsWithHTTPClient(client *http.Client) *GetProductParams {
	var ()
	return &GetProductParams{
		HTTPClient: client,
	}
}

/*GetProductParams contains all the parameters to send to the API endpoint
for the get product operation typically these are written to a http.Request
*/
type GetProductParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get product params
func (o *GetProductParams) WithTimeout(timeout time.Duration) *GetProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product params
func (o *GetProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product params
func (o *GetProductParams) WithContext(ctx context.Context) *GetProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product params
func (o *GetProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product params
func (o *GetProductParams) WithHTTPClient(client *http.Client) *GetProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product params
func (o *GetProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get product params
func (o *GetProductParams) WithID(id string) *GetProductParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get product params
func (o *GetProductParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
