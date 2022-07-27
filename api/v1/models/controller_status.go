// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ControllerStatus Status of a controller
//
// +k8s:deepcopy-gen=true
//
// swagger:model ControllerStatus
type ControllerStatus struct {

	// configuration
	Configuration *ControllerStatusConfiguration `json:"configuration,omitempty"`

	// Name of controller
	Name string `json:"name,omitempty"`

	// status
	Status *ControllerStatusStatus `json:"status,omitempty"`

	// UUID of controller
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this controller status
func (m *ControllerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerStatus) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ControllerStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ControllerStatus) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerStatus) UnmarshalBinary(b []byte) error {
	var res ControllerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ControllerStatusConfiguration Configuration of controller
//
// +deepequal-gen=true
// +k8s:deepcopy-gen=true
//
// swagger:model ControllerStatusConfiguration
type ControllerStatusConfiguration struct {

	// Retry on error
	ErrorRetry bool `json:"error-retry,omitempty"`

	// Base error retry back-off time
	// Format: duration
	ErrorRetryBase strfmt.Duration `json:"error-retry-base,omitempty"`

	// Regular synchronization interval
	// Format: duration
	Interval strfmt.Duration `json:"interval,omitempty"`
}

// Validate validates this controller status configuration
func (m *ControllerStatusConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorRetryBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerStatusConfiguration) validateErrorRetryBase(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorRetryBase) { // not required
		return nil
	}

	if err := validate.FormatOf("configuration"+"."+"error-retry-base", "body", "duration", m.ErrorRetryBase.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ControllerStatusConfiguration) validateInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	if err := validate.FormatOf("configuration"+"."+"interval", "body", "duration", m.Interval.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerStatusConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerStatusConfiguration) UnmarshalBinary(b []byte) error {
	var res ControllerStatusConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ControllerStatusStatus Current status of controller
//
// +k8s:deepcopy-gen=true
//
// swagger:model ControllerStatusStatus
type ControllerStatusStatus struct {

	// Number of consecutive errors since last success
	ConsecutiveFailureCount int64 `json:"consecutive-failure-count,omitempty"`

	// Total number of failed runs
	FailureCount int64 `json:"failure-count,omitempty"`

	// Error message of last failed run
	LastFailureMsg string `json:"last-failure-msg,omitempty"`

	// Timestamp of last error
	// Format: date-time
	LastFailureTimestamp strfmt.DateTime `json:"last-failure-timestamp,omitempty"`

	// Timestamp of last success
	// Format: date-time
	LastSuccessTimestamp strfmt.DateTime `json:"last-success-timestamp,omitempty"`

	// Total number of successful runs
	SuccessCount int64 `json:"success-count,omitempty"`
}

// Validate validates this controller status status
func (m *ControllerStatusStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastFailureTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSuccessTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerStatusStatus) validateLastFailureTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.LastFailureTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("status"+"."+"last-failure-timestamp", "body", "date-time", m.LastFailureTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ControllerStatusStatus) validateLastSuccessTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.LastSuccessTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("status"+"."+"last-success-timestamp", "body", "date-time", m.LastSuccessTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerStatusStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerStatusStatus) UnmarshalBinary(b []byte) error {
	var res ControllerStatusStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
