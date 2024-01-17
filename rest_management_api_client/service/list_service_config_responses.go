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

// ListServiceConfigReader is a Reader for the ListServiceConfig structure.
type ListServiceConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListServiceConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListServiceConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListServiceConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceConfigOK creates a ListServiceConfigOK with default headers values
func NewListServiceConfigOK() *ListServiceConfigOK {
	return &ListServiceConfigOK{}
}

/* ListServiceConfigOK describes a response with status code 200, with default header values.

A list of configs
*/
type ListServiceConfigOK struct {
	Payload *rest_model.ListConfigsEnvelope
}

func (o *ListServiceConfigOK) Error() string {
	return fmt.Sprintf("[GET /services/{id}/configs][%d] listServiceConfigOK  %+v", 200, o.Payload)
}
func (o *ListServiceConfigOK) GetPayload() *rest_model.ListConfigsEnvelope {
	return o.Payload
}

func (o *ListServiceConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListConfigsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceConfigBadRequest creates a ListServiceConfigBadRequest with default headers values
func NewListServiceConfigBadRequest() *ListServiceConfigBadRequest {
	return &ListServiceConfigBadRequest{}
}

/* ListServiceConfigBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListServiceConfigBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /services/{id}/configs][%d] listServiceConfigBadRequest  %+v", 400, o.Payload)
}
func (o *ListServiceConfigBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceConfigUnauthorized creates a ListServiceConfigUnauthorized with default headers values
func NewListServiceConfigUnauthorized() *ListServiceConfigUnauthorized {
	return &ListServiceConfigUnauthorized{}
}

/* ListServiceConfigUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListServiceConfigUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services/{id}/configs][%d] listServiceConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServiceConfigUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceConfigTooManyRequests creates a ListServiceConfigTooManyRequests with default headers values
func NewListServiceConfigTooManyRequests() *ListServiceConfigTooManyRequests {
	return &ListServiceConfigTooManyRequests{}
}

/* ListServiceConfigTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListServiceConfigTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /services/{id}/configs][%d] listServiceConfigTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListServiceConfigTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
