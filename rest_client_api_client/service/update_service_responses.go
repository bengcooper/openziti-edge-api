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

package service

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

// UpdateServiceReader is a Reader for the UpdateService structure.
type UpdateServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateServiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /services/{id}] updateService", response, response.Code())
	}
}

// NewUpdateServiceOK creates a UpdateServiceOK with default headers values
func NewUpdateServiceOK() *UpdateServiceOK {
	return &UpdateServiceOK{}
}

/*
UpdateServiceOK describes a response with status code 200, with default header values.

The update request was successful and the resource has been altered
*/
type UpdateServiceOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this update service o k response has a 2xx status code
func (o *UpdateServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update service o k response has a 3xx status code
func (o *UpdateServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service o k response has a 4xx status code
func (o *UpdateServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service o k response has a 5xx status code
func (o *UpdateServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update service o k response a status code equal to that given
func (o *UpdateServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update service o k response
func (o *UpdateServiceOK) Code() int {
	return 200
}

func (o *UpdateServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceOK %s", 200, payload)
}

func (o *UpdateServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceOK %s", 200, payload)
}

func (o *UpdateServiceOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdateServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceBadRequest creates a UpdateServiceBadRequest with default headers values
func NewUpdateServiceBadRequest() *UpdateServiceBadRequest {
	return &UpdateServiceBadRequest{}
}

/*
UpdateServiceBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdateServiceBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service bad request response has a 2xx status code
func (o *UpdateServiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service bad request response has a 3xx status code
func (o *UpdateServiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service bad request response has a 4xx status code
func (o *UpdateServiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service bad request response has a 5xx status code
func (o *UpdateServiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update service bad request response a status code equal to that given
func (o *UpdateServiceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update service bad request response
func (o *UpdateServiceBadRequest) Code() int {
	return 400
}

func (o *UpdateServiceBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceBadRequest %s", 400, payload)
}

func (o *UpdateServiceBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceBadRequest %s", 400, payload)
}

func (o *UpdateServiceBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceUnauthorized creates a UpdateServiceUnauthorized with default headers values
func NewUpdateServiceUnauthorized() *UpdateServiceUnauthorized {
	return &UpdateServiceUnauthorized{}
}

/*
UpdateServiceUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type UpdateServiceUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service unauthorized response has a 2xx status code
func (o *UpdateServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service unauthorized response has a 3xx status code
func (o *UpdateServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service unauthorized response has a 4xx status code
func (o *UpdateServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service unauthorized response has a 5xx status code
func (o *UpdateServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update service unauthorized response a status code equal to that given
func (o *UpdateServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update service unauthorized response
func (o *UpdateServiceUnauthorized) Code() int {
	return 401
}

func (o *UpdateServiceUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceUnauthorized %s", 401, payload)
}

func (o *UpdateServiceUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceUnauthorized %s", 401, payload)
}

func (o *UpdateServiceUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceNotFound creates a UpdateServiceNotFound with default headers values
func NewUpdateServiceNotFound() *UpdateServiceNotFound {
	return &UpdateServiceNotFound{}
}

/*
UpdateServiceNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type UpdateServiceNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service not found response has a 2xx status code
func (o *UpdateServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service not found response has a 3xx status code
func (o *UpdateServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service not found response has a 4xx status code
func (o *UpdateServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service not found response has a 5xx status code
func (o *UpdateServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update service not found response a status code equal to that given
func (o *UpdateServiceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update service not found response
func (o *UpdateServiceNotFound) Code() int {
	return 404
}

func (o *UpdateServiceNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceNotFound %s", 404, payload)
}

func (o *UpdateServiceNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceNotFound %s", 404, payload)
}

func (o *UpdateServiceNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceTooManyRequests creates a UpdateServiceTooManyRequests with default headers values
func NewUpdateServiceTooManyRequests() *UpdateServiceTooManyRequests {
	return &UpdateServiceTooManyRequests{}
}

/*
UpdateServiceTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type UpdateServiceTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this update service too many requests response has a 2xx status code
func (o *UpdateServiceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service too many requests response has a 3xx status code
func (o *UpdateServiceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service too many requests response has a 4xx status code
func (o *UpdateServiceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service too many requests response has a 5xx status code
func (o *UpdateServiceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update service too many requests response a status code equal to that given
func (o *UpdateServiceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update service too many requests response
func (o *UpdateServiceTooManyRequests) Code() int {
	return 429
}

func (o *UpdateServiceTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceTooManyRequests %s", 429, payload)
}

func (o *UpdateServiceTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /services/{id}][%d] updateServiceTooManyRequests %s", 429, payload)
}

func (o *UpdateServiceTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
