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

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListEnrollmentsOKCode is the HTTP code returned for type ListEnrollmentsOK
const ListEnrollmentsOKCode int = 200

/*ListEnrollmentsOK A list of enrollments

swagger:response listEnrollmentsOK
*/
type ListEnrollmentsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListEnrollmentsEnvelope `json:"body,omitempty"`
}

// NewListEnrollmentsOK creates ListEnrollmentsOK with default headers values
func NewListEnrollmentsOK() *ListEnrollmentsOK {

	return &ListEnrollmentsOK{}
}

// WithPayload adds the payload to the list enrollments o k response
func (o *ListEnrollmentsOK) WithPayload(payload *rest_model.ListEnrollmentsEnvelope) *ListEnrollmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list enrollments o k response
func (o *ListEnrollmentsOK) SetPayload(payload *rest_model.ListEnrollmentsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEnrollmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEnrollmentsBadRequestCode is the HTTP code returned for type ListEnrollmentsBadRequest
const ListEnrollmentsBadRequestCode int = 400

/*ListEnrollmentsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listEnrollmentsBadRequest
*/
type ListEnrollmentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEnrollmentsBadRequest creates ListEnrollmentsBadRequest with default headers values
func NewListEnrollmentsBadRequest() *ListEnrollmentsBadRequest {

	return &ListEnrollmentsBadRequest{}
}

// WithPayload adds the payload to the list enrollments bad request response
func (o *ListEnrollmentsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEnrollmentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list enrollments bad request response
func (o *ListEnrollmentsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEnrollmentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEnrollmentsUnauthorizedCode is the HTTP code returned for type ListEnrollmentsUnauthorized
const ListEnrollmentsUnauthorizedCode int = 401

/*ListEnrollmentsUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listEnrollmentsUnauthorized
*/
type ListEnrollmentsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEnrollmentsUnauthorized creates ListEnrollmentsUnauthorized with default headers values
func NewListEnrollmentsUnauthorized() *ListEnrollmentsUnauthorized {

	return &ListEnrollmentsUnauthorized{}
}

// WithPayload adds the payload to the list enrollments unauthorized response
func (o *ListEnrollmentsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEnrollmentsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list enrollments unauthorized response
func (o *ListEnrollmentsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEnrollmentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEnrollmentsTooManyRequestsCode is the HTTP code returned for type ListEnrollmentsTooManyRequests
const ListEnrollmentsTooManyRequestsCode int = 429

/*ListEnrollmentsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listEnrollmentsTooManyRequests
*/
type ListEnrollmentsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEnrollmentsTooManyRequests creates ListEnrollmentsTooManyRequests with default headers values
func NewListEnrollmentsTooManyRequests() *ListEnrollmentsTooManyRequests {

	return &ListEnrollmentsTooManyRequests{}
}

// WithPayload adds the payload to the list enrollments too many requests response
func (o *ListEnrollmentsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEnrollmentsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list enrollments too many requests response
func (o *ListEnrollmentsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEnrollmentsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
