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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchServicePolicyOKCode is the HTTP code returned for type PatchServicePolicyOK
const PatchServicePolicyOKCode int = 200

/*PatchServicePolicyOK The patch request was successful and the resource has been altered

swagger:response patchServicePolicyOK
*/
type PatchServicePolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchServicePolicyOK creates PatchServicePolicyOK with default headers values
func NewPatchServicePolicyOK() *PatchServicePolicyOK {

	return &PatchServicePolicyOK{}
}

// WithPayload adds the payload to the patch service policy o k response
func (o *PatchServicePolicyOK) WithPayload(payload *rest_model.Empty) *PatchServicePolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service policy o k response
func (o *PatchServicePolicyOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServicePolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServicePolicyBadRequestCode is the HTTP code returned for type PatchServicePolicyBadRequest
const PatchServicePolicyBadRequestCode int = 400

/*PatchServicePolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchServicePolicyBadRequest
*/
type PatchServicePolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServicePolicyBadRequest creates PatchServicePolicyBadRequest with default headers values
func NewPatchServicePolicyBadRequest() *PatchServicePolicyBadRequest {

	return &PatchServicePolicyBadRequest{}
}

// WithPayload adds the payload to the patch service policy bad request response
func (o *PatchServicePolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServicePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service policy bad request response
func (o *PatchServicePolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServicePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServicePolicyUnauthorizedCode is the HTTP code returned for type PatchServicePolicyUnauthorized
const PatchServicePolicyUnauthorizedCode int = 401

/*PatchServicePolicyUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchServicePolicyUnauthorized
*/
type PatchServicePolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServicePolicyUnauthorized creates PatchServicePolicyUnauthorized with default headers values
func NewPatchServicePolicyUnauthorized() *PatchServicePolicyUnauthorized {

	return &PatchServicePolicyUnauthorized{}
}

// WithPayload adds the payload to the patch service policy unauthorized response
func (o *PatchServicePolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServicePolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service policy unauthorized response
func (o *PatchServicePolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServicePolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServicePolicyNotFoundCode is the HTTP code returned for type PatchServicePolicyNotFound
const PatchServicePolicyNotFoundCode int = 404

/*PatchServicePolicyNotFound The requested resource does not exist

swagger:response patchServicePolicyNotFound
*/
type PatchServicePolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServicePolicyNotFound creates PatchServicePolicyNotFound with default headers values
func NewPatchServicePolicyNotFound() *PatchServicePolicyNotFound {

	return &PatchServicePolicyNotFound{}
}

// WithPayload adds the payload to the patch service policy not found response
func (o *PatchServicePolicyNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServicePolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service policy not found response
func (o *PatchServicePolicyNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServicePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchServicePolicyTooManyRequestsCode is the HTTP code returned for type PatchServicePolicyTooManyRequests
const PatchServicePolicyTooManyRequestsCode int = 429

/*PatchServicePolicyTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchServicePolicyTooManyRequests
*/
type PatchServicePolicyTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchServicePolicyTooManyRequests creates PatchServicePolicyTooManyRequests with default headers values
func NewPatchServicePolicyTooManyRequests() *PatchServicePolicyTooManyRequests {

	return &PatchServicePolicyTooManyRequests{}
}

// WithPayload adds the payload to the patch service policy too many requests response
func (o *PatchServicePolicyTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchServicePolicyTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch service policy too many requests response
func (o *PatchServicePolicyTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchServicePolicyTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
