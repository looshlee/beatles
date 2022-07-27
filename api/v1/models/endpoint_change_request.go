// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EndpointChangeRequest Structure which contains the mutable elements of an Endpoint.
//
// swagger:model EndpointChangeRequest

type EndpointChangeRequest struct {

	// addressing
	Addressing *AddressPair `json:"addressing,omitempty"`

	// ID assigned by container runtime
	ContainerID string `json:"container-id,omitempty"`

	// Name assigned to container
	ContainerName string `json:"container-name,omitempty"`

	// ID of datapath tail call map
	DatapathMapID int64 `json:"datapath-map-id,omitempty"`

	// Docker endpoint ID
	DockerEndpointID string `json:"docker-endpoint-id,omitempty"`

	// Docker network ID
	DockerNetworkID string `json:"docker-network-id,omitempty"`

	// MAC address
	HostMac string `json:"host-mac,omitempty"`

	// Local endpoint ID
	ID int64 `json:"id,omitempty"`

	// Index of network device
	InterfaceIndex int64 `json:"interface-index,omitempty"`

	// Name of network device
	InterfaceName string `json:"interface-name,omitempty"`

	// Kubernetes namespace name
	K8sNamespace string `json:"k8s-namespace,omitempty"`

	// Kubernetes pod name
	K8sPodName string `json:"k8s-pod-name,omitempty"`

	// Labels describing the identity
	Labels Labels `json:"labels"`

	// MAC address
	Mac string `json:"mac,omitempty"`

	// Process ID of the workload belonging to this endpoint
	Pid int64 `json:"pid,omitempty"`

	// Whether policy enforcement is enabled or not
	PolicyEnabled bool `json:"policy-enabled,omitempty"`

	// Current state of endpoint
	// Required: true
	State EndpointState `json:"state"`

	// Whether to build an endpoint synchronously
	//
	SyncBuildEndpoint bool `json:"sync-build-endpoint,omitempty"`
}

/* polymorph EndpointChangeRequest addressing false */

/* polymorph EndpointChangeRequest container-id false */

/* polymorph EndpointChangeRequest container-name false */

/* polymorph EndpointChangeRequest datapath-map-id false */

/* polymorph EndpointChangeRequest docker-endpoint-id false */

/* polymorph EndpointChangeRequest docker-network-id false */

/* polymorph EndpointChangeRequest host-mac false */

/* polymorph EndpointChangeRequest id false */

/* polymorph EndpointChangeRequest interface-index false */

/* polymorph EndpointChangeRequest interface-name false */

/* polymorph EndpointChangeRequest k8s-namespace false */

/* polymorph EndpointChangeRequest k8s-pod-name false */

/* polymorph EndpointChangeRequest labels false */

/* polymorph EndpointChangeRequest mac false */

/* polymorph EndpointChangeRequest pid false */

/* polymorph EndpointChangeRequest policy-enabled false */

/* polymorph EndpointChangeRequest state false */

/* polymorph EndpointChangeRequest sync-build-endpoint false */

// Validate validates this endpoint change request
func (m *EndpointChangeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressing(formats); err != nil {
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

func (m *EndpointChangeRequest) validateAddressing(formats strfmt.Registry) error {

	if swag.IsZero(m.Addressing) { // not required
		return nil
	}

	if m.Addressing != nil {

		if err := m.Addressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressing")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointChangeRequest) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointChangeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointChangeRequest) UnmarshalBinary(b []byte) error {
	var res EndpointChangeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
