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

package service_policy

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

// UpdateServicePolicyReader is a Reader for the UpdateServicePolicy structure.
type UpdateServicePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServicePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServicePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServicePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServicePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServicePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateServicePolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /service-policies/{id}] updateServicePolicy", response, response.Code())
	}
}

// NewUpdateServicePolicyOK creates a UpdateServicePolicyOK with default headers values
func NewUpdateServicePolicyOK() *UpdateServicePolicyOK {
	return &UpdateServicePolicyOK{}
}

/*
UpdateServicePolicyOK describes a response with status code 200, with default header values.

The update request was successful and the resource has been altered
*/
type UpdateServicePolicyOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this update service policy o k response has a 2xx status code
func (o *UpdateServicePolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update service policy o k response has a 3xx status code
func (o *UpdateServicePolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service policy o k response has a 4xx status code
func (o *UpdateServicePolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service policy o k response has a 5xx status code
func (o *UpdateServicePolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update service policy o k response a status code equal to that given
func (o *UpdateServicePolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update service policy o k response
func (o *UpdateServicePolicyOK) Code() int {
	return 200
}

func (o *UpdateServicePolicyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyOK %s", 200, payload)
}

func (o *UpdateServicePolicyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyOK %s", 200, payload)
}

func (o *UpdateServicePolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdateServicePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServicePolicyBadRequest creates a UpdateServicePolicyBadRequest with default headers values
func NewUpdateServicePolicyBadRequest() *UpdateServicePolicyBadRequest {
	return &UpdateServicePolicyBadRequest{}
}

/*
UpdateServicePolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdateServicePolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service policy bad request response has a 2xx status code
func (o *UpdateServicePolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service policy bad request response has a 3xx status code
func (o *UpdateServicePolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service policy bad request response has a 4xx status code
func (o *UpdateServicePolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service policy bad request response has a 5xx status code
func (o *UpdateServicePolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update service policy bad request response a status code equal to that given
func (o *UpdateServicePolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update service policy bad request response
func (o *UpdateServicePolicyBadRequest) Code() int {
	return 400
}

func (o *UpdateServicePolicyBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyBadRequest %s", 400, payload)
}

func (o *UpdateServicePolicyBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyBadRequest %s", 400, payload)
}

func (o *UpdateServicePolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServicePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServicePolicyUnauthorized creates a UpdateServicePolicyUnauthorized with default headers values
func NewUpdateServicePolicyUnauthorized() *UpdateServicePolicyUnauthorized {
	return &UpdateServicePolicyUnauthorized{}
}

/*
UpdateServicePolicyUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type UpdateServicePolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service policy unauthorized response has a 2xx status code
func (o *UpdateServicePolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service policy unauthorized response has a 3xx status code
func (o *UpdateServicePolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service policy unauthorized response has a 4xx status code
func (o *UpdateServicePolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service policy unauthorized response has a 5xx status code
func (o *UpdateServicePolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update service policy unauthorized response a status code equal to that given
func (o *UpdateServicePolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update service policy unauthorized response
func (o *UpdateServicePolicyUnauthorized) Code() int {
	return 401
}

func (o *UpdateServicePolicyUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyUnauthorized %s", 401, payload)
}

func (o *UpdateServicePolicyUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyUnauthorized %s", 401, payload)
}

func (o *UpdateServicePolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServicePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServicePolicyNotFound creates a UpdateServicePolicyNotFound with default headers values
func NewUpdateServicePolicyNotFound() *UpdateServicePolicyNotFound {
	return &UpdateServicePolicyNotFound{}
}

/*
UpdateServicePolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type UpdateServicePolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service policy not found response has a 2xx status code
func (o *UpdateServicePolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service policy not found response has a 3xx status code
func (o *UpdateServicePolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service policy not found response has a 4xx status code
func (o *UpdateServicePolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service policy not found response has a 5xx status code
func (o *UpdateServicePolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update service policy not found response a status code equal to that given
func (o *UpdateServicePolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update service policy not found response
func (o *UpdateServicePolicyNotFound) Code() int {
	return 404
}

func (o *UpdateServicePolicyNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyNotFound %s", 404, payload)
}

func (o *UpdateServicePolicyNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyNotFound %s", 404, payload)
}

func (o *UpdateServicePolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServicePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServicePolicyTooManyRequests creates a UpdateServicePolicyTooManyRequests with default headers values
func NewUpdateServicePolicyTooManyRequests() *UpdateServicePolicyTooManyRequests {
	return &UpdateServicePolicyTooManyRequests{}
}

/*
UpdateServicePolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type UpdateServicePolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service policy too many requests response has a 2xx status code
func (o *UpdateServicePolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service policy too many requests response has a 3xx status code
func (o *UpdateServicePolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service policy too many requests response has a 4xx status code
func (o *UpdateServicePolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service policy too many requests response has a 5xx status code
func (o *UpdateServicePolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update service policy too many requests response a status code equal to that given
func (o *UpdateServicePolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update service policy too many requests response
func (o *UpdateServicePolicyTooManyRequests) Code() int {
	return 429
}

func (o *UpdateServicePolicyTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyTooManyRequests %s", 429, payload)
}

func (o *UpdateServicePolicyTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /service-policies/{id}][%d] updateServicePolicyTooManyRequests %s", 429, payload)
}

func (o *UpdateServicePolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServicePolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
