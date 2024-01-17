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

package auth_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteAuthPolicyReader is a Reader for the DeleteAuthPolicy structure.
type DeleteAuthPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAuthPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAuthPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAuthPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAuthPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAuthPolicyOK creates a DeleteAuthPolicyOK with default headers values
func NewDeleteAuthPolicyOK() *DeleteAuthPolicyOK {
	return &DeleteAuthPolicyOK{}
}

/* DeleteAuthPolicyOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteAuthPolicyOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteAuthPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /auth-policies/{id}][%d] deleteAuthPolicyOK  %+v", 200, o.Payload)
}
func (o *DeleteAuthPolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteAuthPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthPolicyBadRequest creates a DeleteAuthPolicyBadRequest with default headers values
func NewDeleteAuthPolicyBadRequest() *DeleteAuthPolicyBadRequest {
	return &DeleteAuthPolicyBadRequest{}
}

/* DeleteAuthPolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteAuthPolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteAuthPolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /auth-policies/{id}][%d] deleteAuthPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAuthPolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteAuthPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthPolicyUnauthorized creates a DeleteAuthPolicyUnauthorized with default headers values
func NewDeleteAuthPolicyUnauthorized() *DeleteAuthPolicyUnauthorized {
	return &DeleteAuthPolicyUnauthorized{}
}

/* DeleteAuthPolicyUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteAuthPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteAuthPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /auth-policies/{id}][%d] deleteAuthPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteAuthPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteAuthPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthPolicyTooManyRequests creates a DeleteAuthPolicyTooManyRequests with default headers values
func NewDeleteAuthPolicyTooManyRequests() *DeleteAuthPolicyTooManyRequests {
	return &DeleteAuthPolicyTooManyRequests{}
}

/* DeleteAuthPolicyTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteAuthPolicyTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteAuthPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /auth-policies/{id}][%d] deleteAuthPolicyTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteAuthPolicyTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteAuthPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
