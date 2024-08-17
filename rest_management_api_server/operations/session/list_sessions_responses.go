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

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListSessionsOKCode is the HTTP code returned for type ListSessionsOK
const ListSessionsOKCode int = 200

/*
ListSessionsOK A list of sessions

swagger:response listSessionsOK
*/
type ListSessionsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListSessionsManagementEnvelope `json:"body,omitempty"`
}

// NewListSessionsOK creates ListSessionsOK with default headers values
func NewListSessionsOK() *ListSessionsOK {

	return &ListSessionsOK{}
}

// WithPayload adds the payload to the list sessions o k response
func (o *ListSessionsOK) WithPayload(payload *rest_model.ListSessionsManagementEnvelope) *ListSessionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list sessions o k response
func (o *ListSessionsOK) SetPayload(payload *rest_model.ListSessionsManagementEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSessionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSessionsBadRequestCode is the HTTP code returned for type ListSessionsBadRequest
const ListSessionsBadRequestCode int = 400

/*
ListSessionsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listSessionsBadRequest
*/
type ListSessionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListSessionsBadRequest creates ListSessionsBadRequest with default headers values
func NewListSessionsBadRequest() *ListSessionsBadRequest {

	return &ListSessionsBadRequest{}
}

// WithPayload adds the payload to the list sessions bad request response
func (o *ListSessionsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListSessionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list sessions bad request response
func (o *ListSessionsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSessionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSessionsUnauthorizedCode is the HTTP code returned for type ListSessionsUnauthorized
const ListSessionsUnauthorizedCode int = 401

/*
ListSessionsUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listSessionsUnauthorized
*/
type ListSessionsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListSessionsUnauthorized creates ListSessionsUnauthorized with default headers values
func NewListSessionsUnauthorized() *ListSessionsUnauthorized {

	return &ListSessionsUnauthorized{}
}

// WithPayload adds the payload to the list sessions unauthorized response
func (o *ListSessionsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListSessionsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list sessions unauthorized response
func (o *ListSessionsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSessionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSessionsTooManyRequestsCode is the HTTP code returned for type ListSessionsTooManyRequests
const ListSessionsTooManyRequestsCode int = 429

/*
ListSessionsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listSessionsTooManyRequests
*/
type ListSessionsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListSessionsTooManyRequests creates ListSessionsTooManyRequests with default headers values
func NewListSessionsTooManyRequests() *ListSessionsTooManyRequests {

	return &ListSessionsTooManyRequests{}
}

// WithPayload adds the payload to the list sessions too many requests response
func (o *ListSessionsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListSessionsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list sessions too many requests response
func (o *ListSessionsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSessionsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
