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

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// CreateEnrollmentReader is a Reader for the CreateEnrollment structure.
type CreateEnrollmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnrollmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEnrollmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEnrollmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEnrollmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEnrollmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateEnrollmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEnrollmentCreated creates a CreateEnrollmentCreated with default headers values
func NewCreateEnrollmentCreated() *CreateEnrollmentCreated {
	return &CreateEnrollmentCreated{}
}

/* CreateEnrollmentCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateEnrollmentCreated struct {
	Payload *rest_model.CreateEnvelope
}

func (o *CreateEnrollmentCreated) Error() string {
	return fmt.Sprintf("[POST /enrollments][%d] createEnrollmentCreated  %+v", 201, o.Payload)
}
func (o *CreateEnrollmentCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateEnrollmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnrollmentBadRequest creates a CreateEnrollmentBadRequest with default headers values
func NewCreateEnrollmentBadRequest() *CreateEnrollmentBadRequest {
	return &CreateEnrollmentBadRequest{}
}

/* CreateEnrollmentBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateEnrollmentBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateEnrollmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /enrollments][%d] createEnrollmentBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEnrollmentBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateEnrollmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnrollmentUnauthorized creates a CreateEnrollmentUnauthorized with default headers values
func NewCreateEnrollmentUnauthorized() *CreateEnrollmentUnauthorized {
	return &CreateEnrollmentUnauthorized{}
}

/* CreateEnrollmentUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type CreateEnrollmentUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateEnrollmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /enrollments][%d] createEnrollmentUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateEnrollmentUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateEnrollmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnrollmentConflict creates a CreateEnrollmentConflict with default headers values
func NewCreateEnrollmentConflict() *CreateEnrollmentConflict {
	return &CreateEnrollmentConflict{}
}

/* CreateEnrollmentConflict describes a response with status code 409, with default header values.

The request could not be completed due to a conflict of configuration or state
*/
type CreateEnrollmentConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateEnrollmentConflict) Error() string {
	return fmt.Sprintf("[POST /enrollments][%d] createEnrollmentConflict  %+v", 409, o.Payload)
}
func (o *CreateEnrollmentConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateEnrollmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnrollmentTooManyRequests creates a CreateEnrollmentTooManyRequests with default headers values
func NewCreateEnrollmentTooManyRequests() *CreateEnrollmentTooManyRequests {
	return &CreateEnrollmentTooManyRequests{}
}

/* CreateEnrollmentTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type CreateEnrollmentTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateEnrollmentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /enrollments][%d] createEnrollmentTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateEnrollmentTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateEnrollmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
