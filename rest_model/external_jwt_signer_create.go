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

// ExternalJWTSignerCreate A create Certificate Authority (CA) object
//
// swagger:model externalJwtSignerCreate
type ExternalJWTSignerCreate struct {

	// audience
	// Required: true
	Audience *string `json:"audience"`

	// cert pem
	CertPem *string `json:"certPem,omitempty"`

	// claims property
	ClaimsProperty *string `json:"claimsProperty,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// external auth Url
	ExternalAuthURL *string `json:"externalAuthUrl,omitempty"`

	// issuer
	// Required: true
	Issuer *string `json:"issuer"`

	// jwks endpoint
	// Format: uri
	JwksEndpoint *strfmt.URI `json:"jwksEndpoint,omitempty"`

	// kid
	Kid *string `json:"kid,omitempty"`

	// name
	// Example: MyApps Signer
	// Required: true
	Name *string `json:"name"`

	// tags
	Tags *Tags `json:"tags,omitempty"`

	// use external Id
	UseExternalID *bool `json:"useExternalId,omitempty"`
}

// Validate validates this external Jwt signer create
func (m *ExternalJWTSignerCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJwksEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalJWTSignerCreate) validateAudience(formats strfmt.Registry) error {

	if err := validate.Required("audience", "body", m.Audience); err != nil {
		return err
	}

	return nil
}

func (m *ExternalJWTSignerCreate) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ExternalJWTSignerCreate) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

func (m *ExternalJWTSignerCreate) validateJwksEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.JwksEndpoint) { // not required
		return nil
	}

	if err := validate.FormatOf("jwksEndpoint", "body", "uri", m.JwksEndpoint.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalJWTSignerCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ExternalJWTSignerCreate) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this external Jwt signer create based on the context it is used
func (m *ExternalJWTSignerCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalJWTSignerCreate) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if m.Tags != nil {

		if swag.IsZero(m.Tags) { // not required
			return nil
		}

		if err := m.Tags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalJWTSignerCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalJWTSignerCreate) UnmarshalBinary(b []byte) error {
	var res ExternalJWTSignerCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
