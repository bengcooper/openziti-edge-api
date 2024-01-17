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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchCaOKCode is the HTTP code returned for type PatchCaOK
const PatchCaOKCode int = 200

/*PatchCaOK The patch request was successful and the resource has been altered

swagger:response patchCaOK
*/
type PatchCaOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchCaOK creates PatchCaOK with default headers values
func NewPatchCaOK() *PatchCaOK {

	return &PatchCaOK{}
}

// WithPayload adds the payload to the patch ca o k response
func (o *PatchCaOK) WithPayload(payload *rest_model.Empty) *PatchCaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch ca o k response
func (o *PatchCaOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCaBadRequestCode is the HTTP code returned for type PatchCaBadRequest
const PatchCaBadRequestCode int = 400

/*PatchCaBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchCaBadRequest
*/
type PatchCaBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCaBadRequest creates PatchCaBadRequest with default headers values
func NewPatchCaBadRequest() *PatchCaBadRequest {

	return &PatchCaBadRequest{}
}

// WithPayload adds the payload to the patch ca bad request response
func (o *PatchCaBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCaBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch ca bad request response
func (o *PatchCaBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCaUnauthorizedCode is the HTTP code returned for type PatchCaUnauthorized
const PatchCaUnauthorizedCode int = 401

/*PatchCaUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchCaUnauthorized
*/
type PatchCaUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCaUnauthorized creates PatchCaUnauthorized with default headers values
func NewPatchCaUnauthorized() *PatchCaUnauthorized {

	return &PatchCaUnauthorized{}
}

// WithPayload adds the payload to the patch ca unauthorized response
func (o *PatchCaUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCaUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch ca unauthorized response
func (o *PatchCaUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCaNotFoundCode is the HTTP code returned for type PatchCaNotFound
const PatchCaNotFoundCode int = 404

/*PatchCaNotFound The requested resource does not exist

swagger:response patchCaNotFound
*/
type PatchCaNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCaNotFound creates PatchCaNotFound with default headers values
func NewPatchCaNotFound() *PatchCaNotFound {

	return &PatchCaNotFound{}
}

// WithPayload adds the payload to the patch ca not found response
func (o *PatchCaNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCaNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch ca not found response
func (o *PatchCaNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCaNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCaTooManyRequestsCode is the HTTP code returned for type PatchCaTooManyRequests
const PatchCaTooManyRequestsCode int = 429

/*PatchCaTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response patchCaTooManyRequests
*/
type PatchCaTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCaTooManyRequests creates PatchCaTooManyRequests with default headers values
func NewPatchCaTooManyRequests() *PatchCaTooManyRequests {

	return &PatchCaTooManyRequests{}
}

// WithPayload adds the payload to the patch ca too many requests response
func (o *PatchCaTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCaTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch ca too many requests response
func (o *PatchCaTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCaTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
