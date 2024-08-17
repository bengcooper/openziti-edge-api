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

package current_identity

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

// GetCurrentIdentityEdgeRoutersReader is a Reader for the GetCurrentIdentityEdgeRouters structure.
type GetCurrentIdentityEdgeRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentIdentityEdgeRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentIdentityEdgeRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentIdentityEdgeRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCurrentIdentityEdgeRoutersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /current-identity/edge-routers] getCurrentIdentityEdgeRouters", response, response.Code())
	}
}

// NewGetCurrentIdentityEdgeRoutersOK creates a GetCurrentIdentityEdgeRoutersOK with default headers values
func NewGetCurrentIdentityEdgeRoutersOK() *GetCurrentIdentityEdgeRoutersOK {
	return &GetCurrentIdentityEdgeRoutersOK{}
}

/*
GetCurrentIdentityEdgeRoutersOK describes a response with status code 200, with default header values.

A list of edge routers
*/
type GetCurrentIdentityEdgeRoutersOK struct {
	Payload *rest_model.ListCurrentIdentityEdgeRoutersEnvelope
}

// IsSuccess returns true when this get current identity edge routers o k response has a 2xx status code
func (o *GetCurrentIdentityEdgeRoutersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current identity edge routers o k response has a 3xx status code
func (o *GetCurrentIdentityEdgeRoutersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current identity edge routers o k response has a 4xx status code
func (o *GetCurrentIdentityEdgeRoutersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current identity edge routers o k response has a 5xx status code
func (o *GetCurrentIdentityEdgeRoutersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current identity edge routers o k response a status code equal to that given
func (o *GetCurrentIdentityEdgeRoutersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get current identity edge routers o k response
func (o *GetCurrentIdentityEdgeRoutersOK) Code() int {
	return 200
}

func (o *GetCurrentIdentityEdgeRoutersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/edge-routers][%d] getCurrentIdentityEdgeRoutersOK %s", 200, payload)
}

func (o *GetCurrentIdentityEdgeRoutersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/edge-routers][%d] getCurrentIdentityEdgeRoutersOK %s", 200, payload)
}

func (o *GetCurrentIdentityEdgeRoutersOK) GetPayload() *rest_model.ListCurrentIdentityEdgeRoutersEnvelope {
	return o.Payload
}

func (o *GetCurrentIdentityEdgeRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListCurrentIdentityEdgeRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentIdentityEdgeRoutersUnauthorized creates a GetCurrentIdentityEdgeRoutersUnauthorized with default headers values
func NewGetCurrentIdentityEdgeRoutersUnauthorized() *GetCurrentIdentityEdgeRoutersUnauthorized {
	return &GetCurrentIdentityEdgeRoutersUnauthorized{}
}

/*
GetCurrentIdentityEdgeRoutersUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type GetCurrentIdentityEdgeRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this get current identity edge routers unauthorized response has a 2xx status code
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current identity edge routers unauthorized response has a 3xx status code
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current identity edge routers unauthorized response has a 4xx status code
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current identity edge routers unauthorized response has a 5xx status code
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get current identity edge routers unauthorized response a status code equal to that given
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get current identity edge routers unauthorized response
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) Code() int {
	return 401
}

func (o *GetCurrentIdentityEdgeRoutersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/edge-routers][%d] getCurrentIdentityEdgeRoutersUnauthorized %s", 401, payload)
}

func (o *GetCurrentIdentityEdgeRoutersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/edge-routers][%d] getCurrentIdentityEdgeRoutersUnauthorized %s", 401, payload)
}

func (o *GetCurrentIdentityEdgeRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetCurrentIdentityEdgeRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentIdentityEdgeRoutersTooManyRequests creates a GetCurrentIdentityEdgeRoutersTooManyRequests with default headers values
func NewGetCurrentIdentityEdgeRoutersTooManyRequests() *GetCurrentIdentityEdgeRoutersTooManyRequests {
	return &GetCurrentIdentityEdgeRoutersTooManyRequests{}
}

/*
GetCurrentIdentityEdgeRoutersTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type GetCurrentIdentityEdgeRoutersTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this get current identity edge routers too many requests response has a 2xx status code
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current identity edge routers too many requests response has a 3xx status code
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current identity edge routers too many requests response has a 4xx status code
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current identity edge routers too many requests response has a 5xx status code
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get current identity edge routers too many requests response a status code equal to that given
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get current identity edge routers too many requests response
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) Code() int {
	return 429
}

func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/edge-routers][%d] getCurrentIdentityEdgeRoutersTooManyRequests %s", 429, payload)
}

func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/edge-routers][%d] getCurrentIdentityEdgeRoutersTooManyRequests %s", 429, payload)
}

func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
