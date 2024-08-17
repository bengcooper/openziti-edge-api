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

package config

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

// DeleteConfigTypeReader is a Reader for the DeleteConfigType structure.
type DeleteConfigTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfigTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConfigTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConfigTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConfigTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteConfigTypeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConfigTypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /config-types/{id}] deleteConfigType", response, response.Code())
	}
}

// NewDeleteConfigTypeOK creates a DeleteConfigTypeOK with default headers values
func NewDeleteConfigTypeOK() *DeleteConfigTypeOK {
	return &DeleteConfigTypeOK{}
}

/*
DeleteConfigTypeOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteConfigTypeOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this delete config type o k response has a 2xx status code
func (o *DeleteConfigTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete config type o k response has a 3xx status code
func (o *DeleteConfigTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete config type o k response has a 4xx status code
func (o *DeleteConfigTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete config type o k response has a 5xx status code
func (o *DeleteConfigTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete config type o k response a status code equal to that given
func (o *DeleteConfigTypeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete config type o k response
func (o *DeleteConfigTypeOK) Code() int {
	return 200
}

func (o *DeleteConfigTypeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeOK %s", 200, payload)
}

func (o *DeleteConfigTypeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeOK %s", 200, payload)
}

func (o *DeleteConfigTypeOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteConfigTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigTypeBadRequest creates a DeleteConfigTypeBadRequest with default headers values
func NewDeleteConfigTypeBadRequest() *DeleteConfigTypeBadRequest {
	return &DeleteConfigTypeBadRequest{}
}

/*
DeleteConfigTypeBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteConfigTypeBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete config type bad request response has a 2xx status code
func (o *DeleteConfigTypeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete config type bad request response has a 3xx status code
func (o *DeleteConfigTypeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete config type bad request response has a 4xx status code
func (o *DeleteConfigTypeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete config type bad request response has a 5xx status code
func (o *DeleteConfigTypeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete config type bad request response a status code equal to that given
func (o *DeleteConfigTypeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete config type bad request response
func (o *DeleteConfigTypeBadRequest) Code() int {
	return 400
}

func (o *DeleteConfigTypeBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeBadRequest %s", 400, payload)
}

func (o *DeleteConfigTypeBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeBadRequest %s", 400, payload)
}

func (o *DeleteConfigTypeBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigTypeUnauthorized creates a DeleteConfigTypeUnauthorized with default headers values
func NewDeleteConfigTypeUnauthorized() *DeleteConfigTypeUnauthorized {
	return &DeleteConfigTypeUnauthorized{}
}

/*
DeleteConfigTypeUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteConfigTypeUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete config type unauthorized response has a 2xx status code
func (o *DeleteConfigTypeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete config type unauthorized response has a 3xx status code
func (o *DeleteConfigTypeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete config type unauthorized response has a 4xx status code
func (o *DeleteConfigTypeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete config type unauthorized response has a 5xx status code
func (o *DeleteConfigTypeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete config type unauthorized response a status code equal to that given
func (o *DeleteConfigTypeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete config type unauthorized response
func (o *DeleteConfigTypeUnauthorized) Code() int {
	return 401
}

func (o *DeleteConfigTypeUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeUnauthorized %s", 401, payload)
}

func (o *DeleteConfigTypeUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeUnauthorized %s", 401, payload)
}

func (o *DeleteConfigTypeUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigTypeConflict creates a DeleteConfigTypeConflict with default headers values
func NewDeleteConfigTypeConflict() *DeleteConfigTypeConflict {
	return &DeleteConfigTypeConflict{}
}

/*
DeleteConfigTypeConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteConfigTypeConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete config type conflict response has a 2xx status code
func (o *DeleteConfigTypeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete config type conflict response has a 3xx status code
func (o *DeleteConfigTypeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete config type conflict response has a 4xx status code
func (o *DeleteConfigTypeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete config type conflict response has a 5xx status code
func (o *DeleteConfigTypeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete config type conflict response a status code equal to that given
func (o *DeleteConfigTypeConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete config type conflict response
func (o *DeleteConfigTypeConflict) Code() int {
	return 409
}

func (o *DeleteConfigTypeConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeConflict %s", 409, payload)
}

func (o *DeleteConfigTypeConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeConflict %s", 409, payload)
}

func (o *DeleteConfigTypeConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigTypeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigTypeTooManyRequests creates a DeleteConfigTypeTooManyRequests with default headers values
func NewDeleteConfigTypeTooManyRequests() *DeleteConfigTypeTooManyRequests {
	return &DeleteConfigTypeTooManyRequests{}
}

/*
DeleteConfigTypeTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteConfigTypeTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete config type too many requests response has a 2xx status code
func (o *DeleteConfigTypeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete config type too many requests response has a 3xx status code
func (o *DeleteConfigTypeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete config type too many requests response has a 4xx status code
func (o *DeleteConfigTypeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete config type too many requests response has a 5xx status code
func (o *DeleteConfigTypeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete config type too many requests response a status code equal to that given
func (o *DeleteConfigTypeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete config type too many requests response
func (o *DeleteConfigTypeTooManyRequests) Code() int {
	return 429
}

func (o *DeleteConfigTypeTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeTooManyRequests %s", 429, payload)
}

func (o *DeleteConfigTypeTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /config-types/{id}][%d] deleteConfigTypeTooManyRequests %s", 429, payload)
}

func (o *DeleteConfigTypeTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteConfigTypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
