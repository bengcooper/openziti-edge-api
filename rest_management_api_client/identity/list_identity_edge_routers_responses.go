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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListIdentityEdgeRoutersReader is a Reader for the ListIdentityEdgeRouters structure.
type ListIdentityEdgeRoutersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIdentityEdgeRoutersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIdentityEdgeRoutersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListIdentityEdgeRoutersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListIdentityEdgeRoutersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListIdentityEdgeRoutersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListIdentityEdgeRoutersOK creates a ListIdentityEdgeRoutersOK with default headers values
func NewListIdentityEdgeRoutersOK() *ListIdentityEdgeRoutersOK {
	return &ListIdentityEdgeRoutersOK{}
}

/* ListIdentityEdgeRoutersOK describes a response with status code 200, with default header values.

A list of edge routers
*/
type ListIdentityEdgeRoutersOK struct {
	Payload *rest_model.ListEdgeRoutersEnvelope
}

func (o *ListIdentityEdgeRoutersOK) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/edge-routers][%d] listIdentityEdgeRoutersOK  %+v", 200, o.Payload)
}
func (o *ListIdentityEdgeRoutersOK) GetPayload() *rest_model.ListEdgeRoutersEnvelope {
	return o.Payload
}

func (o *ListIdentityEdgeRoutersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRoutersEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityEdgeRoutersUnauthorized creates a ListIdentityEdgeRoutersUnauthorized with default headers values
func NewListIdentityEdgeRoutersUnauthorized() *ListIdentityEdgeRoutersUnauthorized {
	return &ListIdentityEdgeRoutersUnauthorized{}
}

/* ListIdentityEdgeRoutersUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListIdentityEdgeRoutersUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityEdgeRoutersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/edge-routers][%d] listIdentityEdgeRoutersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListIdentityEdgeRoutersUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityEdgeRoutersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityEdgeRoutersNotFound creates a ListIdentityEdgeRoutersNotFound with default headers values
func NewListIdentityEdgeRoutersNotFound() *ListIdentityEdgeRoutersNotFound {
	return &ListIdentityEdgeRoutersNotFound{}
}

/* ListIdentityEdgeRoutersNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type ListIdentityEdgeRoutersNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityEdgeRoutersNotFound) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/edge-routers][%d] listIdentityEdgeRoutersNotFound  %+v", 404, o.Payload)
}
func (o *ListIdentityEdgeRoutersNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityEdgeRoutersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityEdgeRoutersTooManyRequests creates a ListIdentityEdgeRoutersTooManyRequests with default headers values
func NewListIdentityEdgeRoutersTooManyRequests() *ListIdentityEdgeRoutersTooManyRequests {
	return &ListIdentityEdgeRoutersTooManyRequests{}
}

/* ListIdentityEdgeRoutersTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListIdentityEdgeRoutersTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityEdgeRoutersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/edge-routers][%d] listIdentityEdgeRoutersTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListIdentityEdgeRoutersTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityEdgeRoutersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
