// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EndpointHealthStatus A common set of statuses for endpoint health * ``OK`` = All components operational * ``Bootstrap`` = This component is being created * ``Pending`` = A change is being processed to be applied * ``Warning`` = This component is not applying up-to-date policies (but is still applying the previous version) * ``Failure`` = An error has occurred and no policy is being applied * ``Disabled`` = This endpoint is disabled and will not handle traffic
//
//
// swagger:model EndpointHealthStatus
type EndpointHealthStatus string

const (

	// EndpointHealthStatusOK captures enum value "OK"
	EndpointHealthStatusOK EndpointHealthStatus = "OK"

	// EndpointHealthStatusBootstrap captures enum value "Bootstrap"
	EndpointHealthStatusBootstrap EndpointHealthStatus = "Bootstrap"

	// EndpointHealthStatusPending captures enum value "Pending"
	EndpointHealthStatusPending EndpointHealthStatus = "Pending"

	// EndpointHealthStatusWarning captures enum value "Warning"
	EndpointHealthStatusWarning EndpointHealthStatus = "Warning"

	// EndpointHealthStatusFailure captures enum value "Failure"
	EndpointHealthStatusFailure EndpointHealthStatus = "Failure"

	// EndpointHealthStatusDisabled captures enum value "Disabled"
	EndpointHealthStatusDisabled EndpointHealthStatus = "Disabled"
)

// for schema
var endpointHealthStatusEnum []interface{}

func init() {
	var res []EndpointHealthStatus
	if err := json.Unmarshal([]byte(`["OK","Bootstrap","Pending","Warning","Failure","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointHealthStatusEnum = append(endpointHealthStatusEnum, v)
	}
}

func (m EndpointHealthStatus) validateEndpointHealthStatusEnum(path, location string, value EndpointHealthStatus) error {
	if err := validate.EnumCase(path, location, value, endpointHealthStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this endpoint health status
func (m EndpointHealthStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEndpointHealthStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
