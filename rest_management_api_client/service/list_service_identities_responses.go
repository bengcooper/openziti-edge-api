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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListServiceIdentitiesReader is a Reader for the ListServiceIdentities structure.
type ListServiceIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServiceIdentitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListServiceIdentitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServiceIdentitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceIdentitiesOK creates a ListServiceIdentitiesOK with default headers values
func NewListServiceIdentitiesOK() *ListServiceIdentitiesOK {
	return &ListServiceIdentitiesOK{}
}

/* ListServiceIdentitiesOK describes a response with status code 200, with default header values.

A list of identities
*/
type ListServiceIdentitiesOK struct {
	Payload *rest_model.ListIdentitiesEnvelope
}

func (o *ListServiceIdentitiesOK) Error() string {
	return fmt.Sprintf("[GET /services/{id}/identities][%d] listServiceIdentitiesOK  %+v", 200, o.Payload)
}
func (o *ListServiceIdentitiesOK) GetPayload() *rest_model.ListIdentitiesEnvelope {
	return o.Payload
}

func (o *ListServiceIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListIdentitiesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceIdentitiesBadRequest creates a ListServiceIdentitiesBadRequest with default headers values
func NewListServiceIdentitiesBadRequest() *ListServiceIdentitiesBadRequest {
	return &ListServiceIdentitiesBadRequest{}
}

/* ListServiceIdentitiesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListServiceIdentitiesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceIdentitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /services/{id}/identities][%d] listServiceIdentitiesBadRequest  %+v", 400, o.Payload)
}
func (o *ListServiceIdentitiesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceIdentitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceIdentitiesUnauthorized creates a ListServiceIdentitiesUnauthorized with default headers values
func NewListServiceIdentitiesUnauthorized() *ListServiceIdentitiesUnauthorized {
	return &ListServiceIdentitiesUnauthorized{}
}

/* ListServiceIdentitiesUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListServiceIdentitiesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceIdentitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services/{id}/identities][%d] listServiceIdentitiesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServiceIdentitiesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceIdentitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceIdentitiesTooManyRequests creates a ListServiceIdentitiesTooManyRequests with default headers values
func NewListServiceIdentitiesTooManyRequests() *ListServiceIdentitiesTooManyRequests {
	return &ListServiceIdentitiesTooManyRequests{}
}

/* ListServiceIdentitiesTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListServiceIdentitiesTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceIdentitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /services/{id}/identities][%d] listServiceIdentitiesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServiceIdentitiesTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceIdentitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
