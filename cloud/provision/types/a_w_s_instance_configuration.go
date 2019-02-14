// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AWSInstanceConfiguration AWS instance configuration
// swagger:model AWSInstanceConfiguration
type AWSInstanceConfiguration struct {

	// AMI for this instance
	Ami string `json:"ami,omitempty"`

	// Flag indicating whether to associate a public IP with this instance
	AssociatePublicIPAddress bool `json:"associate_public_ip_address,omitempty"`

	// Availability zone for this instance
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// Dependencies for this instance
	DependsOn []string `json:"depends_on"`

	// AWS instance type
	InstanceType string `json:"instance_type,omitempty"`

	// Subnet this instance belongs to
	SubnetID string `json:"subnet_id,omitempty"`

	// VPC security group IDs
	VpcSecurityGroupIds []string `json:"vpc_security_group_ids"`
}

// Validate validates this a w s instance configuration
func (m *AWSInstanceConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AWSInstanceConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSInstanceConfiguration) UnmarshalBinary(b []byte) error {
	var res AWSInstanceConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
