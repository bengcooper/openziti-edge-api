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

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchTransitRouterOKCode is the HTTP code returned for type PatchTransitRouterOK
const PatchTransitRouterOKCode int = 200

/*PatchTransitRouterOK The patch request was successful and the resource has been altered

swagger:response patchTransitRouterOK
*/
type PatchTransitRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchTransitRouterOK creates PatchTransitRouterOK with default headers values
func NewPatchTransitRouterOK() *PatchTransitRouterOK {

	return &PatchTransitRouterOK{}
}

// WithPayload adds the payload to the patch transit router o k response
func (o *PatchTransitRouterOK) WithPayload(payload *rest_model.Empty) *PatchTransitRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch transit router o k response
func (o *PatchTransitRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTransitRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTransitRouterBadRequestCode is the HTTP code returned for type PatchTransitRouterBadRequest
const PatchTransitRouterBadRequestCode int = 400

/*PatchTransitRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchTransitRouterBadRequest
*/
type PatchTransitRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchTransitRouterBadRequest creates PatchTransitRouterBadRequest with default headers values
func NewPatchTransitRouterBadRequest() *PatchTransitRouterBadRequest {

	return &PatchTransitRouterBadRequest{}
}

// WithPayload adds the payload to the patch transit router bad request response
func (o *PatchTransitRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchTransitRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch transit router bad request response
func (o *PatchTransitRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTransitRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTransitRouterUnauthorizedCode is the HTTP code returned for type PatchTransitRouterUnauthorized
const PatchTransitRouterUnauthorizedCode int = 401

/*PatchTransitRouterUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchTransitRouterUnauthorized
*/
type PatchTransitRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchTransitRouterUnauthorized creates PatchTransitRouterUnauthorized with default headers values
func NewPatchTransitRouterUnauthorized() *PatchTransitRouterUnauthorized {

	return &PatchTransitRouterUnauthorized{}
}

// WithPayload adds the payload to the patch transit router unauthorized response
func (o *PatchTransitRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchTransitRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch transit router unauthorized response
func (o *PatchTransitRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTransitRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTransitRouterNotFoundCode is the HTTP code returned for type PatchTransitRouterNotFound
const PatchTransitRouterNotFoundCode int = 404

/*PatchTransitRouterNotFound The requested resource does not exist

swagger:response patchTransitRouterNotFound
*/
type PatchTransitRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchTransitRouterNotFound creates PatchTransitRouterNotFound with default headers values
func NewPatchTransitRouterNotFound() *PatchTransitRouterNotFound {

	return &PatchTransitRouterNotFound{}
}

// WithPayload adds the payload to the patch transit router not found response
func (o *PatchTransitRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchTransitRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch transit router not found response
func (o *PatchTransitRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTransitRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTransitRouterTooManyRequestsCode is the HTTP code returned for type PatchTransitRouterTooManyRequests
const PatchTransitRouterTooManyRequestsCode int = 429

/*PatchTransitRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchTransitRouterTooManyRequests
*/
type PatchTransitRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchTransitRouterTooManyRequests creates PatchTransitRouterTooManyRequests with default headers values
func NewPatchTransitRouterTooManyRequests() *PatchTransitRouterTooManyRequests {

	return &PatchTransitRouterTooManyRequests{}
}

// WithPayload adds the payload to the patch transit router too many requests response
func (o *PatchTransitRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchTransitRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch transit router too many requests response
func (o *PatchTransitRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTransitRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
