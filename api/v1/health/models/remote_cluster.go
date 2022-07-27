// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RemoteCluster Status of remote cluster
// swagger:model remoteCluster
type RemoteCluster struct {

	// Name of the cluster
	Name string `json:"name,omitempty"`

	// Number of identities in the cluster
	NumIdentities int64 `json:"num-identities,omitempty"`

	// Number of nodes in the cluster
	NumNodes int64 `json:"num-nodes,omitempty"`

	// Number of services in the cluster
	NumSharedServices int64 `json:"num-shared-services,omitempty"`

	// Indicates readiness of the remote cluser
	Ready bool `json:"ready,omitempty"`

	// Status of the control plane
	Status string `json:"status,omitempty"`
}

// Validate validates this remote cluster
func (m *RemoteCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoteCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteCluster) UnmarshalBinary(b []byte) error {
	var res RemoteCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
