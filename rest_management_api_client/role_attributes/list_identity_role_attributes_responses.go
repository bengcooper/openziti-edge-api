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

package role_attributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListIdentityRoleAttributesReader is a Reader for the ListIdentityRoleAttributes structure.
type ListIdentityRoleAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIdentityRoleAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIdentityRoleAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListIdentityRoleAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListIdentityRoleAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListIdentityRoleAttributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListIdentityRoleAttributesOK creates a ListIdentityRoleAttributesOK with default headers values
func NewListIdentityRoleAttributesOK() *ListIdentityRoleAttributesOK {
	return &ListIdentityRoleAttributesOK{}
}

/* ListIdentityRoleAttributesOK describes a response with status code 200, with default header values.

A list of role attributes
*/
type ListIdentityRoleAttributesOK struct {
	Payload *rest_model.ListRoleAttributesEnvelope
}

func (o *ListIdentityRoleAttributesOK) Error() string {
	return fmt.Sprintf("[GET /identity-role-attributes][%d] listIdentityRoleAttributesOK  %+v", 200, o.Payload)
}
func (o *ListIdentityRoleAttributesOK) GetPayload() *rest_model.ListRoleAttributesEnvelope {
	return o.Payload
}

func (o *ListIdentityRoleAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListRoleAttributesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityRoleAttributesBadRequest creates a ListIdentityRoleAttributesBadRequest with default headers values
func NewListIdentityRoleAttributesBadRequest() *ListIdentityRoleAttributesBadRequest {
	return &ListIdentityRoleAttributesBadRequest{}
}

/* ListIdentityRoleAttributesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListIdentityRoleAttributesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityRoleAttributesBadRequest) Error() string {
	return fmt.Sprintf("[GET /identity-role-attributes][%d] listIdentityRoleAttributesBadRequest  %+v", 400, o.Payload)
}
func (o *ListIdentityRoleAttributesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityRoleAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityRoleAttributesUnauthorized creates a ListIdentityRoleAttributesUnauthorized with default headers values
func NewListIdentityRoleAttributesUnauthorized() *ListIdentityRoleAttributesUnauthorized {
	return &ListIdentityRoleAttributesUnauthorized{}
}

/* ListIdentityRoleAttributesUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListIdentityRoleAttributesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityRoleAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identity-role-attributes][%d] listIdentityRoleAttributesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListIdentityRoleAttributesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityRoleAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentityRoleAttributesTooManyRequests creates a ListIdentityRoleAttributesTooManyRequests with default headers values
func NewListIdentityRoleAttributesTooManyRequests() *ListIdentityRoleAttributesTooManyRequests {
	return &ListIdentityRoleAttributesTooManyRequests{}
}

/* ListIdentityRoleAttributesTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListIdentityRoleAttributesTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListIdentityRoleAttributesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /identity-role-attributes][%d] listIdentityRoleAttributesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListIdentityRoleAttributesTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListIdentityRoleAttributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
