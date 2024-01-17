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

package external_jwt_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// CreateExternalJWTSignerReader is a Reader for the CreateExternalJWTSigner structure.
type CreateExternalJWTSignerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExternalJWTSignerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateExternalJWTSignerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateExternalJWTSignerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateExternalJWTSignerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateExternalJWTSignerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateExternalJWTSignerCreated creates a CreateExternalJWTSignerCreated with default headers values
func NewCreateExternalJWTSignerCreated() *CreateExternalJWTSignerCreated {
	return &CreateExternalJWTSignerCreated{}
}

/* CreateExternalJWTSignerCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateExternalJWTSignerCreated struct {
	Payload *rest_model.CreateEnvelope
}

func (o *CreateExternalJWTSignerCreated) Error() string {
	return fmt.Sprintf("[POST /external-jwt-signers][%d] createExternalJwtSignerCreated  %+v", 201, o.Payload)
}
func (o *CreateExternalJWTSignerCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateExternalJWTSignerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExternalJWTSignerBadRequest creates a CreateExternalJWTSignerBadRequest with default headers values
func NewCreateExternalJWTSignerBadRequest() *CreateExternalJWTSignerBadRequest {
	return &CreateExternalJWTSignerBadRequest{}
}

/* CreateExternalJWTSignerBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateExternalJWTSignerBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateExternalJWTSignerBadRequest) Error() string {
	return fmt.Sprintf("[POST /external-jwt-signers][%d] createExternalJwtSignerBadRequest  %+v", 400, o.Payload)
}
func (o *CreateExternalJWTSignerBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateExternalJWTSignerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExternalJWTSignerUnauthorized creates a CreateExternalJWTSignerUnauthorized with default headers values
func NewCreateExternalJWTSignerUnauthorized() *CreateExternalJWTSignerUnauthorized {
	return &CreateExternalJWTSignerUnauthorized{}
}

/* CreateExternalJWTSignerUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type CreateExternalJWTSignerUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateExternalJWTSignerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /external-jwt-signers][%d] createExternalJwtSignerUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateExternalJWTSignerUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateExternalJWTSignerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExternalJWTSignerTooManyRequests creates a CreateExternalJWTSignerTooManyRequests with default headers values
func NewCreateExternalJWTSignerTooManyRequests() *CreateExternalJWTSignerTooManyRequests {
	return &CreateExternalJWTSignerTooManyRequests{}
}

/* CreateExternalJWTSignerTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type CreateExternalJWTSignerTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateExternalJWTSignerTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /external-jwt-signers][%d] createExternalJwtSignerTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateExternalJWTSignerTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateExternalJWTSignerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
