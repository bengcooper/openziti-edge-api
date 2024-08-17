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

package current_api_session

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

// ListCurrentIdentityAuthenticatorsReader is a Reader for the ListCurrentIdentityAuthenticators structure.
type ListCurrentIdentityAuthenticatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCurrentIdentityAuthenticatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCurrentIdentityAuthenticatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCurrentIdentityAuthenticatorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCurrentIdentityAuthenticatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /current-identity/authenticators] listCurrentIdentityAuthenticators", response, response.Code())
	}
}

// NewListCurrentIdentityAuthenticatorsOK creates a ListCurrentIdentityAuthenticatorsOK with default headers values
func NewListCurrentIdentityAuthenticatorsOK() *ListCurrentIdentityAuthenticatorsOK {
	return &ListCurrentIdentityAuthenticatorsOK{}
}

/*
ListCurrentIdentityAuthenticatorsOK describes a response with status code 200, with default header values.

A list of authenticators
*/
type ListCurrentIdentityAuthenticatorsOK struct {
	Payload *rest_model.ListAuthenticatorsEnvelope
}

// IsSuccess returns true when this list current identity authenticators o k response has a 2xx status code
func (o *ListCurrentIdentityAuthenticatorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list current identity authenticators o k response has a 3xx status code
func (o *ListCurrentIdentityAuthenticatorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list current identity authenticators o k response has a 4xx status code
func (o *ListCurrentIdentityAuthenticatorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list current identity authenticators o k response has a 5xx status code
func (o *ListCurrentIdentityAuthenticatorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list current identity authenticators o k response a status code equal to that given
func (o *ListCurrentIdentityAuthenticatorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list current identity authenticators o k response
func (o *ListCurrentIdentityAuthenticatorsOK) Code() int {
	return 200
}

func (o *ListCurrentIdentityAuthenticatorsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/authenticators][%d] listCurrentIdentityAuthenticatorsOK %s", 200, payload)
}

func (o *ListCurrentIdentityAuthenticatorsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/authenticators][%d] listCurrentIdentityAuthenticatorsOK %s", 200, payload)
}

func (o *ListCurrentIdentityAuthenticatorsOK) GetPayload() *rest_model.ListAuthenticatorsEnvelope {
	return o.Payload
}

func (o *ListCurrentIdentityAuthenticatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListAuthenticatorsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCurrentIdentityAuthenticatorsBadRequest creates a ListCurrentIdentityAuthenticatorsBadRequest with default headers values
func NewListCurrentIdentityAuthenticatorsBadRequest() *ListCurrentIdentityAuthenticatorsBadRequest {
	return &ListCurrentIdentityAuthenticatorsBadRequest{}
}

/*
ListCurrentIdentityAuthenticatorsBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListCurrentIdentityAuthenticatorsBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this list current identity authenticators bad request response has a 2xx status code
func (o *ListCurrentIdentityAuthenticatorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list current identity authenticators bad request response has a 3xx status code
func (o *ListCurrentIdentityAuthenticatorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list current identity authenticators bad request response has a 4xx status code
func (o *ListCurrentIdentityAuthenticatorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list current identity authenticators bad request response has a 5xx status code
func (o *ListCurrentIdentityAuthenticatorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list current identity authenticators bad request response a status code equal to that given
func (o *ListCurrentIdentityAuthenticatorsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list current identity authenticators bad request response
func (o *ListCurrentIdentityAuthenticatorsBadRequest) Code() int {
	return 400
}

func (o *ListCurrentIdentityAuthenticatorsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/authenticators][%d] listCurrentIdentityAuthenticatorsBadRequest %s", 400, payload)
}

func (o *ListCurrentIdentityAuthenticatorsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/authenticators][%d] listCurrentIdentityAuthenticatorsBadRequest %s", 400, payload)
}

func (o *ListCurrentIdentityAuthenticatorsBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCurrentIdentityAuthenticatorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCurrentIdentityAuthenticatorsUnauthorized creates a ListCurrentIdentityAuthenticatorsUnauthorized with default headers values
func NewListCurrentIdentityAuthenticatorsUnauthorized() *ListCurrentIdentityAuthenticatorsUnauthorized {
	return &ListCurrentIdentityAuthenticatorsUnauthorized{}
}

/*
ListCurrentIdentityAuthenticatorsUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListCurrentIdentityAuthenticatorsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this list current identity authenticators unauthorized response has a 2xx status code
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list current identity authenticators unauthorized response has a 3xx status code
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list current identity authenticators unauthorized response has a 4xx status code
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list current identity authenticators unauthorized response has a 5xx status code
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list current identity authenticators unauthorized response a status code equal to that given
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list current identity authenticators unauthorized response
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) Code() int {
	return 401
}

func (o *ListCurrentIdentityAuthenticatorsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/authenticators][%d] listCurrentIdentityAuthenticatorsUnauthorized %s", 401, payload)
}

func (o *ListCurrentIdentityAuthenticatorsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /current-identity/authenticators][%d] listCurrentIdentityAuthenticatorsUnauthorized %s", 401, payload)
}

func (o *ListCurrentIdentityAuthenticatorsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCurrentIdentityAuthenticatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
