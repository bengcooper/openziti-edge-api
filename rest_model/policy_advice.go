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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyAdvice policy advice
//
// swagger:model policyAdvice
type PolicyAdvice struct {

	// common routers
	CommonRouters []*RouterEntityRef `json:"commonRouters"`

	// identity
	Identity *EntityRef `json:"identity,omitempty"`

	// identity Id
	IdentityID string `json:"identityId,omitempty"`

	// identity router count
	IdentityRouterCount int32 `json:"identityRouterCount,omitempty"`

	// is bind allowed
	IsBindAllowed bool `json:"isBindAllowed,omitempty"`

	// is dial allowed
	IsDialAllowed bool `json:"isDialAllowed,omitempty"`

	// service
	Service *EntityRef `json:"service,omitempty"`

	// service Id
	ServiceID string `json:"serviceId,omitempty"`

	// service router count
	ServiceRouterCount int32 `json:"serviceRouterCount,omitempty"`
}

// Validate validates this policy advice
func (m *PolicyAdvice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommonRouters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyAdvice) validateCommonRouters(formats strfmt.Registry) error {
	if swag.IsZero(m.CommonRouters) { // not required
		return nil
	}

	for i := 0; i < len(m.CommonRouters); i++ {
		if swag.IsZero(m.CommonRouters[i]) { // not required
			continue
		}

		if m.CommonRouters[i] != nil {
			if err := m.CommonRouters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commonRouters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commonRouters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyAdvice) validateIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyAdvice) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy advice based on the context it is used
func (m *PolicyAdvice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommonRouters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyAdvice) contextValidateCommonRouters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CommonRouters); i++ {

		if m.CommonRouters[i] != nil {

			if swag.IsZero(m.CommonRouters[i]) { // not required
				return nil
			}

			if err := m.CommonRouters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commonRouters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commonRouters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyAdvice) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.Identity != nil {

		if swag.IsZero(m.Identity) { // not required
			return nil
		}

		if err := m.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyAdvice) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if swag.IsZero(m.Service) { // not required
			return nil
		}

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyAdvice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyAdvice) UnmarshalBinary(b []byte) error {
	var res PolicyAdvice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
