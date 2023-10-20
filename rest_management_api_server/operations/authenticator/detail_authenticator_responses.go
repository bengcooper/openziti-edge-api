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

package authenticator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DetailAuthenticatorOKCode is the HTTP code returned for type DetailAuthenticatorOK
const DetailAuthenticatorOKCode int = 200

/*DetailAuthenticatorOK A singular authenticator resource

swagger:response detailAuthenticatorOK
*/
type DetailAuthenticatorOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailAuthenticatorEnvelope `json:"body,omitempty"`
}

// NewDetailAuthenticatorOK creates DetailAuthenticatorOK with default headers values
func NewDetailAuthenticatorOK() *DetailAuthenticatorOK {

	return &DetailAuthenticatorOK{}
}

// WithPayload adds the payload to the detail authenticator o k response
func (o *DetailAuthenticatorOK) WithPayload(payload *rest_model.DetailAuthenticatorEnvelope) *DetailAuthenticatorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail authenticator o k response
func (o *DetailAuthenticatorOK) SetPayload(payload *rest_model.DetailAuthenticatorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAuthenticatorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailAuthenticatorUnauthorizedCode is the HTTP code returned for type DetailAuthenticatorUnauthorized
const DetailAuthenticatorUnauthorizedCode int = 401

/*DetailAuthenticatorUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailAuthenticatorUnauthorized
*/
type DetailAuthenticatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailAuthenticatorUnauthorized creates DetailAuthenticatorUnauthorized with default headers values
func NewDetailAuthenticatorUnauthorized() *DetailAuthenticatorUnauthorized {

	return &DetailAuthenticatorUnauthorized{}
}

// WithPayload adds the payload to the detail authenticator unauthorized response
func (o *DetailAuthenticatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailAuthenticatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail authenticator unauthorized response
func (o *DetailAuthenticatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAuthenticatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailAuthenticatorNotFoundCode is the HTTP code returned for type DetailAuthenticatorNotFound
const DetailAuthenticatorNotFoundCode int = 404

/*DetailAuthenticatorNotFound The requested resource does not exist

swagger:response detailAuthenticatorNotFound
*/
type DetailAuthenticatorNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailAuthenticatorNotFound creates DetailAuthenticatorNotFound with default headers values
func NewDetailAuthenticatorNotFound() *DetailAuthenticatorNotFound {

	return &DetailAuthenticatorNotFound{}
}

// WithPayload adds the payload to the detail authenticator not found response
func (o *DetailAuthenticatorNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailAuthenticatorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail authenticator not found response
func (o *DetailAuthenticatorNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAuthenticatorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailAuthenticatorTooManyRequestsCode is the HTTP code returned for type DetailAuthenticatorTooManyRequests
const DetailAuthenticatorTooManyRequestsCode int = 429

/*DetailAuthenticatorTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailAuthenticatorTooManyRequests
*/
type DetailAuthenticatorTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailAuthenticatorTooManyRequests creates DetailAuthenticatorTooManyRequests with default headers values
func NewDetailAuthenticatorTooManyRequests() *DetailAuthenticatorTooManyRequests {

	return &DetailAuthenticatorTooManyRequests{}
}

// WithPayload adds the payload to the detail authenticator too many requests response
func (o *DetailAuthenticatorTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailAuthenticatorTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail authenticator too many requests response
func (o *DetailAuthenticatorTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAuthenticatorTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
