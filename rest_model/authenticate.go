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
)

// Authenticate A generic authenticate object meant for use with the /authenticate path. Required fields depend on authentication method.
//
// swagger:model authenticate
type Authenticate struct {

	// config types
	ConfigTypes ConfigTypes `json:"configTypes,omitempty"`

	// env info
	EnvInfo *EnvInfo `json:"envInfo,omitempty"`

	// password
	Password Password `json:"password,omitempty"`

	// sdk info
	SdkInfo *SdkInfo `json:"sdkInfo,omitempty"`

	// username
	Username Username `json:"username,omitempty"`
}

// Validate validates this authenticate
func (m *Authenticate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSdkInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authenticate) validateConfigTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigTypes) { // not required
		return nil
	}

	if err := m.ConfigTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("configTypes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("configTypes")
		}
		return err
	}

	return nil
}

func (m *Authenticate) validateEnvInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvInfo) { // not required
		return nil
	}

	if m.EnvInfo != nil {
		if err := m.EnvInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("envInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("envInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Authenticate) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := m.Password.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("password")
		}
		return err
	}

	return nil
}

func (m *Authenticate) validateSdkInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SdkInfo) { // not required
		return nil
	}

	if m.SdkInfo != nil {
		if err := m.SdkInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sdkInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sdkInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Authenticate) validateUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := m.Username.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("username")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("username")
		}
		return err
	}

	return nil
}

// ContextValidate validate this authenticate based on the context it is used
func (m *Authenticate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePassword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSdkInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsername(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authenticate) contextValidateConfigTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConfigTypes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("configTypes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("configTypes")
		}
		return err
	}

	return nil
}

func (m *Authenticate) contextValidateEnvInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.EnvInfo != nil {

		if swag.IsZero(m.EnvInfo) { // not required
			return nil
		}

		if err := m.EnvInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("envInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("envInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Authenticate) contextValidatePassword(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := m.Password.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("password")
		}
		return err
	}

	return nil
}

func (m *Authenticate) contextValidateSdkInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SdkInfo != nil {

		if swag.IsZero(m.SdkInfo) { // not required
			return nil
		}

		if err := m.SdkInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sdkInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sdkInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Authenticate) contextValidateUsername(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := m.Username.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("username")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("username")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Authenticate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Authenticate) UnmarshalBinary(b []byte) error {
	var res Authenticate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
