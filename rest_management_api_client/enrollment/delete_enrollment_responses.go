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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteEnrollmentReader is a Reader for the DeleteEnrollment structure.
type DeleteEnrollmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnrollmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEnrollmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEnrollmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteEnrollmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteEnrollmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /enrollments/{id}] deleteEnrollment", response, response.Code())
	}
}

// NewDeleteEnrollmentOK creates a DeleteEnrollmentOK with default headers values
func NewDeleteEnrollmentOK() *DeleteEnrollmentOK {
	return &DeleteEnrollmentOK{}
}

/*
DeleteEnrollmentOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteEnrollmentOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this delete enrollment o k response has a 2xx status code
func (o *DeleteEnrollmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete enrollment o k response has a 3xx status code
func (o *DeleteEnrollmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete enrollment o k response has a 4xx status code
func (o *DeleteEnrollmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete enrollment o k response has a 5xx status code
func (o *DeleteEnrollmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete enrollment o k response a status code equal to that given
func (o *DeleteEnrollmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete enrollment o k response
func (o *DeleteEnrollmentOK) Code() int {
	return 200
}

func (o *DeleteEnrollmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentOK %s", 200, payload)
}

func (o *DeleteEnrollmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentOK %s", 200, payload)
}

func (o *DeleteEnrollmentOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteEnrollmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentBadRequest creates a DeleteEnrollmentBadRequest with default headers values
func NewDeleteEnrollmentBadRequest() *DeleteEnrollmentBadRequest {
	return &DeleteEnrollmentBadRequest{}
}

/*
DeleteEnrollmentBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteEnrollmentBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete enrollment bad request response has a 2xx status code
func (o *DeleteEnrollmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete enrollment bad request response has a 3xx status code
func (o *DeleteEnrollmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete enrollment bad request response has a 4xx status code
func (o *DeleteEnrollmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete enrollment bad request response has a 5xx status code
func (o *DeleteEnrollmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete enrollment bad request response a status code equal to that given
func (o *DeleteEnrollmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete enrollment bad request response
func (o *DeleteEnrollmentBadRequest) Code() int {
	return 400
}

func (o *DeleteEnrollmentBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentBadRequest %s", 400, payload)
}

func (o *DeleteEnrollmentBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentBadRequest %s", 400, payload)
}

func (o *DeleteEnrollmentBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentUnauthorized creates a DeleteEnrollmentUnauthorized with default headers values
func NewDeleteEnrollmentUnauthorized() *DeleteEnrollmentUnauthorized {
	return &DeleteEnrollmentUnauthorized{}
}

/*
DeleteEnrollmentUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteEnrollmentUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete enrollment unauthorized response has a 2xx status code
func (o *DeleteEnrollmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete enrollment unauthorized response has a 3xx status code
func (o *DeleteEnrollmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete enrollment unauthorized response has a 4xx status code
func (o *DeleteEnrollmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete enrollment unauthorized response has a 5xx status code
func (o *DeleteEnrollmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete enrollment unauthorized response a status code equal to that given
func (o *DeleteEnrollmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete enrollment unauthorized response
func (o *DeleteEnrollmentUnauthorized) Code() int {
	return 401
}

func (o *DeleteEnrollmentUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentUnauthorized %s", 401, payload)
}

func (o *DeleteEnrollmentUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentUnauthorized %s", 401, payload)
}

func (o *DeleteEnrollmentUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnrollmentTooManyRequests creates a DeleteEnrollmentTooManyRequests with default headers values
func NewDeleteEnrollmentTooManyRequests() *DeleteEnrollmentTooManyRequests {
	return &DeleteEnrollmentTooManyRequests{}
}

/*
DeleteEnrollmentTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteEnrollmentTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this delete enrollment too many requests response has a 2xx status code
func (o *DeleteEnrollmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete enrollment too many requests response has a 3xx status code
func (o *DeleteEnrollmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete enrollment too many requests response has a 4xx status code
func (o *DeleteEnrollmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete enrollment too many requests response has a 5xx status code
func (o *DeleteEnrollmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete enrollment too many requests response a status code equal to that given
func (o *DeleteEnrollmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete enrollment too many requests response
func (o *DeleteEnrollmentTooManyRequests) Code() int {
	return 429
}

func (o *DeleteEnrollmentTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentTooManyRequests %s", 429, payload)
}

func (o *DeleteEnrollmentTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /enrollments/{id}][%d] deleteEnrollmentTooManyRequests %s", 429, payload)
}

func (o *DeleteEnrollmentTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteEnrollmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
