// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

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

	"github.com/cilium/cilium/api/v1/models"
)

// NewGetPolicyParams creates a new GetPolicyParams object
// with the default values initialized.
func NewGetPolicyParams() *GetPolicyParams {
	var ()
	return &GetPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyParamsWithTimeout creates a new GetPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyParamsWithTimeout(timeout time.Duration) *GetPolicyParams {
	var ()
	return &GetPolicyParams{

		timeout: timeout,
	}
}

// NewGetPolicyParamsWithContext creates a new GetPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyParamsWithContext(ctx context.Context) *GetPolicyParams {
	var ()
	return &GetPolicyParams{

		Context: ctx,
	}
}

// NewGetPolicyParamsWithHTTPClient creates a new GetPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyParamsWithHTTPClient(client *http.Client) *GetPolicyParams {
	var ()
	return &GetPolicyParams{
		HTTPClient: client,
	}
}

/*GetPolicyParams contains all the parameters to send to the API endpoint
for the get policy operation typically these are written to a http.Request
*/
type GetPolicyParams struct {

	/*Labels*/
	Labels models.Labels

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy params
func (o *GetPolicyParams) WithTimeout(timeout time.Duration) *GetPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy params
func (o *GetPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy params
func (o *GetPolicyParams) WithContext(ctx context.Context) *GetPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy params
func (o *GetPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy params
func (o *GetPolicyParams) WithHTTPClient(client *http.Client) *GetPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy params
func (o *GetPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabels adds the labels to the get policy params
func (o *GetPolicyParams) WithLabels(labels models.Labels) *GetPolicyParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the get policy params
func (o *GetPolicyParams) SetLabels(labels models.Labels) {
	o.Labels = labels
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Labels != nil {
		if err := r.SetBodyParam(o.Labels); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
