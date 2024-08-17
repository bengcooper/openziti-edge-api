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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DisableIdentityReader is a Reader for the DisableIdentity structure.
type DisableIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDisableIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisableIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDisableIdentityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /identities/{id}/disable] disableIdentity", response, response.Code())
	}
}

// NewDisableIdentityOK creates a DisableIdentityOK with default headers values
func NewDisableIdentityOK() *DisableIdentityOK {
	return &DisableIdentityOK{}
}

/*
DisableIdentityOK describes a response with status code 200, with default header values.

Base empty response
*/
type DisableIdentityOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this disable identity o k response has a 2xx status code
func (o *DisableIdentityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable identity o k response has a 3xx status code
func (o *DisableIdentityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable identity o k response has a 4xx status code
func (o *DisableIdentityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable identity o k response has a 5xx status code
func (o *DisableIdentityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disable identity o k response a status code equal to that given
func (o *DisableIdentityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disable identity o k response
func (o *DisableIdentityOK) Code() int {
	return 200
}

func (o *DisableIdentityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityOK %s", 200, payload)
}

func (o *DisableIdentityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityOK %s", 200, payload)
}

func (o *DisableIdentityOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DisableIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableIdentityUnauthorized creates a DisableIdentityUnauthorized with default headers values
func NewDisableIdentityUnauthorized() *DisableIdentityUnauthorized {
	return &DisableIdentityUnauthorized{}
}

/*
DisableIdentityUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DisableIdentityUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this disable identity unauthorized response has a 2xx status code
func (o *DisableIdentityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable identity unauthorized response has a 3xx status code
func (o *DisableIdentityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable identity unauthorized response has a 4xx status code
func (o *DisableIdentityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable identity unauthorized response has a 5xx status code
func (o *DisableIdentityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this disable identity unauthorized response a status code equal to that given
func (o *DisableIdentityUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the disable identity unauthorized response
func (o *DisableIdentityUnauthorized) Code() int {
	return 401
}

func (o *DisableIdentityUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityUnauthorized %s", 401, payload)
}

func (o *DisableIdentityUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityUnauthorized %s", 401, payload)
}

func (o *DisableIdentityUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DisableIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableIdentityNotFound creates a DisableIdentityNotFound with default headers values
func NewDisableIdentityNotFound() *DisableIdentityNotFound {
	return &DisableIdentityNotFound{}
}

/*
DisableIdentityNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DisableIdentityNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this disable identity not found response has a 2xx status code
func (o *DisableIdentityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable identity not found response has a 3xx status code
func (o *DisableIdentityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable identity not found response has a 4xx status code
func (o *DisableIdentityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable identity not found response has a 5xx status code
func (o *DisableIdentityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this disable identity not found response a status code equal to that given
func (o *DisableIdentityNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the disable identity not found response
func (o *DisableIdentityNotFound) Code() int {
	return 404
}

func (o *DisableIdentityNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityNotFound %s", 404, payload)
}

func (o *DisableIdentityNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityNotFound %s", 404, payload)
}

func (o *DisableIdentityNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DisableIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableIdentityTooManyRequests creates a DisableIdentityTooManyRequests with default headers values
func NewDisableIdentityTooManyRequests() *DisableIdentityTooManyRequests {
	return &DisableIdentityTooManyRequests{}
}

/*
DisableIdentityTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DisableIdentityTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this disable identity too many requests response has a 2xx status code
func (o *DisableIdentityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable identity too many requests response has a 3xx status code
func (o *DisableIdentityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable identity too many requests response has a 4xx status code
func (o *DisableIdentityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable identity too many requests response has a 5xx status code
func (o *DisableIdentityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this disable identity too many requests response a status code equal to that given
func (o *DisableIdentityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the disable identity too many requests response
func (o *DisableIdentityTooManyRequests) Code() int {
	return 429
}

func (o *DisableIdentityTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityTooManyRequests %s", 429, payload)
}

func (o *DisableIdentityTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identities/{id}/disable][%d] disableIdentityTooManyRequests %s", 429, payload)
}

func (o *DisableIdentityTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DisableIdentityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
