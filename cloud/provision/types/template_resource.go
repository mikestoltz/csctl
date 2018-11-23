// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TemplateResource template resource
// swagger:model TemplateResource
type TemplateResource struct {

	// digitalocean droplet
	DigitaloceanDroplet DigitalOceanDropletMap `json:"digitalocean_droplet,omitempty"`

	// packet device
	PacketDevice PacketDeviceMap `json:"packet_device,omitempty"`
}

// Validate validates this template resource
func (m *TemplateResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDigitaloceanDroplet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacketDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateResource) validateDigitaloceanDroplet(formats strfmt.Registry) error {

	if swag.IsZero(m.DigitaloceanDroplet) { // not required
		return nil
	}

	if err := m.DigitaloceanDroplet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("digitalocean_droplet")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validatePacketDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.PacketDevice) { // not required
		return nil
	}

	if err := m.PacketDevice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packet_device")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateResource) UnmarshalBinary(b []byte) error {
	var res TemplateResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}