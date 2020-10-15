// Code generated by go-swagger; DO NOT EDIT.

package public

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
)

// NewGetSelfServiceSettingsFlowParams creates a new GetSelfServiceSettingsFlowParams object
// with the default values initialized.
func NewGetSelfServiceSettingsFlowParams() *GetSelfServiceSettingsFlowParams {
	var ()
	return &GetSelfServiceSettingsFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSelfServiceSettingsFlowParamsWithTimeout creates a new GetSelfServiceSettingsFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSelfServiceSettingsFlowParamsWithTimeout(timeout time.Duration) *GetSelfServiceSettingsFlowParams {
	var ()
	return &GetSelfServiceSettingsFlowParams{

		timeout: timeout,
	}
}

// NewGetSelfServiceSettingsFlowParamsWithContext creates a new GetSelfServiceSettingsFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSelfServiceSettingsFlowParamsWithContext(ctx context.Context) *GetSelfServiceSettingsFlowParams {
	var ()
	return &GetSelfServiceSettingsFlowParams{

		Context: ctx,
	}
}

// NewGetSelfServiceSettingsFlowParamsWithHTTPClient creates a new GetSelfServiceSettingsFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSelfServiceSettingsFlowParamsWithHTTPClient(client *http.Client) *GetSelfServiceSettingsFlowParams {
	var ()
	return &GetSelfServiceSettingsFlowParams{
		HTTPClient: client,
	}
}

/*GetSelfServiceSettingsFlowParams contains all the parameters to send to the API endpoint
for the get self service settings flow operation typically these are written to a http.Request
*/
type GetSelfServiceSettingsFlowParams struct {

	/*ID
	  ID is the Settings Flow ID

	The value for this parameter comes from `flow` URL Query parameter sent to your
	application (e.g. `/settings?flow=abcde`).

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) WithTimeout(timeout time.Duration) *GetSelfServiceSettingsFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) WithContext(ctx context.Context) *GetSelfServiceSettingsFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) WithHTTPClient(client *http.Client) *GetSelfServiceSettingsFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) WithID(id string) *GetSelfServiceSettingsFlowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get self service settings flow params
func (o *GetSelfServiceSettingsFlowParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSelfServiceSettingsFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {
		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
