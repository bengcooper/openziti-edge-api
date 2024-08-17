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

// CaDetail A Certificate Authority (CA) resource
//
// swagger:model caDetail
type CaDetail struct {
	BaseEntity

	// cert pem
	// Required: true
	CertPem *string `json:"certPem"`

	// external Id claim
	ExternalIDClaim *ExternalIDClaim `json:"externalIdClaim,omitempty"`

	// fingerprint
	// Required: true
	Fingerprint *string `json:"fingerprint"`

	// identity name format
	// Required: true
	IdentityNameFormat *string `json:"identityNameFormat"`

	// identity roles
	// Required: true
	IdentityRoles Roles `json:"identityRoles"`

	// is auth enabled
	// Example: true
	// Required: true
	IsAuthEnabled *bool `json:"isAuthEnabled"`

	// is auto ca enrollment enabled
	// Example: true
	// Required: true
	IsAutoCaEnrollmentEnabled *bool `json:"isAutoCaEnrollmentEnabled"`

	// is ott ca enrollment enabled
	// Example: true
	// Required: true
	IsOttCaEnrollmentEnabled *bool `json:"isOttCaEnrollmentEnabled"`

	// is verified
	// Example: false
	// Required: true
	IsVerified *bool `json:"isVerified"`

	// name
	// Required: true
	Name *string `json:"name"`

	// verification token
	// Format: uuid
	VerificationToken strfmt.UUID `json:"verificationToken,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CaDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		CertPem *string `json:"certPem"`

		ExternalIDClaim *ExternalIDClaim `json:"externalIdClaim,omitempty"`

		Fingerprint *string `json:"fingerprint"`

		IdentityNameFormat *string `json:"identityNameFormat"`

		IdentityRoles Roles `json:"identityRoles"`

		IsAuthEnabled *bool `json:"isAuthEnabled"`

		IsAutoCaEnrollmentEnabled *bool `json:"isAutoCaEnrollmentEnabled"`

		IsOttCaEnrollmentEnabled *bool `json:"isOttCaEnrollmentEnabled"`

		IsVerified *bool `json:"isVerified"`

		Name *string `json:"name"`

		VerificationToken strfmt.UUID `json:"verificationToken,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CertPem = dataAO1.CertPem

	m.ExternalIDClaim = dataAO1.ExternalIDClaim

	m.Fingerprint = dataAO1.Fingerprint

	m.IdentityNameFormat = dataAO1.IdentityNameFormat

	m.IdentityRoles = dataAO1.IdentityRoles

	m.IsAuthEnabled = dataAO1.IsAuthEnabled

	m.IsAutoCaEnrollmentEnabled = dataAO1.IsAutoCaEnrollmentEnabled

	m.IsOttCaEnrollmentEnabled = dataAO1.IsOttCaEnrollmentEnabled

	m.IsVerified = dataAO1.IsVerified

	m.Name = dataAO1.Name

	m.VerificationToken = dataAO1.VerificationToken

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CaDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CertPem *string `json:"certPem"`

		ExternalIDClaim *ExternalIDClaim `json:"externalIdClaim,omitempty"`

		Fingerprint *string `json:"fingerprint"`

		IdentityNameFormat *string `json:"identityNameFormat"`

		IdentityRoles Roles `json:"identityRoles"`

		IsAuthEnabled *bool `json:"isAuthEnabled"`

		IsAutoCaEnrollmentEnabled *bool `json:"isAutoCaEnrollmentEnabled"`

		IsOttCaEnrollmentEnabled *bool `json:"isOttCaEnrollmentEnabled"`

		IsVerified *bool `json:"isVerified"`

		Name *string `json:"name"`

		VerificationToken strfmt.UUID `json:"verificationToken,omitempty"`
	}

	dataAO1.CertPem = m.CertPem

	dataAO1.ExternalIDClaim = m.ExternalIDClaim

	dataAO1.Fingerprint = m.Fingerprint

	dataAO1.IdentityNameFormat = m.IdentityNameFormat

	dataAO1.IdentityRoles = m.IdentityRoles

	dataAO1.IsAuthEnabled = m.IsAuthEnabled

	dataAO1.IsAutoCaEnrollmentEnabled = m.IsAutoCaEnrollmentEnabled

	dataAO1.IsOttCaEnrollmentEnabled = m.IsOttCaEnrollmentEnabled

	dataAO1.IsVerified = m.IsVerified

	dataAO1.Name = m.Name

	dataAO1.VerificationToken = m.VerificationToken

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ca detail
func (m *CaDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertPem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalIDClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityNameFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAuthEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAutoCaEnrollmentEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsOttCaEnrollmentEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsVerified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerificationToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaDetail) validateCertPem(formats strfmt.Registry) error {

	if err := validate.Required("certPem", "body", m.CertPem); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateExternalIDClaim(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalIDClaim) { // not required
		return nil
	}

	if m.ExternalIDClaim != nil {
		if err := m.ExternalIDClaim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalIdClaim")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalIdClaim")
			}
			return err
		}
	}

	return nil
}

func (m *CaDetail) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateIdentityNameFormat(formats strfmt.Registry) error {

	if err := validate.Required("identityNameFormat", "body", m.IdentityNameFormat); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateIdentityRoles(formats strfmt.Registry) error {

	if err := validate.Required("identityRoles", "body", m.IdentityRoles); err != nil {
		return err
	}

	if err := m.IdentityRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identityRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identityRoles")
		}
		return err
	}

	return nil
}

func (m *CaDetail) validateIsAuthEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isAuthEnabled", "body", m.IsAuthEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateIsAutoCaEnrollmentEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isAutoCaEnrollmentEnabled", "body", m.IsAutoCaEnrollmentEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateIsOttCaEnrollmentEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isOttCaEnrollmentEnabled", "body", m.IsOttCaEnrollmentEnabled); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateIsVerified(formats strfmt.Registry) error {

	if err := validate.Required("isVerified", "body", m.IsVerified); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CaDetail) validateVerificationToken(formats strfmt.Registry) error {

	if swag.IsZero(m.VerificationToken) { // not required
		return nil
	}

	if err := validate.FormatOf("verificationToken", "body", "uuid", m.VerificationToken.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ca detail based on the context it is used
func (m *CaDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalIDClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaDetail) contextValidateExternalIDClaim(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalIDClaim != nil {

		if swag.IsZero(m.ExternalIDClaim) { // not required
			return nil
		}

		if err := m.ExternalIDClaim.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalIdClaim")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalIdClaim")
			}
			return err
		}
	}

	return nil
}

func (m *CaDetail) contextValidateIdentityRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IdentityRoles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identityRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identityRoles")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CaDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CaDetail) UnmarshalBinary(b []byte) error {
	var res CaDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
