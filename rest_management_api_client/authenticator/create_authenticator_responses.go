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

package authenticator

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

// CreateAuthenticatorReader is a Reader for the CreateAuthenticator structure.
type CreateAuthenticatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthenticatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAuthenticatorCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAuthenticatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAuthenticatorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateAuthenticatorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /authenticators] createAuthenticator", response, response.Code())
	}
}

// NewCreateAuthenticatorCreated creates a CreateAuthenticatorCreated with default headers values
func NewCreateAuthenticatorCreated() *CreateAuthenticatorCreated {
	return &CreateAuthenticatorCreated{}
}

/*
CreateAuthenticatorCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateAuthenticatorCreated struct {
	Payload *rest_model.CreateEnvelope
}

// IsSuccess returns true when this create authenticator created response has a 2xx status code
func (o *CreateAuthenticatorCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create authenticator created response has a 3xx status code
func (o *CreateAuthenticatorCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authenticator created response has a 4xx status code
func (o *CreateAuthenticatorCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create authenticator created response has a 5xx status code
func (o *CreateAuthenticatorCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create authenticator created response a status code equal to that given
func (o *CreateAuthenticatorCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create authenticator created response
func (o *CreateAuthenticatorCreated) Code() int {
	return 201
}

func (o *CreateAuthenticatorCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorCreated %s", 201, payload)
}

func (o *CreateAuthenticatorCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorCreated %s", 201, payload)
}

func (o *CreateAuthenticatorCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateAuthenticatorCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticatorBadRequest creates a CreateAuthenticatorBadRequest with default headers values
func NewCreateAuthenticatorBadRequest() *CreateAuthenticatorBadRequest {
	return &CreateAuthenticatorBadRequest{}
}

/*
CreateAuthenticatorBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateAuthenticatorBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create authenticator bad request response has a 2xx status code
func (o *CreateAuthenticatorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create authenticator bad request response has a 3xx status code
func (o *CreateAuthenticatorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authenticator bad request response has a 4xx status code
func (o *CreateAuthenticatorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create authenticator bad request response has a 5xx status code
func (o *CreateAuthenticatorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create authenticator bad request response a status code equal to that given
func (o *CreateAuthenticatorBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create authenticator bad request response
func (o *CreateAuthenticatorBadRequest) Code() int {
	return 400
}

func (o *CreateAuthenticatorBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorBadRequest %s", 400, payload)
}

func (o *CreateAuthenticatorBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorBadRequest %s", 400, payload)
}

func (o *CreateAuthenticatorBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateAuthenticatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticatorUnauthorized creates a CreateAuthenticatorUnauthorized with default headers values
func NewCreateAuthenticatorUnauthorized() *CreateAuthenticatorUnauthorized {
	return &CreateAuthenticatorUnauthorized{}
}

/*
CreateAuthenticatorUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type CreateAuthenticatorUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create authenticator unauthorized response has a 2xx status code
func (o *CreateAuthenticatorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create authenticator unauthorized response has a 3xx status code
func (o *CreateAuthenticatorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authenticator unauthorized response has a 4xx status code
func (o *CreateAuthenticatorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create authenticator unauthorized response has a 5xx status code
func (o *CreateAuthenticatorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create authenticator unauthorized response a status code equal to that given
func (o *CreateAuthenticatorUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create authenticator unauthorized response
func (o *CreateAuthenticatorUnauthorized) Code() int {
	return 401
}

func (o *CreateAuthenticatorUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorUnauthorized %s", 401, payload)
}

func (o *CreateAuthenticatorUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorUnauthorized %s", 401, payload)
}

func (o *CreateAuthenticatorUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateAuthenticatorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticatorTooManyRequests creates a CreateAuthenticatorTooManyRequests with default headers values
func NewCreateAuthenticatorTooManyRequests() *CreateAuthenticatorTooManyRequests {
	return &CreateAuthenticatorTooManyRequests{}
}

/*
CreateAuthenticatorTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type CreateAuthenticatorTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this create authenticator too many requests response has a 2xx status code
func (o *CreateAuthenticatorTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create authenticator too many requests response has a 3xx status code
func (o *CreateAuthenticatorTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authenticator too many requests response has a 4xx status code
func (o *CreateAuthenticatorTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create authenticator too many requests response has a 5xx status code
func (o *CreateAuthenticatorTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create authenticator too many requests response a status code equal to that given
func (o *CreateAuthenticatorTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create authenticator too many requests response
func (o *CreateAuthenticatorTooManyRequests) Code() int {
	return 429
}

func (o *CreateAuthenticatorTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorTooManyRequests %s", 429, payload)
}

func (o *CreateAuthenticatorTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators][%d] createAuthenticatorTooManyRequests %s", 429, payload)
}

func (o *CreateAuthenticatorTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateAuthenticatorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
