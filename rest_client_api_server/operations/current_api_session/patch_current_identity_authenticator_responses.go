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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// PatchCurrentIdentityAuthenticatorOKCode is the HTTP code returned for type PatchCurrentIdentityAuthenticatorOK
const PatchCurrentIdentityAuthenticatorOKCode int = 200

/*PatchCurrentIdentityAuthenticatorOK The patch request was successful and the resource has been altered

swagger:response patchCurrentIdentityAuthenticatorOK
*/
type PatchCurrentIdentityAuthenticatorOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewPatchCurrentIdentityAuthenticatorOK creates PatchCurrentIdentityAuthenticatorOK with default headers values
func NewPatchCurrentIdentityAuthenticatorOK() *PatchCurrentIdentityAuthenticatorOK {

	return &PatchCurrentIdentityAuthenticatorOK{}
}

// WithPayload adds the payload to the patch current identity authenticator o k response
func (o *PatchCurrentIdentityAuthenticatorOK) WithPayload(payload *rest_model.Empty) *PatchCurrentIdentityAuthenticatorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch current identity authenticator o k response
func (o *PatchCurrentIdentityAuthenticatorOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCurrentIdentityAuthenticatorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCurrentIdentityAuthenticatorBadRequestCode is the HTTP code returned for type PatchCurrentIdentityAuthenticatorBadRequest
const PatchCurrentIdentityAuthenticatorBadRequestCode int = 400

/*PatchCurrentIdentityAuthenticatorBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response patchCurrentIdentityAuthenticatorBadRequest
*/
type PatchCurrentIdentityAuthenticatorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCurrentIdentityAuthenticatorBadRequest creates PatchCurrentIdentityAuthenticatorBadRequest with default headers values
func NewPatchCurrentIdentityAuthenticatorBadRequest() *PatchCurrentIdentityAuthenticatorBadRequest {

	return &PatchCurrentIdentityAuthenticatorBadRequest{}
}

// WithPayload adds the payload to the patch current identity authenticator bad request response
func (o *PatchCurrentIdentityAuthenticatorBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCurrentIdentityAuthenticatorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch current identity authenticator bad request response
func (o *PatchCurrentIdentityAuthenticatorBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCurrentIdentityAuthenticatorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCurrentIdentityAuthenticatorUnauthorizedCode is the HTTP code returned for type PatchCurrentIdentityAuthenticatorUnauthorized
const PatchCurrentIdentityAuthenticatorUnauthorizedCode int = 401

/*PatchCurrentIdentityAuthenticatorUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response patchCurrentIdentityAuthenticatorUnauthorized
*/
type PatchCurrentIdentityAuthenticatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCurrentIdentityAuthenticatorUnauthorized creates PatchCurrentIdentityAuthenticatorUnauthorized with default headers values
func NewPatchCurrentIdentityAuthenticatorUnauthorized() *PatchCurrentIdentityAuthenticatorUnauthorized {

	return &PatchCurrentIdentityAuthenticatorUnauthorized{}
}

// WithPayload adds the payload to the patch current identity authenticator unauthorized response
func (o *PatchCurrentIdentityAuthenticatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCurrentIdentityAuthenticatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch current identity authenticator unauthorized response
func (o *PatchCurrentIdentityAuthenticatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCurrentIdentityAuthenticatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCurrentIdentityAuthenticatorNotFoundCode is the HTTP code returned for type PatchCurrentIdentityAuthenticatorNotFound
const PatchCurrentIdentityAuthenticatorNotFoundCode int = 404

/*PatchCurrentIdentityAuthenticatorNotFound The requested resource does not exist

swagger:response patchCurrentIdentityAuthenticatorNotFound
*/
type PatchCurrentIdentityAuthenticatorNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewPatchCurrentIdentityAuthenticatorNotFound creates PatchCurrentIdentityAuthenticatorNotFound with default headers values
func NewPatchCurrentIdentityAuthenticatorNotFound() *PatchCurrentIdentityAuthenticatorNotFound {

	return &PatchCurrentIdentityAuthenticatorNotFound{}
}

// WithPayload adds the payload to the patch current identity authenticator not found response
func (o *PatchCurrentIdentityAuthenticatorNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *PatchCurrentIdentityAuthenticatorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch current identity authenticator not found response
func (o *PatchCurrentIdentityAuthenticatorNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCurrentIdentityAuthenticatorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
