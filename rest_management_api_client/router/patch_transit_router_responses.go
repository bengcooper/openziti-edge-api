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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// PatchTransitRouterReader is a Reader for the PatchTransitRouter structure.
type PatchTransitRouterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTransitRouterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTransitRouterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchTransitRouterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchTransitRouterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchTransitRouterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchTransitRouterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /transit-routers/{id}] patchTransitRouter", response, response.Code())
	}
}

// NewPatchTransitRouterOK creates a PatchTransitRouterOK with default headers values
func NewPatchTransitRouterOK() *PatchTransitRouterOK {
	return &PatchTransitRouterOK{}
}

/*
PatchTransitRouterOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchTransitRouterOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this patch transit router o k response has a 2xx status code
func (o *PatchTransitRouterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch transit router o k response has a 3xx status code
func (o *PatchTransitRouterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch transit router o k response has a 4xx status code
func (o *PatchTransitRouterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch transit router o k response has a 5xx status code
func (o *PatchTransitRouterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch transit router o k response a status code equal to that given
func (o *PatchTransitRouterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch transit router o k response
func (o *PatchTransitRouterOK) Code() int {
	return 200
}

func (o *PatchTransitRouterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterOK %s", 200, payload)
}

func (o *PatchTransitRouterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterOK %s", 200, payload)
}

func (o *PatchTransitRouterOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchTransitRouterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterBadRequest creates a PatchTransitRouterBadRequest with default headers values
func NewPatchTransitRouterBadRequest() *PatchTransitRouterBadRequest {
	return &PatchTransitRouterBadRequest{}
}

/*
PatchTransitRouterBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchTransitRouterBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch transit router bad request response has a 2xx status code
func (o *PatchTransitRouterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch transit router bad request response has a 3xx status code
func (o *PatchTransitRouterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch transit router bad request response has a 4xx status code
func (o *PatchTransitRouterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch transit router bad request response has a 5xx status code
func (o *PatchTransitRouterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch transit router bad request response a status code equal to that given
func (o *PatchTransitRouterBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch transit router bad request response
func (o *PatchTransitRouterBadRequest) Code() int {
	return 400
}

func (o *PatchTransitRouterBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterBadRequest %s", 400, payload)
}

func (o *PatchTransitRouterBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterBadRequest %s", 400, payload)
}

func (o *PatchTransitRouterBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterUnauthorized creates a PatchTransitRouterUnauthorized with default headers values
func NewPatchTransitRouterUnauthorized() *PatchTransitRouterUnauthorized {
	return &PatchTransitRouterUnauthorized{}
}

/*
PatchTransitRouterUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type PatchTransitRouterUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch transit router unauthorized response has a 2xx status code
func (o *PatchTransitRouterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch transit router unauthorized response has a 3xx status code
func (o *PatchTransitRouterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch transit router unauthorized response has a 4xx status code
func (o *PatchTransitRouterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch transit router unauthorized response has a 5xx status code
func (o *PatchTransitRouterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch transit router unauthorized response a status code equal to that given
func (o *PatchTransitRouterUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch transit router unauthorized response
func (o *PatchTransitRouterUnauthorized) Code() int {
	return 401
}

func (o *PatchTransitRouterUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterUnauthorized %s", 401, payload)
}

func (o *PatchTransitRouterUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterUnauthorized %s", 401, payload)
}

func (o *PatchTransitRouterUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterNotFound creates a PatchTransitRouterNotFound with default headers values
func NewPatchTransitRouterNotFound() *PatchTransitRouterNotFound {
	return &PatchTransitRouterNotFound{}
}

/*
PatchTransitRouterNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchTransitRouterNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch transit router not found response has a 2xx status code
func (o *PatchTransitRouterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch transit router not found response has a 3xx status code
func (o *PatchTransitRouterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch transit router not found response has a 4xx status code
func (o *PatchTransitRouterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch transit router not found response has a 5xx status code
func (o *PatchTransitRouterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch transit router not found response a status code equal to that given
func (o *PatchTransitRouterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch transit router not found response
func (o *PatchTransitRouterNotFound) Code() int {
	return 404
}

func (o *PatchTransitRouterNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterNotFound %s", 404, payload)
}

func (o *PatchTransitRouterNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterNotFound %s", 404, payload)
}

func (o *PatchTransitRouterNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransitRouterTooManyRequests creates a PatchTransitRouterTooManyRequests with default headers values
func NewPatchTransitRouterTooManyRequests() *PatchTransitRouterTooManyRequests {
	return &PatchTransitRouterTooManyRequests{}
}

/*
PatchTransitRouterTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchTransitRouterTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch transit router too many requests response has a 2xx status code
func (o *PatchTransitRouterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch transit router too many requests response has a 3xx status code
func (o *PatchTransitRouterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch transit router too many requests response has a 4xx status code
func (o *PatchTransitRouterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch transit router too many requests response has a 5xx status code
func (o *PatchTransitRouterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch transit router too many requests response a status code equal to that given
func (o *PatchTransitRouterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the patch transit router too many requests response
func (o *PatchTransitRouterTooManyRequests) Code() int {
	return 429
}

func (o *PatchTransitRouterTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterTooManyRequests %s", 429, payload)
}

func (o *PatchTransitRouterTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /transit-routers/{id}][%d] patchTransitRouterTooManyRequests %s", 429, payload)
}

func (o *PatchTransitRouterTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchTransitRouterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
