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

// DetailCurrentIdentityAuthenticatorOKCode is the HTTP code returned for type DetailCurrentIdentityAuthenticatorOK
const DetailCurrentIdentityAuthenticatorOKCode int = 200

/*DetailCurrentIdentityAuthenticatorOK A singular authenticator resource

swagger:response detailCurrentIdentityAuthenticatorOK
*/
type DetailCurrentIdentityAuthenticatorOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailAuthenticatorEnvelope `json:"body,omitempty"`
}

// NewDetailCurrentIdentityAuthenticatorOK creates DetailCurrentIdentityAuthenticatorOK with default headers values
func NewDetailCurrentIdentityAuthenticatorOK() *DetailCurrentIdentityAuthenticatorOK {

	return &DetailCurrentIdentityAuthenticatorOK{}
}

// WithPayload adds the payload to the detail current identity authenticator o k response
func (o *DetailCurrentIdentityAuthenticatorOK) WithPayload(payload *rest_model.DetailAuthenticatorEnvelope) *DetailCurrentIdentityAuthenticatorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail current identity authenticator o k response
func (o *DetailCurrentIdentityAuthenticatorOK) SetPayload(payload *rest_model.DetailAuthenticatorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCurrentIdentityAuthenticatorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCurrentIdentityAuthenticatorUnauthorizedCode is the HTTP code returned for type DetailCurrentIdentityAuthenticatorUnauthorized
const DetailCurrentIdentityAuthenticatorUnauthorizedCode int = 401

/*DetailCurrentIdentityAuthenticatorUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailCurrentIdentityAuthenticatorUnauthorized
*/
type DetailCurrentIdentityAuthenticatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCurrentIdentityAuthenticatorUnauthorized creates DetailCurrentIdentityAuthenticatorUnauthorized with default headers values
func NewDetailCurrentIdentityAuthenticatorUnauthorized() *DetailCurrentIdentityAuthenticatorUnauthorized {

	return &DetailCurrentIdentityAuthenticatorUnauthorized{}
}

// WithPayload adds the payload to the detail current identity authenticator unauthorized response
func (o *DetailCurrentIdentityAuthenticatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCurrentIdentityAuthenticatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail current identity authenticator unauthorized response
func (o *DetailCurrentIdentityAuthenticatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCurrentIdentityAuthenticatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCurrentIdentityAuthenticatorNotFoundCode is the HTTP code returned for type DetailCurrentIdentityAuthenticatorNotFound
const DetailCurrentIdentityAuthenticatorNotFoundCode int = 404

/*DetailCurrentIdentityAuthenticatorNotFound The requested resource does not exist

swagger:response detailCurrentIdentityAuthenticatorNotFound
*/
type DetailCurrentIdentityAuthenticatorNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCurrentIdentityAuthenticatorNotFound creates DetailCurrentIdentityAuthenticatorNotFound with default headers values
func NewDetailCurrentIdentityAuthenticatorNotFound() *DetailCurrentIdentityAuthenticatorNotFound {

	return &DetailCurrentIdentityAuthenticatorNotFound{}
}

// WithPayload adds the payload to the detail current identity authenticator not found response
func (o *DetailCurrentIdentityAuthenticatorNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCurrentIdentityAuthenticatorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail current identity authenticator not found response
func (o *DetailCurrentIdentityAuthenticatorNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCurrentIdentityAuthenticatorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
