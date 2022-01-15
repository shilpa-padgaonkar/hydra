// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewGetLogoutRequestParams creates a new GetLogoutRequestParams object
// with the default values initialized.
func NewGetLogoutRequestParams() *GetLogoutRequestParams {
	var ()
	return &GetLogoutRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogoutRequestParamsWithTimeout creates a new GetLogoutRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogoutRequestParamsWithTimeout(timeout time.Duration) *GetLogoutRequestParams {
	var ()
	return &GetLogoutRequestParams{

		timeout: timeout,
	}
}

// NewGetLogoutRequestParamsWithContext creates a new GetLogoutRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogoutRequestParamsWithContext(ctx context.Context) *GetLogoutRequestParams {
	var ()
	return &GetLogoutRequestParams{

		Context: ctx,
	}
}

// NewGetLogoutRequestParamsWithHTTPClient creates a new GetLogoutRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogoutRequestParamsWithHTTPClient(client *http.Client) *GetLogoutRequestParams {
	var ()
	return &GetLogoutRequestParams{
		HTTPClient: client,
	}
}

/*GetLogoutRequestParams contains all the parameters to send to the API endpoint
for the get logout request operation typically these are written to a http.Request
*/
type GetLogoutRequestParams struct {

	/*LogoutChallenge*/
	LogoutChallenge string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logout request params
func (o *GetLogoutRequestParams) WithTimeout(timeout time.Duration) *GetLogoutRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logout request params
func (o *GetLogoutRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logout request params
func (o *GetLogoutRequestParams) WithContext(ctx context.Context) *GetLogoutRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logout request params
func (o *GetLogoutRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logout request params
func (o *GetLogoutRequestParams) WithHTTPClient(client *http.Client) *GetLogoutRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logout request params
func (o *GetLogoutRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogoutChallenge adds the logoutChallenge to the get logout request params
func (o *GetLogoutRequestParams) WithLogoutChallenge(logoutChallenge string) *GetLogoutRequestParams {
	o.SetLogoutChallenge(logoutChallenge)
	return o
}

// SetLogoutChallenge adds the logoutChallenge to the get logout request params
func (o *GetLogoutRequestParams) SetLogoutChallenge(logoutChallenge string) {
	o.LogoutChallenge = logoutChallenge
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogoutRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param logout_challenge
	qrLogoutChallenge := o.LogoutChallenge
	qLogoutChallenge := qrLogoutChallenge
	if qLogoutChallenge != "" {
		if err := r.SetQueryParam("logout_challenge", qLogoutChallenge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
