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

// EdgeRouterPolicyDetail edge router policy detail
//
// swagger:model edgeRouterPolicyDetail
type EdgeRouterPolicyDetail struct {
	BaseEntity

	// edge router roles
	// Required: true
	EdgeRouterRoles Roles `json:"edgeRouterRoles"`

	// edge router roles display
	// Required: true
	EdgeRouterRolesDisplay NamedRoles `json:"edgeRouterRolesDisplay"`

	// identity roles
	// Required: true
	IdentityRoles Roles `json:"identityRoles"`

	// identity roles display
	// Required: true
	IdentityRolesDisplay NamedRoles `json:"identityRolesDisplay"`

	// is system
	// Required: true
	IsSystem *bool `json:"isSystem"`

	// name
	// Required: true
	Name *string `json:"name"`

	// semantic
	// Required: true
	Semantic *Semantic `json:"semantic"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EdgeRouterPolicyDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		EdgeRouterRoles Roles `json:"edgeRouterRoles"`

		EdgeRouterRolesDisplay NamedRoles `json:"edgeRouterRolesDisplay"`

		IdentityRoles Roles `json:"identityRoles"`

		IdentityRolesDisplay NamedRoles `json:"identityRolesDisplay"`

		IsSystem *bool `json:"isSystem"`

		Name *string `json:"name"`

		Semantic *Semantic `json:"semantic"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EdgeRouterRoles = dataAO1.EdgeRouterRoles

	m.EdgeRouterRolesDisplay = dataAO1.EdgeRouterRolesDisplay

	m.IdentityRoles = dataAO1.IdentityRoles

	m.IdentityRolesDisplay = dataAO1.IdentityRolesDisplay

	m.IsSystem = dataAO1.IsSystem

	m.Name = dataAO1.Name

	m.Semantic = dataAO1.Semantic

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EdgeRouterPolicyDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		EdgeRouterRoles Roles `json:"edgeRouterRoles"`

		EdgeRouterRolesDisplay NamedRoles `json:"edgeRouterRolesDisplay"`

		IdentityRoles Roles `json:"identityRoles"`

		IdentityRolesDisplay NamedRoles `json:"identityRolesDisplay"`

		IsSystem *bool `json:"isSystem"`

		Name *string `json:"name"`

		Semantic *Semantic `json:"semantic"`
	}

	dataAO1.EdgeRouterRoles = m.EdgeRouterRoles

	dataAO1.EdgeRouterRolesDisplay = m.EdgeRouterRolesDisplay

	dataAO1.IdentityRoles = m.IdentityRoles

	dataAO1.IdentityRolesDisplay = m.IdentityRolesDisplay

	dataAO1.IsSystem = m.IsSystem

	dataAO1.Name = m.Name

	dataAO1.Semantic = m.Semantic

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this edge router policy detail
func (m *EdgeRouterPolicyDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeRouterRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeRouterRolesDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRolesDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemantic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeRouterPolicyDetail) validateEdgeRouterRoles(formats strfmt.Registry) error {

	if err := validate.Required("edgeRouterRoles", "body", m.EdgeRouterRoles); err != nil {
		return err
	}

	if err := m.EdgeRouterRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRoles")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) validateEdgeRouterRolesDisplay(formats strfmt.Registry) error {

	if err := validate.Required("edgeRouterRolesDisplay", "body", m.EdgeRouterRolesDisplay); err != nil {
		return err
	}

	if err := m.EdgeRouterRolesDisplay.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRolesDisplay")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) validateIdentityRoles(formats strfmt.Registry) error {

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

func (m *EdgeRouterPolicyDetail) validateIdentityRolesDisplay(formats strfmt.Registry) error {

	if err := validate.Required("identityRolesDisplay", "body", m.IdentityRolesDisplay); err != nil {
		return err
	}

	if err := m.IdentityRolesDisplay.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identityRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identityRolesDisplay")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) validateIsSystem(formats strfmt.Registry) error {

	if err := validate.Required("isSystem", "body", m.IsSystem); err != nil {
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) validateSemantic(formats strfmt.Registry) error {

	if err := validate.Required("semantic", "body", m.Semantic); err != nil {
		return err
	}

	if err := validate.Required("semantic", "body", m.Semantic); err != nil {
		return err
	}

	if m.Semantic != nil {
		if err := m.Semantic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semantic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semantic")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edge router policy detail based on the context it is used
func (m *EdgeRouterPolicyDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeRouterRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeRouterRolesDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityRolesDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSemantic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeRouterPolicyDetail) contextValidateEdgeRouterRoles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EdgeRouterRoles.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRoles")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRoles")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) contextValidateEdgeRouterRolesDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EdgeRouterRolesDisplay.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("edgeRouterRolesDisplay")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) contextValidateIdentityRoles(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EdgeRouterPolicyDetail) contextValidateIdentityRolesDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IdentityRolesDisplay.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("identityRolesDisplay")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("identityRolesDisplay")
		}
		return err
	}

	return nil
}

func (m *EdgeRouterPolicyDetail) contextValidateSemantic(ctx context.Context, formats strfmt.Registry) error {

	if m.Semantic != nil {

		if err := m.Semantic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semantic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semantic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeRouterPolicyDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeRouterPolicyDetail) UnmarshalBinary(b []byte) error {
	var res EdgeRouterPolicyDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
