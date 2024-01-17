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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// EnableIdentityReader is a Reader for the EnableIdentity structure.
type EnableIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnableIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewEnableIdentityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableIdentityOK creates a EnableIdentityOK with default headers values
func NewEnableIdentityOK() *EnableIdentityOK {
	return &EnableIdentityOK{}
}

/* EnableIdentityOK describes a response with status code 200, with default header values.

Base empty response
*/
type EnableIdentityOK struct {
	Payload *rest_model.Empty
}

func (o *EnableIdentityOK) Error() string {
	return fmt.Sprintf("[POST /identities/{id}/enable][%d] enableIdentityOK  %+v", 200, o.Payload)
}
func (o *EnableIdentityOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *EnableIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableIdentityUnauthorized creates a EnableIdentityUnauthorized with default headers values
func NewEnableIdentityUnauthorized() *EnableIdentityUnauthorized {
	return &EnableIdentityUnauthorized{}
}

/* EnableIdentityUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type EnableIdentityUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnableIdentityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /identities/{id}/enable][%d] enableIdentityUnauthorized  %+v", 401, o.Payload)
}
func (o *EnableIdentityUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnableIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableIdentityNotFound creates a EnableIdentityNotFound with default headers values
func NewEnableIdentityNotFound() *EnableIdentityNotFound {
	return &EnableIdentityNotFound{}
}

/* EnableIdentityNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type EnableIdentityNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnableIdentityNotFound) Error() string {
	return fmt.Sprintf("[POST /identities/{id}/enable][%d] enableIdentityNotFound  %+v", 404, o.Payload)
}
func (o *EnableIdentityNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnableIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableIdentityTooManyRequests creates a EnableIdentityTooManyRequests with default headers values
func NewEnableIdentityTooManyRequests() *EnableIdentityTooManyRequests {
	return &EnableIdentityTooManyRequests{}
}

/* EnableIdentityTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type EnableIdentityTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *EnableIdentityTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /identities/{id}/enable][%d] enableIdentityTooManyRequests  %+v", 429, o.Payload)
}
func (o *EnableIdentityTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *EnableIdentityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
