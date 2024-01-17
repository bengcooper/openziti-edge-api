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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// PatchPostureCheckReader is a Reader for the PatchPostureCheck structure.
type PatchPostureCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPostureCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPostureCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchPostureCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchPostureCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchPostureCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchPostureCheckTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchPostureCheckOK creates a PatchPostureCheckOK with default headers values
func NewPatchPostureCheckOK() *PatchPostureCheckOK {
	return &PatchPostureCheckOK{}
}

/* PatchPostureCheckOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchPostureCheckOK struct {
	Payload *rest_model.Empty
}

func (o *PatchPostureCheckOK) Error() string {
	return fmt.Sprintf("[PATCH /posture-checks/{id}][%d] patchPostureCheckOK  %+v", 200, o.Payload)
}
func (o *PatchPostureCheckOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchPostureCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPostureCheckBadRequest creates a PatchPostureCheckBadRequest with default headers values
func NewPatchPostureCheckBadRequest() *PatchPostureCheckBadRequest {
	return &PatchPostureCheckBadRequest{}
}

/* PatchPostureCheckBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchPostureCheckBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchPostureCheckBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /posture-checks/{id}][%d] patchPostureCheckBadRequest  %+v", 400, o.Payload)
}
func (o *PatchPostureCheckBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchPostureCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPostureCheckUnauthorized creates a PatchPostureCheckUnauthorized with default headers values
func NewPatchPostureCheckUnauthorized() *PatchPostureCheckUnauthorized {
	return &PatchPostureCheckUnauthorized{}
}

/* PatchPostureCheckUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type PatchPostureCheckUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchPostureCheckUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /posture-checks/{id}][%d] patchPostureCheckUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchPostureCheckUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchPostureCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPostureCheckNotFound creates a PatchPostureCheckNotFound with default headers values
func NewPatchPostureCheckNotFound() *PatchPostureCheckNotFound {
	return &PatchPostureCheckNotFound{}
}

/* PatchPostureCheckNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchPostureCheckNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchPostureCheckNotFound) Error() string {
	return fmt.Sprintf("[PATCH /posture-checks/{id}][%d] patchPostureCheckNotFound  %+v", 404, o.Payload)
}
func (o *PatchPostureCheckNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchPostureCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPostureCheckTooManyRequests creates a PatchPostureCheckTooManyRequests with default headers values
func NewPatchPostureCheckTooManyRequests() *PatchPostureCheckTooManyRequests {
	return &PatchPostureCheckTooManyRequests{}
}

/* PatchPostureCheckTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchPostureCheckTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchPostureCheckTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /posture-checks/{id}][%d] patchPostureCheckTooManyRequests  %+v", 429, o.Payload)
}
func (o *PatchPostureCheckTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchPostureCheckTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
