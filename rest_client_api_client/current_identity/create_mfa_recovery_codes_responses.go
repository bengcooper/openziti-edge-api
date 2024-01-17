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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// CreateMfaRecoveryCodesReader is a Reader for the CreateMfaRecoveryCodes structure.
type CreateMfaRecoveryCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMfaRecoveryCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMfaRecoveryCodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMfaRecoveryCodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateMfaRecoveryCodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMfaRecoveryCodesOK creates a CreateMfaRecoveryCodesOK with default headers values
func NewCreateMfaRecoveryCodesOK() *CreateMfaRecoveryCodesOK {
	return &CreateMfaRecoveryCodesOK{}
}

/* CreateMfaRecoveryCodesOK describes a response with status code 200, with default header values.

The recovery codes of an MFA enrollment
*/
type CreateMfaRecoveryCodesOK struct {
	Payload *rest_model.DetailMfaRecoveryCodesEnvelope
}

func (o *CreateMfaRecoveryCodesOK) Error() string {
	return fmt.Sprintf("[POST /current-identity/mfa/recovery-codes][%d] createMfaRecoveryCodesOK  %+v", 200, o.Payload)
}
func (o *CreateMfaRecoveryCodesOK) GetPayload() *rest_model.DetailMfaRecoveryCodesEnvelope {
	return o.Payload
}

func (o *CreateMfaRecoveryCodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailMfaRecoveryCodesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMfaRecoveryCodesUnauthorized creates a CreateMfaRecoveryCodesUnauthorized with default headers values
func NewCreateMfaRecoveryCodesUnauthorized() *CreateMfaRecoveryCodesUnauthorized {
	return &CreateMfaRecoveryCodesUnauthorized{}
}

/* CreateMfaRecoveryCodesUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type CreateMfaRecoveryCodesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateMfaRecoveryCodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /current-identity/mfa/recovery-codes][%d] createMfaRecoveryCodesUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateMfaRecoveryCodesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateMfaRecoveryCodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMfaRecoveryCodesNotFound creates a CreateMfaRecoveryCodesNotFound with default headers values
func NewCreateMfaRecoveryCodesNotFound() *CreateMfaRecoveryCodesNotFound {
	return &CreateMfaRecoveryCodesNotFound{}
}

/* CreateMfaRecoveryCodesNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type CreateMfaRecoveryCodesNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateMfaRecoveryCodesNotFound) Error() string {
	return fmt.Sprintf("[POST /current-identity/mfa/recovery-codes][%d] createMfaRecoveryCodesNotFound  %+v", 404, o.Payload)
}
func (o *CreateMfaRecoveryCodesNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateMfaRecoveryCodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
