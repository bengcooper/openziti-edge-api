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

// GetIdentityAuthenticatorsReader is a Reader for the GetIdentityAuthenticators structure.
type GetIdentityAuthenticatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityAuthenticatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityAuthenticatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetIdentityAuthenticatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityAuthenticatorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdentityAuthenticatorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /identities/{id}/authenticators] getIdentityAuthenticators", response, response.Code())
	}
}

// NewGetIdentityAuthenticatorsOK creates a GetIdentityAuthenticatorsOK with default headers values
func NewGetIdentityAuthenticatorsOK() *GetIdentityAuthenticatorsOK {
	return &GetIdentityAuthenticatorsOK{}
}

/*
GetIdentityAuthenticatorsOK describes a response with status code 200, with default header values.

A list of authenticators
*/
type GetIdentityAuthenticatorsOK struct {
	Payload *rest_model.ListAuthenticatorsEnvelope
}

// IsSuccess returns true when this get identity authenticators o k response has a 2xx status code
func (o *GetIdentityAuthenticatorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get identity authenticators o k response has a 3xx status code
func (o *GetIdentityAuthenticatorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity authenticators o k response has a 4xx status code
func (o *GetIdentityAuthenticatorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity authenticators o k response has a 5xx status code
func (o *GetIdentityAuthenticatorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity authenticators o k response a status code equal to that given
func (o *GetIdentityAuthenticatorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get identity authenticators o k response
func (o *GetIdentityAuthenticatorsOK) Code() int {
	return 200
}

func (o *GetIdentityAuthenticatorsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsOK %s", 200, payload)
}

func (o *GetIdentityAuthenticatorsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsOK %s", 200, payload)
}

func (o *GetIdentityAuthenticatorsOK) GetPayload() *rest_model.ListAuthenticatorsEnvelope {
	return o.Payload
}

func (o *GetIdentityAuthenticatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListAuthenticatorsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityAuthenticatorsUnauthorized creates a GetIdentityAuthenticatorsUnauthorized with default headers values
func NewGetIdentityAuthenticatorsUnauthorized() *GetIdentityAuthenticatorsUnauthorized {
	return &GetIdentityAuthenticatorsUnauthorized{}
}

/*
GetIdentityAuthenticatorsUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type GetIdentityAuthenticatorsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this get identity authenticators unauthorized response has a 2xx status code
func (o *GetIdentityAuthenticatorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity authenticators unauthorized response has a 3xx status code
func (o *GetIdentityAuthenticatorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity authenticators unauthorized response has a 4xx status code
func (o *GetIdentityAuthenticatorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity authenticators unauthorized response has a 5xx status code
func (o *GetIdentityAuthenticatorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity authenticators unauthorized response a status code equal to that given
func (o *GetIdentityAuthenticatorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get identity authenticators unauthorized response
func (o *GetIdentityAuthenticatorsUnauthorized) Code() int {
	return 401
}

func (o *GetIdentityAuthenticatorsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsUnauthorized %s", 401, payload)
}

func (o *GetIdentityAuthenticatorsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsUnauthorized %s", 401, payload)
}

func (o *GetIdentityAuthenticatorsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetIdentityAuthenticatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityAuthenticatorsNotFound creates a GetIdentityAuthenticatorsNotFound with default headers values
func NewGetIdentityAuthenticatorsNotFound() *GetIdentityAuthenticatorsNotFound {
	return &GetIdentityAuthenticatorsNotFound{}
}

/*
GetIdentityAuthenticatorsNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type GetIdentityAuthenticatorsNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this get identity authenticators not found response has a 2xx status code
func (o *GetIdentityAuthenticatorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity authenticators not found response has a 3xx status code
func (o *GetIdentityAuthenticatorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity authenticators not found response has a 4xx status code
func (o *GetIdentityAuthenticatorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity authenticators not found response has a 5xx status code
func (o *GetIdentityAuthenticatorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity authenticators not found response a status code equal to that given
func (o *GetIdentityAuthenticatorsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get identity authenticators not found response
func (o *GetIdentityAuthenticatorsNotFound) Code() int {
	return 404
}

func (o *GetIdentityAuthenticatorsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsNotFound %s", 404, payload)
}

func (o *GetIdentityAuthenticatorsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsNotFound %s", 404, payload)
}

func (o *GetIdentityAuthenticatorsNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetIdentityAuthenticatorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityAuthenticatorsTooManyRequests creates a GetIdentityAuthenticatorsTooManyRequests with default headers values
func NewGetIdentityAuthenticatorsTooManyRequests() *GetIdentityAuthenticatorsTooManyRequests {
	return &GetIdentityAuthenticatorsTooManyRequests{}
}

/*
GetIdentityAuthenticatorsTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type GetIdentityAuthenticatorsTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this get identity authenticators too many requests response has a 2xx status code
func (o *GetIdentityAuthenticatorsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity authenticators too many requests response has a 3xx status code
func (o *GetIdentityAuthenticatorsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity authenticators too many requests response has a 4xx status code
func (o *GetIdentityAuthenticatorsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity authenticators too many requests response has a 5xx status code
func (o *GetIdentityAuthenticatorsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity authenticators too many requests response a status code equal to that given
func (o *GetIdentityAuthenticatorsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get identity authenticators too many requests response
func (o *GetIdentityAuthenticatorsTooManyRequests) Code() int {
	return 429
}

func (o *GetIdentityAuthenticatorsTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsTooManyRequests %s", 429, payload)
}

func (o *GetIdentityAuthenticatorsTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /identities/{id}/authenticators][%d] getIdentityAuthenticatorsTooManyRequests %s", 429, payload)
}

func (o *GetIdentityAuthenticatorsTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetIdentityAuthenticatorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
