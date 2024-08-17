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

// DeleteIdentityReader is a Reader for the DeleteIdentity structure.
type DeleteIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteIdentityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /identities/{id}] deleteIdentity", response, response.Code())
	}
}

// NewDeleteIdentityOK creates a DeleteIdentityOK with default headers values
func NewDeleteIdentityOK() *DeleteIdentityOK {
	return &DeleteIdentityOK{}
}

/*
DeleteIdentityOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteIdentityOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this delete identity o k response has a 2xx status code
func (o *DeleteIdentityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete identity o k response has a 3xx status code
func (o *DeleteIdentityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identity o k response has a 4xx status code
func (o *DeleteIdentityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identity o k response has a 5xx status code
func (o *DeleteIdentityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identity o k response a status code equal to that given
func (o *DeleteIdentityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete identity o k response
func (o *DeleteIdentityOK) Code() int {
	return 200
}

func (o *DeleteIdentityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityOK %s", 200, payload)
}

func (o *DeleteIdentityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityOK %s", 200, payload)
}

func (o *DeleteIdentityOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityBadRequest creates a DeleteIdentityBadRequest with default headers values
func NewDeleteIdentityBadRequest() *DeleteIdentityBadRequest {
	return &DeleteIdentityBadRequest{}
}

/*
DeleteIdentityBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteIdentityBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete identity bad request response has a 2xx status code
func (o *DeleteIdentityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identity bad request response has a 3xx status code
func (o *DeleteIdentityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identity bad request response has a 4xx status code
func (o *DeleteIdentityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identity bad request response has a 5xx status code
func (o *DeleteIdentityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identity bad request response a status code equal to that given
func (o *DeleteIdentityBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete identity bad request response
func (o *DeleteIdentityBadRequest) Code() int {
	return 400
}

func (o *DeleteIdentityBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityBadRequest %s", 400, payload)
}

func (o *DeleteIdentityBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityBadRequest %s", 400, payload)
}

func (o *DeleteIdentityBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityUnauthorized creates a DeleteIdentityUnauthorized with default headers values
func NewDeleteIdentityUnauthorized() *DeleteIdentityUnauthorized {
	return &DeleteIdentityUnauthorized{}
}

/*
DeleteIdentityUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteIdentityUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete identity unauthorized response has a 2xx status code
func (o *DeleteIdentityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identity unauthorized response has a 3xx status code
func (o *DeleteIdentityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identity unauthorized response has a 4xx status code
func (o *DeleteIdentityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identity unauthorized response has a 5xx status code
func (o *DeleteIdentityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identity unauthorized response a status code equal to that given
func (o *DeleteIdentityUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete identity unauthorized response
func (o *DeleteIdentityUnauthorized) Code() int {
	return 401
}

func (o *DeleteIdentityUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityUnauthorized %s", 401, payload)
}

func (o *DeleteIdentityUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityUnauthorized %s", 401, payload)
}

func (o *DeleteIdentityUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityConflict creates a DeleteIdentityConflict with default headers values
func NewDeleteIdentityConflict() *DeleteIdentityConflict {
	return &DeleteIdentityConflict{}
}

/*
DeleteIdentityConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteIdentityConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete identity conflict response has a 2xx status code
func (o *DeleteIdentityConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identity conflict response has a 3xx status code
func (o *DeleteIdentityConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identity conflict response has a 4xx status code
func (o *DeleteIdentityConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identity conflict response has a 5xx status code
func (o *DeleteIdentityConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identity conflict response a status code equal to that given
func (o *DeleteIdentityConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete identity conflict response
func (o *DeleteIdentityConflict) Code() int {
	return 409
}

func (o *DeleteIdentityConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityConflict %s", 409, payload)
}

func (o *DeleteIdentityConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityConflict %s", 409, payload)
}

func (o *DeleteIdentityConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityTooManyRequests creates a DeleteIdentityTooManyRequests with default headers values
func NewDeleteIdentityTooManyRequests() *DeleteIdentityTooManyRequests {
	return &DeleteIdentityTooManyRequests{}
}

/*
DeleteIdentityTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteIdentityTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete identity too many requests response has a 2xx status code
func (o *DeleteIdentityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identity too many requests response has a 3xx status code
func (o *DeleteIdentityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identity too many requests response has a 4xx status code
func (o *DeleteIdentityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identity too many requests response has a 5xx status code
func (o *DeleteIdentityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identity too many requests response a status code equal to that given
func (o *DeleteIdentityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete identity too many requests response
func (o *DeleteIdentityTooManyRequests) Code() int {
	return 429
}

func (o *DeleteIdentityTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityTooManyRequests %s", 429, payload)
}

func (o *DeleteIdentityTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /identities/{id}][%d] deleteIdentityTooManyRequests %s", 429, payload)
}

func (o *DeleteIdentityTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteIdentityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
