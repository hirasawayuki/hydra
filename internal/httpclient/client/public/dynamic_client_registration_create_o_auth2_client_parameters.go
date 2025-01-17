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

	"github.com/ory/hydra/internal/httpclient/models"
)

// NewDynamicClientRegistrationCreateOAuth2ClientParams creates a new DynamicClientRegistrationCreateOAuth2ClientParams object
// with the default values initialized.
func NewDynamicClientRegistrationCreateOAuth2ClientParams() *DynamicClientRegistrationCreateOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationCreateOAuth2ClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDynamicClientRegistrationCreateOAuth2ClientParamsWithTimeout creates a new DynamicClientRegistrationCreateOAuth2ClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDynamicClientRegistrationCreateOAuth2ClientParamsWithTimeout(timeout time.Duration) *DynamicClientRegistrationCreateOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationCreateOAuth2ClientParams{

		timeout: timeout,
	}
}

// NewDynamicClientRegistrationCreateOAuth2ClientParamsWithContext creates a new DynamicClientRegistrationCreateOAuth2ClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewDynamicClientRegistrationCreateOAuth2ClientParamsWithContext(ctx context.Context) *DynamicClientRegistrationCreateOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationCreateOAuth2ClientParams{

		Context: ctx,
	}
}

// NewDynamicClientRegistrationCreateOAuth2ClientParamsWithHTTPClient creates a new DynamicClientRegistrationCreateOAuth2ClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDynamicClientRegistrationCreateOAuth2ClientParamsWithHTTPClient(client *http.Client) *DynamicClientRegistrationCreateOAuth2ClientParams {
	var ()
	return &DynamicClientRegistrationCreateOAuth2ClientParams{
		HTTPClient: client,
	}
}

/*DynamicClientRegistrationCreateOAuth2ClientParams contains all the parameters to send to the API endpoint
for the dynamic client registration create o auth2 client operation typically these are written to a http.Request
*/
type DynamicClientRegistrationCreateOAuth2ClientParams struct {

	/*Body*/
	Body *models.OAuth2Client

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) WithTimeout(timeout time.Duration) *DynamicClientRegistrationCreateOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) WithContext(ctx context.Context) *DynamicClientRegistrationCreateOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) WithHTTPClient(client *http.Client) *DynamicClientRegistrationCreateOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) WithBody(body *models.OAuth2Client) *DynamicClientRegistrationCreateOAuth2ClientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the dynamic client registration create o auth2 client params
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) SetBody(body *models.OAuth2Client) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DynamicClientRegistrationCreateOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
