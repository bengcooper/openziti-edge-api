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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchEdgeRouterOKCode is the HTTP code returned for type PatchEdgeRouterOK
const PatchEdgeRouterOKCode int = 200

/*PatchEdgeRouterOK The patch request was successful and the resource has been altered

swagger:response patchEdgeRouterOK
*/
type PatchEdgeRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchEdgeRouterOK creates PatchEdgeRouterOK with default headers values
func NewPatchEdgeRouterOK() *PatchEdgeRouterOK {

	return &PatchEdgeRouterOK{}
}

// WithPayload adds the payload to the patch edge router o k response
func (o *PatchEdgeRouterOK) WithPayload(payload *rest_model.Empty) *PatchEdgeRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch edge router o k response
func (o *PatchEdgeRouterOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEdgeRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchEdgeRouterBadRequestCode is the HTTP code returned for type PatchEdgeRouterBadRequest
const PatchEdgeRouterBadRequestCode int = 400

/*PatchEdgeRouterBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchEdgeRouterBadRequest
*/
type PatchEdgeRouterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchEdgeRouterBadRequest creates PatchEdgeRouterBadRequest with default headers values
func NewPatchEdgeRouterBadRequest() *PatchEdgeRouterBadRequest {

	return &PatchEdgeRouterBadRequest{}
}

// WithPayload adds the payload to the patch edge router bad request response
func (o *PatchEdgeRouterBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchEdgeRouterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch edge router bad request response
func (o *PatchEdgeRouterBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEdgeRouterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchEdgeRouterUnauthorizedCode is the HTTP code returned for type PatchEdgeRouterUnauthorized
const PatchEdgeRouterUnauthorizedCode int = 401

/*PatchEdgeRouterUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchEdgeRouterUnauthorized
*/
type PatchEdgeRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchEdgeRouterUnauthorized creates PatchEdgeRouterUnauthorized with default headers values
func NewPatchEdgeRouterUnauthorized() *PatchEdgeRouterUnauthorized {

	return &PatchEdgeRouterUnauthorized{}
}

// WithPayload adds the payload to the patch edge router unauthorized response
func (o *PatchEdgeRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchEdgeRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch edge router unauthorized response
func (o *PatchEdgeRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEdgeRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchEdgeRouterNotFoundCode is the HTTP code returned for type PatchEdgeRouterNotFound
const PatchEdgeRouterNotFoundCode int = 404

/*PatchEdgeRouterNotFound The requested resource does not exist

swagger:response patchEdgeRouterNotFound
*/
type PatchEdgeRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchEdgeRouterNotFound creates PatchEdgeRouterNotFound with default headers values
func NewPatchEdgeRouterNotFound() *PatchEdgeRouterNotFound {

	return &PatchEdgeRouterNotFound{}
}

// WithPayload adds the payload to the patch edge router not found response
func (o *PatchEdgeRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchEdgeRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch edge router not found response
func (o *PatchEdgeRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEdgeRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchEdgeRouterTooManyRequestsCode is the HTTP code returned for type PatchEdgeRouterTooManyRequests
const PatchEdgeRouterTooManyRequestsCode int = 429

/*PatchEdgeRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchEdgeRouterTooManyRequests
*/
type PatchEdgeRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchEdgeRouterTooManyRequests creates PatchEdgeRouterTooManyRequests with default headers values
func NewPatchEdgeRouterTooManyRequests() *PatchEdgeRouterTooManyRequests {

	return &PatchEdgeRouterTooManyRequests{}
}

// WithPayload adds the payload to the patch edge router too many requests response
func (o *PatchEdgeRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchEdgeRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch edge router too many requests response
func (o *PatchEdgeRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEdgeRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
