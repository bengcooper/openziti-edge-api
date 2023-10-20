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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListPostureChecksReader is a Reader for the ListPostureChecks structure.
type ListPostureChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPostureChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPostureChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPostureChecksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListPostureChecksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListPostureChecksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPostureChecksOK creates a ListPostureChecksOK with default headers values
func NewListPostureChecksOK() *ListPostureChecksOK {
	return &ListPostureChecksOK{}
}

/* ListPostureChecksOK describes a response with status code 200, with default header values.

A list of posture checks
*/
type ListPostureChecksOK struct {
	Payload *rest_model.ListPostureCheckEnvelope
}

func (o *ListPostureChecksOK) Error() string {
	return fmt.Sprintf("[GET /posture-checks][%d] listPostureChecksOK  %+v", 200, o.Payload)
}
func (o *ListPostureChecksOK) GetPayload() *rest_model.ListPostureCheckEnvelope {
	return o.Payload
}

func (o *ListPostureChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListPostureCheckEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostureChecksBadRequest creates a ListPostureChecksBadRequest with default headers values
func NewListPostureChecksBadRequest() *ListPostureChecksBadRequest {
	return &ListPostureChecksBadRequest{}
}

/* ListPostureChecksBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListPostureChecksBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListPostureChecksBadRequest) Error() string {
	return fmt.Sprintf("[GET /posture-checks][%d] listPostureChecksBadRequest  %+v", 400, o.Payload)
}
func (o *ListPostureChecksBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListPostureChecksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostureChecksUnauthorized creates a ListPostureChecksUnauthorized with default headers values
func NewListPostureChecksUnauthorized() *ListPostureChecksUnauthorized {
	return &ListPostureChecksUnauthorized{}
}

/* ListPostureChecksUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListPostureChecksUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListPostureChecksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /posture-checks][%d] listPostureChecksUnauthorized  %+v", 401, o.Payload)
}
func (o *ListPostureChecksUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListPostureChecksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPostureChecksTooManyRequests creates a ListPostureChecksTooManyRequests with default headers values
func NewListPostureChecksTooManyRequests() *ListPostureChecksTooManyRequests {
	return &ListPostureChecksTooManyRequests{}
}

/* ListPostureChecksTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListPostureChecksTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListPostureChecksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /posture-checks][%d] listPostureChecksTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListPostureChecksTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListPostureChecksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
