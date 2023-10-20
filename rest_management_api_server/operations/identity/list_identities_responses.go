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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListIdentitiesOKCode is the HTTP code returned for type ListIdentitiesOK
const ListIdentitiesOKCode int = 200

/*ListIdentitiesOK A list of identities

swagger:response listIdentitiesOK
*/
type ListIdentitiesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListIdentitiesEnvelope `json:"body,omitempty"`
}

// NewListIdentitiesOK creates ListIdentitiesOK with default headers values
func NewListIdentitiesOK() *ListIdentitiesOK {

	return &ListIdentitiesOK{}
}

// WithPayload adds the payload to the list identities o k response
func (o *ListIdentitiesOK) WithPayload(payload *rest_model.ListIdentitiesEnvelope) *ListIdentitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identities o k response
func (o *ListIdentitiesOK) SetPayload(payload *rest_model.ListIdentitiesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitiesBadRequestCode is the HTTP code returned for type ListIdentitiesBadRequest
const ListIdentitiesBadRequestCode int = 400

/*ListIdentitiesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listIdentitiesBadRequest
*/
type ListIdentitiesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitiesBadRequest creates ListIdentitiesBadRequest with default headers values
func NewListIdentitiesBadRequest() *ListIdentitiesBadRequest {

	return &ListIdentitiesBadRequest{}
}

// WithPayload adds the payload to the list identities bad request response
func (o *ListIdentitiesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitiesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identities bad request response
func (o *ListIdentitiesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitiesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitiesUnauthorizedCode is the HTTP code returned for type ListIdentitiesUnauthorized
const ListIdentitiesUnauthorizedCode int = 401

/*ListIdentitiesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listIdentitiesUnauthorized
*/
type ListIdentitiesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitiesUnauthorized creates ListIdentitiesUnauthorized with default headers values
func NewListIdentitiesUnauthorized() *ListIdentitiesUnauthorized {

	return &ListIdentitiesUnauthorized{}
}

// WithPayload adds the payload to the list identities unauthorized response
func (o *ListIdentitiesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitiesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identities unauthorized response
func (o *ListIdentitiesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitiesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentitiesTooManyRequestsCode is the HTTP code returned for type ListIdentitiesTooManyRequests
const ListIdentitiesTooManyRequestsCode int = 429

/*ListIdentitiesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listIdentitiesTooManyRequests
*/
type ListIdentitiesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentitiesTooManyRequests creates ListIdentitiesTooManyRequests with default headers values
func NewListIdentitiesTooManyRequests() *ListIdentitiesTooManyRequests {

	return &ListIdentitiesTooManyRequests{}
}

// WithPayload adds the payload to the list identities too many requests response
func (o *ListIdentitiesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentitiesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identities too many requests response
func (o *ListIdentitiesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentitiesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
