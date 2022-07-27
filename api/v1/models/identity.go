// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Identity Security identity
// swagger:model Identity

type Identity struct {

	// Unique identifier
	ID int64 `json:"id,omitempty"`

	// Labels describing the identity
	Labels Labels `json:"labels"`

	// SHA256 of labels
	LabelsSHA256 string `json:"labelsSHA256,omitempty"`
}

/* polymorph Identity id false */

/* polymorph Identity labels false */

/* polymorph Identity labelsSHA256 false */

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
