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

package edge_router

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

// PatchEdgeRouterReader is a Reader for the PatchEdgeRouter structure.
type PatchEdgeRouterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEdgeRouterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEdgeRouterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchEdgeRouterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchEdgeRouterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchEdgeRouterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchEdgeRouterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /edge-routers/{id}] patchEdgeRouter", response, response.Code())
	}
}

// NewPatchEdgeRouterOK creates a PatchEdgeRouterOK with default headers values
func NewPatchEdgeRouterOK() *PatchEdgeRouterOK {
	return &PatchEdgeRouterOK{}
}

/*
PatchEdgeRouterOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchEdgeRouterOK struct {
	Payload *rest_model.Empty
}

// IsSuccess returns true when this patch edge router o k response has a 2xx status code
func (o *PatchEdgeRouterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch edge router o k response has a 3xx status code
func (o *PatchEdgeRouterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch edge router o k response has a 4xx status code
func (o *PatchEdgeRouterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch edge router o k response has a 5xx status code
func (o *PatchEdgeRouterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch edge router o k response a status code equal to that given
func (o *PatchEdgeRouterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch edge router o k response
func (o *PatchEdgeRouterOK) Code() int {
	return 200
}

func (o *PatchEdgeRouterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterOK %s", 200, payload)
}

func (o *PatchEdgeRouterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterOK %s", 200, payload)
}

func (o *PatchEdgeRouterOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchEdgeRouterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterBadRequest creates a PatchEdgeRouterBadRequest with default headers values
func NewPatchEdgeRouterBadRequest() *PatchEdgeRouterBadRequest {
	return &PatchEdgeRouterBadRequest{}
}

/*
PatchEdgeRouterBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchEdgeRouterBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch edge router bad request response has a 2xx status code
func (o *PatchEdgeRouterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch edge router bad request response has a 3xx status code
func (o *PatchEdgeRouterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch edge router bad request response has a 4xx status code
func (o *PatchEdgeRouterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch edge router bad request response has a 5xx status code
func (o *PatchEdgeRouterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch edge router bad request response a status code equal to that given
func (o *PatchEdgeRouterBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch edge router bad request response
func (o *PatchEdgeRouterBadRequest) Code() int {
	return 400
}

func (o *PatchEdgeRouterBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterBadRequest %s", 400, payload)
}

func (o *PatchEdgeRouterBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterBadRequest %s", 400, payload)
}

func (o *PatchEdgeRouterBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterUnauthorized creates a PatchEdgeRouterUnauthorized with default headers values
func NewPatchEdgeRouterUnauthorized() *PatchEdgeRouterUnauthorized {
	return &PatchEdgeRouterUnauthorized{}
}

/*
PatchEdgeRouterUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type PatchEdgeRouterUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch edge router unauthorized response has a 2xx status code
func (o *PatchEdgeRouterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch edge router unauthorized response has a 3xx status code
func (o *PatchEdgeRouterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch edge router unauthorized response has a 4xx status code
func (o *PatchEdgeRouterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch edge router unauthorized response has a 5xx status code
func (o *PatchEdgeRouterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch edge router unauthorized response a status code equal to that given
func (o *PatchEdgeRouterUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch edge router unauthorized response
func (o *PatchEdgeRouterUnauthorized) Code() int {
	return 401
}

func (o *PatchEdgeRouterUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterUnauthorized %s", 401, payload)
}

func (o *PatchEdgeRouterUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterUnauthorized %s", 401, payload)
}

func (o *PatchEdgeRouterUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterNotFound creates a PatchEdgeRouterNotFound with default headers values
func NewPatchEdgeRouterNotFound() *PatchEdgeRouterNotFound {
	return &PatchEdgeRouterNotFound{}
}

/*
PatchEdgeRouterNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchEdgeRouterNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch edge router not found response has a 2xx status code
func (o *PatchEdgeRouterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch edge router not found response has a 3xx status code
func (o *PatchEdgeRouterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch edge router not found response has a 4xx status code
func (o *PatchEdgeRouterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch edge router not found response has a 5xx status code
func (o *PatchEdgeRouterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch edge router not found response a status code equal to that given
func (o *PatchEdgeRouterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch edge router not found response
func (o *PatchEdgeRouterNotFound) Code() int {
	return 404
}

func (o *PatchEdgeRouterNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterNotFound %s", 404, payload)
}

func (o *PatchEdgeRouterNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterNotFound %s", 404, payload)
}

func (o *PatchEdgeRouterNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterTooManyRequests creates a PatchEdgeRouterTooManyRequests with default headers values
func NewPatchEdgeRouterTooManyRequests() *PatchEdgeRouterTooManyRequests {
	return &PatchEdgeRouterTooManyRequests{}
}

/*
PatchEdgeRouterTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type PatchEdgeRouterTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this patch edge router too many requests response has a 2xx status code
func (o *PatchEdgeRouterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch edge router too many requests response has a 3xx status code
func (o *PatchEdgeRouterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch edge router too many requests response has a 4xx status code
func (o *PatchEdgeRouterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch edge router too many requests response has a 5xx status code
func (o *PatchEdgeRouterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch edge router too many requests response a status code equal to that given
func (o *PatchEdgeRouterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the patch edge router too many requests response
func (o *PatchEdgeRouterTooManyRequests) Code() int {
	return 429
}

func (o *PatchEdgeRouterTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterTooManyRequests %s", 429, payload)
}

func (o *PatchEdgeRouterTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /edge-routers/{id}][%d] patchEdgeRouterTooManyRequests %s", 429, payload)
}

func (o *PatchEdgeRouterTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
