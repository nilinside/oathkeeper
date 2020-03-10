// Passcode generated by go-swagger; DO NOT EDIT.

package api

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

// NewDecisionsParams creates a new DecisionsParams object
// with the default values initialized.
func NewDecisionsParams() *DecisionsParams {

	return &DecisionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDecisionsParamsWithTimeout creates a new DecisionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDecisionsParamsWithTimeout(timeout time.Duration) *DecisionsParams {

	return &DecisionsParams{

		timeout: timeout,
	}
}

// NewDecisionsParamsWithContext creates a new DecisionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDecisionsParamsWithContext(ctx context.Context) *DecisionsParams {

	return &DecisionsParams{

		Context: ctx,
	}
}

// NewDecisionsParamsWithHTTPClient creates a new DecisionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDecisionsParamsWithHTTPClient(client *http.Client) *DecisionsParams {

	return &DecisionsParams{
		HTTPClient: client,
	}
}

/*DecisionsParams contains all the parameters to send to the API endpoint
for the decisions operation typically these are written to a http.Request
*/
type DecisionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the decisions params
func (o *DecisionsParams) WithTimeout(timeout time.Duration) *DecisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decisions params
func (o *DecisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decisions params
func (o *DecisionsParams) WithContext(ctx context.Context) *DecisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decisions params
func (o *DecisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decisions params
func (o *DecisionsParams) WithHTTPClient(client *http.Client) *DecisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decisions params
func (o *DecisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DecisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
