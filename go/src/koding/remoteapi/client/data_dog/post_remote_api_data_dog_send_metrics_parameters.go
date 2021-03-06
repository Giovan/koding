package data_dog

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPIDataDogSendMetricsParams creates a new PostRemoteAPIDataDogSendMetricsParams object
// with the default values initialized.
func NewPostRemoteAPIDataDogSendMetricsParams() *PostRemoteAPIDataDogSendMetricsParams {
	var ()
	return &PostRemoteAPIDataDogSendMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIDataDogSendMetricsParamsWithTimeout creates a new PostRemoteAPIDataDogSendMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIDataDogSendMetricsParamsWithTimeout(timeout time.Duration) *PostRemoteAPIDataDogSendMetricsParams {
	var ()
	return &PostRemoteAPIDataDogSendMetricsParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIDataDogSendMetricsParamsWithContext creates a new PostRemoteAPIDataDogSendMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIDataDogSendMetricsParamsWithContext(ctx context.Context) *PostRemoteAPIDataDogSendMetricsParams {
	var ()
	return &PostRemoteAPIDataDogSendMetricsParams{

		Context: ctx,
	}
}

/*PostRemoteAPIDataDogSendMetricsParams contains all the parameters to send to the API endpoint
for the post remote API data dog send metrics operation typically these are written to a http.Request
*/
type PostRemoteAPIDataDogSendMetricsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API data dog send metrics params
func (o *PostRemoteAPIDataDogSendMetricsParams) WithTimeout(timeout time.Duration) *PostRemoteAPIDataDogSendMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API data dog send metrics params
func (o *PostRemoteAPIDataDogSendMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API data dog send metrics params
func (o *PostRemoteAPIDataDogSendMetricsParams) WithContext(ctx context.Context) *PostRemoteAPIDataDogSendMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API data dog send metrics params
func (o *PostRemoteAPIDataDogSendMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API data dog send metrics params
func (o *PostRemoteAPIDataDogSendMetricsParams) WithBody(body models.DefaultSelector) *PostRemoteAPIDataDogSendMetricsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API data dog send metrics params
func (o *PostRemoteAPIDataDogSendMetricsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIDataDogSendMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
