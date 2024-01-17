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

package controllers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListControllersOKCode is the HTTP code returned for type ListControllersOK
const ListControllersOKCode int = 200

/*ListControllersOK A list of controllers

swagger:response listControllersOK
*/
type ListControllersOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListControllersEnvelope `json:"body,omitempty"`
}

// NewListControllersOK creates ListControllersOK with default headers values
func NewListControllersOK() *ListControllersOK {

	return &ListControllersOK{}
}

// WithPayload adds the payload to the list controllers o k response
func (o *ListControllersOK) WithPayload(payload *rest_model.ListControllersEnvelope) *ListControllersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list controllers o k response
func (o *ListControllersOK) SetPayload(payload *rest_model.ListControllersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListControllersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListControllersBadRequestCode is the HTTP code returned for type ListControllersBadRequest
const ListControllersBadRequestCode int = 400

/*ListControllersBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listControllersBadRequest
*/
type ListControllersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListControllersBadRequest creates ListControllersBadRequest with default headers values
func NewListControllersBadRequest() *ListControllersBadRequest {

	return &ListControllersBadRequest{}
}

// WithPayload adds the payload to the list controllers bad request response
func (o *ListControllersBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListControllersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list controllers bad request response
func (o *ListControllersBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListControllersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListControllersUnauthorizedCode is the HTTP code returned for type ListControllersUnauthorized
const ListControllersUnauthorizedCode int = 401

/*ListControllersUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listControllersUnauthorized
*/
type ListControllersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListControllersUnauthorized creates ListControllersUnauthorized with default headers values
func NewListControllersUnauthorized() *ListControllersUnauthorized {

	return &ListControllersUnauthorized{}
}

// WithPayload adds the payload to the list controllers unauthorized response
func (o *ListControllersUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListControllersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list controllers unauthorized response
func (o *ListControllersUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListControllersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListControllersTooManyRequestsCode is the HTTP code returned for type ListControllersTooManyRequests
const ListControllersTooManyRequestsCode int = 429

/*ListControllersTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listControllersTooManyRequests
*/
type ListControllersTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListControllersTooManyRequests creates ListControllersTooManyRequests with default headers values
func NewListControllersTooManyRequests() *ListControllersTooManyRequests {

	return &ListControllersTooManyRequests{}
}

// WithPayload adds the payload to the list controllers too many requests response
func (o *ListControllersTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListControllersTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list controllers too many requests response
func (o *ListControllersTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListControllersTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
