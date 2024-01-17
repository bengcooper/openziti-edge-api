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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteRouterReader is a Reader for the DeleteRouter structure.
type DeleteRouterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRouterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRouterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRouterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRouterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteRouterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRouterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRouterOK creates a DeleteRouterOK with default headers values
func NewDeleteRouterOK() *DeleteRouterOK {
	return &DeleteRouterOK{}
}

/* DeleteRouterOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteRouterOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteRouterOK) Error() string {
	return fmt.Sprintf("[DELETE /routers/{id}][%d] deleteRouterOK  %+v", 200, o.Payload)
}
func (o *DeleteRouterOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteRouterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouterBadRequest creates a DeleteRouterBadRequest with default headers values
func NewDeleteRouterBadRequest() *DeleteRouterBadRequest {
	return &DeleteRouterBadRequest{}
}

/* DeleteRouterBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteRouterBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteRouterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /routers/{id}][%d] deleteRouterBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteRouterBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteRouterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouterUnauthorized creates a DeleteRouterUnauthorized with default headers values
func NewDeleteRouterUnauthorized() *DeleteRouterUnauthorized {
	return &DeleteRouterUnauthorized{}
}

/* DeleteRouterUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DeleteRouterUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteRouterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /routers/{id}][%d] deleteRouterUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteRouterUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteRouterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouterConflict creates a DeleteRouterConflict with default headers values
func NewDeleteRouterConflict() *DeleteRouterConflict {
	return &DeleteRouterConflict{}
}

/* DeleteRouterConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteRouterConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteRouterConflict) Error() string {
	return fmt.Sprintf("[DELETE /routers/{id}][%d] deleteRouterConflict  %+v", 409, o.Payload)
}
func (o *DeleteRouterConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteRouterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouterTooManyRequests creates a DeleteRouterTooManyRequests with default headers values
func NewDeleteRouterTooManyRequests() *DeleteRouterTooManyRequests {
	return &DeleteRouterTooManyRequests{}
}

/* DeleteRouterTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteRouterTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteRouterTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /routers/{id}][%d] deleteRouterTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteRouterTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteRouterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
