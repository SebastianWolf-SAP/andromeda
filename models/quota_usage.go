// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaUsage quota usage
//
// swagger:model quota_usage
type QuotaUsage struct {

	// The current quota usage of datacenter.
	// Example: 5
	InUseDatacenter int64 `json:"in_use_datacenter" db:"in_use_datacenter"`

	// The current quota usage of domain (provider = Akamai).
	// Example: 5
	InUseDomainAkamai int64 `json:"in_use_domain_akamai" db:"in_use_domain_akamai"`

	// The current quota usage of domain (provider = F5).
	// Example: 5
	InUseDomainF5 int64 `json:"in_use_domain_f5" db:"in_use_domain_f5"`

	// The current quota usage of member.
	// Example: 5
	InUseMember int64 `json:"in_use_member" db:"in_use_member"`

	// The current quota usage of monitor.
	// Example: 5
	InUseMonitor int64 `json:"in_use_monitor" db:"in_use_monitor"`

	// The current quota usage of pool.
	// Example: 5
	InUsePool int64 `json:"in_use_pool" db:"in_use_pool"`
}

// Validate validates this quota usage
func (m *QuotaUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota usage based on context it is used
func (m *QuotaUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaUsage) UnmarshalBinary(b []byte) error {
	var res QuotaUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
