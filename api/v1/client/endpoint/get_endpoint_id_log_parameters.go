// Code generated by go-swagger; DO NOT EDIT.

package endpoint

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

// NewGetEndpointIDLogParams creates a new GetEndpointIDLogParams object
// with the default values initialized.
func NewGetEndpointIDLogParams() *GetEndpointIDLogParams {
	var ()
	return &GetEndpointIDLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointIDLogParamsWithTimeout creates a new GetEndpointIDLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEndpointIDLogParamsWithTimeout(timeout time.Duration) *GetEndpointIDLogParams {
	var ()
	return &GetEndpointIDLogParams{

		timeout: timeout,
	}
}

// NewGetEndpointIDLogParamsWithContext creates a new GetEndpointIDLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEndpointIDLogParamsWithContext(ctx context.Context) *GetEndpointIDLogParams {
	var ()
	return &GetEndpointIDLogParams{

		Context: ctx,
	}
}

// NewGetEndpointIDLogParamsWithHTTPClient creates a new GetEndpointIDLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEndpointIDLogParamsWithHTTPClient(client *http.Client) *GetEndpointIDLogParams {
	var ()
	return &GetEndpointIDLogParams{
		HTTPClient: client,
	}
}

/*GetEndpointIDLogParams contains all the parameters to send to the API endpoint
for the get endpoint ID log operation typically these are written to a http.Request
*/
type GetEndpointIDLogParams struct {

	/*ID
	  String describing an endpoint with the format ``[prefix:]id``. If no prefix
	is specified, a prefix of ``cilium-local:`` is assumed. Not all endpoints
	will be addressable by all endpoint ID prefixes with the exception of the
	local Cilium UUID which is assigned to all endpoints.

	Supported endpoint id prefixes:
	  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595
	  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343
	  - container-id: Container runtime ID, e.g. container-id:22222
	  - container-name: Container name, e.g. container-name:foobar
	  - pod-name: pod name for this container if K8s is enabled, e.g. pod-name:default:foobar
	  - docker-endpoint: Docker libnetwork endpoint ID, e.g. docker-endpoint:4444


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get endpoint ID log params
func (o *GetEndpointIDLogParams) WithTimeout(timeout time.Duration) *GetEndpointIDLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint ID log params
func (o *GetEndpointIDLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint ID log params
func (o *GetEndpointIDLogParams) WithContext(ctx context.Context) *GetEndpointIDLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint ID log params
func (o *GetEndpointIDLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint ID log params
func (o *GetEndpointIDLogParams) WithHTTPClient(client *http.Client) *GetEndpointIDLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint ID log params
func (o *GetEndpointIDLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get endpoint ID log params
func (o *GetEndpointIDLogParams) WithID(id string) *GetEndpointIDLogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get endpoint ID log params
func (o *GetEndpointIDLogParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointIDLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
