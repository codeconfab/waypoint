// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointListHostnamesResponse hashicorp waypoint list hostnames response
//
// swagger:model hashicorp.waypoint.ListHostnamesResponse
type HashicorpWaypointListHostnamesResponse struct {

	// hostnames
	Hostnames []*HashicorpWaypointHostname `json:"hostnames"`
}

// Validate validates this hashicorp waypoint list hostnames response
func (m *HashicorpWaypointListHostnamesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostnames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListHostnamesResponse) validateHostnames(formats strfmt.Registry) error {
	if swag.IsZero(m.Hostnames) { // not required
		return nil
	}

	for i := 0; i < len(m.Hostnames); i++ {
		if swag.IsZero(m.Hostnames[i]) { // not required
			continue
		}

		if m.Hostnames[i] != nil {
			if err := m.Hostnames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostnames" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostnames" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp waypoint list hostnames response based on the context it is used
func (m *HashicorpWaypointListHostnamesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostnames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointListHostnamesResponse) contextValidateHostnames(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hostnames); i++ {

		if m.Hostnames[i] != nil {
			if err := m.Hostnames[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostnames" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostnames" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointListHostnamesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointListHostnamesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointListHostnamesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}