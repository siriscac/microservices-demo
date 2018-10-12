// Code generated by go-swagger; DO NOT EDIT.

package recommendation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListRecommendationsParams creates a new ListRecommendationsParams object
// with the default values initialized.
func NewListRecommendationsParams() *ListRecommendationsParams {
	var ()
	return &ListRecommendationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRecommendationsParamsWithTimeout creates a new ListRecommendationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRecommendationsParamsWithTimeout(timeout time.Duration) *ListRecommendationsParams {
	var ()
	return &ListRecommendationsParams{

		timeout: timeout,
	}
}

// NewListRecommendationsParamsWithContext creates a new ListRecommendationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRecommendationsParamsWithContext(ctx context.Context) *ListRecommendationsParams {
	var ()
	return &ListRecommendationsParams{

		Context: ctx,
	}
}

// NewListRecommendationsParamsWithHTTPClient creates a new ListRecommendationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRecommendationsParamsWithHTTPClient(client *http.Client) *ListRecommendationsParams {
	var ()
	return &ListRecommendationsParams{
		HTTPClient: client,
	}
}

/*ListRecommendationsParams contains all the parameters to send to the API endpoint
for the list recommendations operation typically these are written to a http.Request
*/
type ListRecommendationsParams struct {

	/*ProductIds*/
	ProductIds []string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list recommendations params
func (o *ListRecommendationsParams) WithTimeout(timeout time.Duration) *ListRecommendationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list recommendations params
func (o *ListRecommendationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list recommendations params
func (o *ListRecommendationsParams) WithContext(ctx context.Context) *ListRecommendationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list recommendations params
func (o *ListRecommendationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list recommendations params
func (o *ListRecommendationsParams) WithHTTPClient(client *http.Client) *ListRecommendationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list recommendations params
func (o *ListRecommendationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductIds adds the productIds to the list recommendations params
func (o *ListRecommendationsParams) WithProductIds(productIds []string) *ListRecommendationsParams {
	o.SetProductIds(productIds)
	return o
}

// SetProductIds adds the productIds to the list recommendations params
func (o *ListRecommendationsParams) SetProductIds(productIds []string) {
	o.ProductIds = productIds
}

// WithUserID adds the userID to the list recommendations params
func (o *ListRecommendationsParams) WithUserID(userID string) *ListRecommendationsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the list recommendations params
func (o *ListRecommendationsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRecommendationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesProductIds := o.ProductIds

	joinedProductIds := swag.JoinByFormat(valuesProductIds, "")
	// query array param product_ids
	if err := r.SetQueryParam("product_ids", joinedProductIds...); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
