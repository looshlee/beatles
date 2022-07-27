// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EndpointStatus The current state and configuration of the endpoint, its policy & datapath, and subcomponents
// swagger:model EndpointStatus

type EndpointStatus struct {

	// Status of internal controllers attached to this endpoint
	Controllers ControllerStatuses `json:"controllers"`

	// Unique identifiers for this endpoint from outside cilium
	ExternalIdentifiers *EndpointIdentifiers `json:"external-identifiers,omitempty"`

	// Summary overall endpoint & subcomponent health
	Health *EndpointHealth `json:"health,omitempty"`

	// The security identity for this endpoint
	Identity *Identity `json:"identity,omitempty"`

	// Labels applied to this endpoint
	Labels *LabelConfigurationStatus `json:"labels,omitempty"`

	// Most recent status log. See endpoint/{id}/log for the complete log.
	Log EndpointStatusLog `json:"log"`

	// Networking properties of the endpoint
	Networking *EndpointNetworking `json:"networking,omitempty"`

	// The policy applied to this endpoint from the policy repository
	Policy *EndpointPolicyStatus `json:"policy,omitempty"`

	// The configuration in effect on this endpoint
	Realized *EndpointConfigurationSpec `json:"realized,omitempty"`

	// Current state of endpoint
	// Required: true
	State EndpointState `json:"state"`
}

/* polymorph EndpointStatus controllers false */

/* polymorph EndpointStatus external-identifiers false */

/* polymorph EndpointStatus health false */

/* polymorph EndpointStatus identity false */

/* polymorph EndpointStatus labels false */

/* polymorph EndpointStatus log false */

/* polymorph EndpointStatus networking false */

/* polymorph EndpointStatus policy false */

/* polymorph EndpointStatus realized false */

/* polymorph EndpointStatus state false */

// Validate validates this endpoint status
func (m *EndpointStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalIdentifiers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworking(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRealized(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointStatus) validateExternalIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalIdentifiers) { // not required
		return nil
	}

	if m.ExternalIdentifiers != nil {

		if err := m.ExternalIdentifiers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external-identifiers")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {

		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {

		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {

		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateNetworking(formats strfmt.Registry) error {

	if swag.IsZero(m.Networking) { // not required
		return nil
	}

	if m.Networking != nil {

		if err := m.Networking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networking")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {

		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateRealized(formats strfmt.Registry) error {

	if swag.IsZero(m.Realized) { // not required
		return nil
	}

	if m.Realized != nil {

		if err := m.Realized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointStatus) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointStatus) UnmarshalBinary(b []byte) error {
	var res EndpointStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
