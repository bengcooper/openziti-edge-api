// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthPolicyDetail A Auth Policy resource
//
// swagger:model authPolicyDetail
type AuthPolicyDetail struct {
	BaseEntity

	// name
	// Required: true
	Name *string `json:"name"`

	// primary
	// Required: true
	Primary *AuthPolicyPrimary `json:"primary"`

	// secondary
	// Required: true
	Secondary *AuthPolicySecondary `json:"secondary"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AuthPolicyDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		Name *string `json:"name"`

		Primary *AuthPolicyPrimary `json:"primary"`

		Secondary *AuthPolicySecondary `json:"secondary"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	m.Primary = dataAO1.Primary

	m.Secondary = dataAO1.Secondary

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AuthPolicyDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Name *string `json:"name"`

		Primary *AuthPolicyPrimary `json:"primary"`

		Secondary *AuthPolicySecondary `json:"secondary"`
	}

	dataAO1.Name = m.Name

	dataAO1.Primary = m.Primary

	dataAO1.Secondary = m.Secondary

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this auth policy detail
func (m *AuthPolicyDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthPolicyDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AuthPolicyDetail) validatePrimary(formats strfmt.Registry) error {

	if err := validate.Required("primary", "body", m.Primary); err != nil {
		return err
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

func (m *AuthPolicyDetail) validateSecondary(formats strfmt.Registry) error {

	if err := validate.Required("secondary", "body", m.Secondary); err != nil {
		return err
	}

	if m.Secondary != nil {
		if err := m.Secondary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auth policy detail based on the context it is used
func (m *AuthPolicyDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthPolicyDetail) contextValidatePrimary(ctx context.Context, formats strfmt.Registry) error {

	if m.Primary != nil {

		if err := m.Primary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

func (m *AuthPolicyDetail) contextValidateSecondary(ctx context.Context, formats strfmt.Registry) error {

	if m.Secondary != nil {

		if err := m.Secondary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthPolicyDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthPolicyDetail) UnmarshalBinary(b []byte) error {
	var res AuthPolicyDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
