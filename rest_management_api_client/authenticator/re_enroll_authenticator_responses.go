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

// ReEnrollAuthenticatorReader is a Reader for the ReEnrollAuthenticator structure.
type ReEnrollAuthenticatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReEnrollAuthenticatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewReEnrollAuthenticatorCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReEnrollAuthenticatorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReEnrollAuthenticatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReEnrollAuthenticatorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /authenticators/{id}/re-enroll] reEnrollAuthenticator", response, response.Code())
	}
}

// NewReEnrollAuthenticatorCreated creates a ReEnrollAuthenticatorCreated with default headers values
func NewReEnrollAuthenticatorCreated() *ReEnrollAuthenticatorCreated {
	return &ReEnrollAuthenticatorCreated{}
}

/*
ReEnrollAuthenticatorCreated describes a response with status code 201, with default header values.

The create request was successful and the resource has been added at the following location
*/
type ReEnrollAuthenticatorCreated struct {
	Payload *rest_model.CreateEnvelope
}

// IsSuccess returns true when this re enroll authenticator created response has a 2xx status code
func (o *ReEnrollAuthenticatorCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this re enroll authenticator created response has a 3xx status code
func (o *ReEnrollAuthenticatorCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this re enroll authenticator created response has a 4xx status code
func (o *ReEnrollAuthenticatorCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this re enroll authenticator created response has a 5xx status code
func (o *ReEnrollAuthenticatorCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this re enroll authenticator created response a status code equal to that given
func (o *ReEnrollAuthenticatorCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the re enroll authenticator created response
func (o *ReEnrollAuthenticatorCreated) Code() int {
	return 201
}

func (o *ReEnrollAuthenticatorCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorCreated %s", 201, payload)
}

func (o *ReEnrollAuthenticatorCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorCreated %s", 201, payload)
}

func (o *ReEnrollAuthenticatorCreated) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *ReEnrollAuthenticatorCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReEnrollAuthenticatorUnauthorized creates a ReEnrollAuthenticatorUnauthorized with default headers values
func NewReEnrollAuthenticatorUnauthorized() *ReEnrollAuthenticatorUnauthorized {
	return &ReEnrollAuthenticatorUnauthorized{}
}

/*
ReEnrollAuthenticatorUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ReEnrollAuthenticatorUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this re enroll authenticator unauthorized response has a 2xx status code
func (o *ReEnrollAuthenticatorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this re enroll authenticator unauthorized response has a 3xx status code
func (o *ReEnrollAuthenticatorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this re enroll authenticator unauthorized response has a 4xx status code
func (o *ReEnrollAuthenticatorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this re enroll authenticator unauthorized response has a 5xx status code
func (o *ReEnrollAuthenticatorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this re enroll authenticator unauthorized response a status code equal to that given
func (o *ReEnrollAuthenticatorUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the re enroll authenticator unauthorized response
func (o *ReEnrollAuthenticatorUnauthorized) Code() int {
	return 401
}

func (o *ReEnrollAuthenticatorUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorUnauthorized %s", 401, payload)
}

func (o *ReEnrollAuthenticatorUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorUnauthorized %s", 401, payload)
}

func (o *ReEnrollAuthenticatorUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ReEnrollAuthenticatorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReEnrollAuthenticatorNotFound creates a ReEnrollAuthenticatorNotFound with default headers values
func NewReEnrollAuthenticatorNotFound() *ReEnrollAuthenticatorNotFound {
	return &ReEnrollAuthenticatorNotFound{}
}

/*
ReEnrollAuthenticatorNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ReEnrollAuthenticatorNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this re enroll authenticator not found response has a 2xx status code
func (o *ReEnrollAuthenticatorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this re enroll authenticator not found response has a 3xx status code
func (o *ReEnrollAuthenticatorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this re enroll authenticator not found response has a 4xx status code
func (o *ReEnrollAuthenticatorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this re enroll authenticator not found response has a 5xx status code
func (o *ReEnrollAuthenticatorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this re enroll authenticator not found response a status code equal to that given
func (o *ReEnrollAuthenticatorNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the re enroll authenticator not found response
func (o *ReEnrollAuthenticatorNotFound) Code() int {
	return 404
}

func (o *ReEnrollAuthenticatorNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorNotFound %s", 404, payload)
}

func (o *ReEnrollAuthenticatorNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorNotFound %s", 404, payload)
}

func (o *ReEnrollAuthenticatorNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ReEnrollAuthenticatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReEnrollAuthenticatorTooManyRequests creates a ReEnrollAuthenticatorTooManyRequests with default headers values
func NewReEnrollAuthenticatorTooManyRequests() *ReEnrollAuthenticatorTooManyRequests {
	return &ReEnrollAuthenticatorTooManyRequests{}
}

/*
ReEnrollAuthenticatorTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ReEnrollAuthenticatorTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this re enroll authenticator too many requests response has a 2xx status code
func (o *ReEnrollAuthenticatorTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this re enroll authenticator too many requests response has a 3xx status code
func (o *ReEnrollAuthenticatorTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this re enroll authenticator too many requests response has a 4xx status code
func (o *ReEnrollAuthenticatorTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this re enroll authenticator too many requests response has a 5xx status code
func (o *ReEnrollAuthenticatorTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this re enroll authenticator too many requests response a status code equal to that given
func (o *ReEnrollAuthenticatorTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the re enroll authenticator too many requests response
func (o *ReEnrollAuthenticatorTooManyRequests) Code() int {
	return 429
}

func (o *ReEnrollAuthenticatorTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorTooManyRequests %s", 429, payload)
}

func (o *ReEnrollAuthenticatorTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /authenticators/{id}/re-enroll][%d] reEnrollAuthenticatorTooManyRequests %s", 429, payload)
}

func (o *ReEnrollAuthenticatorTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ReEnrollAuthenticatorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
