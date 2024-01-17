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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// EnrollMfaCreatedCode is the HTTP code returned for type EnrollMfaCreated
const EnrollMfaCreatedCode int = 201

/*EnrollMfaCreated The create request was successful and the resource has been added at the following location

swagger:response enrollMfaCreated
*/
type EnrollMfaCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewEnrollMfaCreated creates EnrollMfaCreated with default headers values
func NewEnrollMfaCreated() *EnrollMfaCreated {

	return &EnrollMfaCreated{}
}

// WithPayload adds the payload to the enroll mfa created response
func (o *EnrollMfaCreated) WithPayload(payload *rest_model.CreateEnvelope) *EnrollMfaCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll mfa created response
func (o *EnrollMfaCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollMfaCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollMfaUnauthorizedCode is the HTTP code returned for type EnrollMfaUnauthorized
const EnrollMfaUnauthorizedCode int = 401

/*EnrollMfaUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response enrollMfaUnauthorized
*/
type EnrollMfaUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollMfaUnauthorized creates EnrollMfaUnauthorized with default headers values
func NewEnrollMfaUnauthorized() *EnrollMfaUnauthorized {

	return &EnrollMfaUnauthorized{}
}

// WithPayload adds the payload to the enroll mfa unauthorized response
func (o *EnrollMfaUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollMfaUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll mfa unauthorized response
func (o *EnrollMfaUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollMfaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollMfaConflictCode is the HTTP code returned for type EnrollMfaConflict
const EnrollMfaConflictCode int = 409

/*EnrollMfaConflict The identity is already enrolled in MFA

swagger:response enrollMfaConflict
*/
type EnrollMfaConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollMfaConflict creates EnrollMfaConflict with default headers values
func NewEnrollMfaConflict() *EnrollMfaConflict {

	return &EnrollMfaConflict{}
}

// WithPayload adds the payload to the enroll mfa conflict response
func (o *EnrollMfaConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollMfaConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll mfa conflict response
func (o *EnrollMfaConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollMfaConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
