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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListConfigTypesOKCode is the HTTP code returned for type ListConfigTypesOK
const ListConfigTypesOKCode int = 200

/*
ListConfigTypesOK A list of config-types

swagger:response listConfigTypesOK
*/
type ListConfigTypesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListConfigTypesEnvelope `json:"body,omitempty"`
}

// NewListConfigTypesOK creates ListConfigTypesOK with default headers values
func NewListConfigTypesOK() *ListConfigTypesOK {

	return &ListConfigTypesOK{}
}

// WithPayload adds the payload to the list config types o k response
func (o *ListConfigTypesOK) WithPayload(payload *rest_model.ListConfigTypesEnvelope) *ListConfigTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config types o k response
func (o *ListConfigTypesOK) SetPayload(payload *rest_model.ListConfigTypesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigTypesBadRequestCode is the HTTP code returned for type ListConfigTypesBadRequest
const ListConfigTypesBadRequestCode int = 400

/*
ListConfigTypesBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listConfigTypesBadRequest
*/
type ListConfigTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigTypesBadRequest creates ListConfigTypesBadRequest with default headers values
func NewListConfigTypesBadRequest() *ListConfigTypesBadRequest {

	return &ListConfigTypesBadRequest{}
}

// WithPayload adds the payload to the list config types bad request response
func (o *ListConfigTypesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config types bad request response
func (o *ListConfigTypesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigTypesUnauthorizedCode is the HTTP code returned for type ListConfigTypesUnauthorized
const ListConfigTypesUnauthorizedCode int = 401

/*
ListConfigTypesUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listConfigTypesUnauthorized
*/
type ListConfigTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigTypesUnauthorized creates ListConfigTypesUnauthorized with default headers values
func NewListConfigTypesUnauthorized() *ListConfigTypesUnauthorized {

	return &ListConfigTypesUnauthorized{}
}

// WithPayload adds the payload to the list config types unauthorized response
func (o *ListConfigTypesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config types unauthorized response
func (o *ListConfigTypesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigTypesTooManyRequestsCode is the HTTP code returned for type ListConfigTypesTooManyRequests
const ListConfigTypesTooManyRequestsCode int = 429

/*
ListConfigTypesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listConfigTypesTooManyRequests
*/
type ListConfigTypesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigTypesTooManyRequests creates ListConfigTypesTooManyRequests with default headers values
func NewListConfigTypesTooManyRequests() *ListConfigTypesTooManyRequests {

	return &ListConfigTypesTooManyRequests{}
}

// WithPayload adds the payload to the list config types too many requests response
func (o *ListConfigTypesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigTypesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config types too many requests response
func (o *ListConfigTypesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigTypesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
