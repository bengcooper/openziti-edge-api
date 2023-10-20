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

// PatchExternalJWTSignerReader is a Reader for the PatchExternalJWTSigner structure.
type PatchExternalJWTSignerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchExternalJWTSignerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchExternalJWTSignerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchExternalJWTSignerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchExternalJWTSignerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchExternalJWTSignerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchExternalJWTSignerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchExternalJWTSignerOK creates a PatchExternalJWTSignerOK with default headers values
func NewPatchExternalJWTSignerOK() *PatchExternalJWTSignerOK {
	return &PatchExternalJWTSignerOK{}
}

/* PatchExternalJWTSignerOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchExternalJWTSignerOK struct {
	Payload *rest_model.Empty
}

func (o *PatchExternalJWTSignerOK) Error() string {
	return fmt.Sprintf("[PATCH /external-jwt-signers/{id}][%d] patchExternalJwtSignerOK  %+v", 200, o.Payload)
}
func (o *PatchExternalJWTSignerOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchExternalJWTSignerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchExternalJWTSignerBadRequest creates a PatchExternalJWTSignerBadRequest with default headers values
func NewPatchExternalJWTSignerBadRequest() *PatchExternalJWTSignerBadRequest {
	return &PatchExternalJWTSignerBadRequest{}
}

/* PatchExternalJWTSignerBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchExternalJWTSignerBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchExternalJWTSignerBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /external-jwt-signers/{id}][%d] patchExternalJwtSignerBadRequest  %+v", 400, o.Payload)
}
func (o *PatchExternalJWTSignerBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchExternalJWTSignerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchExternalJWTSignerUnauthorized creates a PatchExternalJWTSignerUnauthorized with default headers values
func NewPatchExternalJWTSignerUnauthorized() *PatchExternalJWTSignerUnauthorized {
	return &PatchExternalJWTSignerUnauthorized{}
}

/* PatchExternalJWTSignerUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type PatchExternalJWTSignerUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchExternalJWTSignerUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /external-jwt-signers/{id}][%d] patchExternalJwtSignerUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchExternalJWTSignerUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchExternalJWTSignerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchExternalJWTSignerNotFound creates a PatchExternalJWTSignerNotFound with default headers values
func NewPatchExternalJWTSignerNotFound() *PatchExternalJWTSignerNotFound {
	return &PatchExternalJWTSignerNotFound{}
}

/* PatchExternalJWTSignerNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchExternalJWTSignerNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchExternalJWTSignerNotFound) Error() string {
	return fmt.Sprintf("[PATCH /external-jwt-signers/{id}][%d] patchExternalJwtSignerNotFound  %+v", 404, o.Payload)
}
func (o *PatchExternalJWTSignerNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchExternalJWTSignerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchExternalJWTSignerTooManyRequests creates a PatchExternalJWTSignerTooManyRequests with default headers values
func NewPatchExternalJWTSignerTooManyRequests() *PatchExternalJWTSignerTooManyRequests {
	return &PatchExternalJWTSignerTooManyRequests{}
}

/* PatchExternalJWTSignerTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchExternalJWTSignerTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchExternalJWTSignerTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /external-jwt-signers/{id}][%d] patchExternalJwtSignerTooManyRequests  %+v", 429, o.Payload)
}
func (o *PatchExternalJWTSignerTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchExternalJWTSignerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
