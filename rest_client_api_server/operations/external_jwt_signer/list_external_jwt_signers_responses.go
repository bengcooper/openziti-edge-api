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

package external_jwt_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListExternalJWTSignersOKCode is the HTTP code returned for type ListExternalJWTSignersOK
const ListExternalJWTSignersOKCode int = 200

/*ListExternalJWTSignersOK A list of External JWT Signers

swagger:response listExternalJwtSignersOK
*/
type ListExternalJWTSignersOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListClientExternalJWTSignersEnvelope `json:"body,omitempty"`
}

// NewListExternalJWTSignersOK creates ListExternalJWTSignersOK with default headers values
func NewListExternalJWTSignersOK() *ListExternalJWTSignersOK {

	return &ListExternalJWTSignersOK{}
}

// WithPayload adds the payload to the list external Jwt signers o k response
func (o *ListExternalJWTSignersOK) WithPayload(payload *rest_model.ListClientExternalJWTSignersEnvelope) *ListExternalJWTSignersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list external Jwt signers o k response
func (o *ListExternalJWTSignersOK) SetPayload(payload *rest_model.ListClientExternalJWTSignersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExternalJWTSignersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListExternalJWTSignersBadRequestCode is the HTTP code returned for type ListExternalJWTSignersBadRequest
const ListExternalJWTSignersBadRequestCode int = 400

/*ListExternalJWTSignersBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listExternalJwtSignersBadRequest
*/
type ListExternalJWTSignersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListExternalJWTSignersBadRequest creates ListExternalJWTSignersBadRequest with default headers values
func NewListExternalJWTSignersBadRequest() *ListExternalJWTSignersBadRequest {

	return &ListExternalJWTSignersBadRequest{}
}

// WithPayload adds the payload to the list external Jwt signers bad request response
func (o *ListExternalJWTSignersBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListExternalJWTSignersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list external Jwt signers bad request response
func (o *ListExternalJWTSignersBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExternalJWTSignersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListExternalJWTSignersUnauthorizedCode is the HTTP code returned for type ListExternalJWTSignersUnauthorized
const ListExternalJWTSignersUnauthorizedCode int = 401

/*ListExternalJWTSignersUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listExternalJwtSignersUnauthorized
*/
type ListExternalJWTSignersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListExternalJWTSignersUnauthorized creates ListExternalJWTSignersUnauthorized with default headers values
func NewListExternalJWTSignersUnauthorized() *ListExternalJWTSignersUnauthorized {

	return &ListExternalJWTSignersUnauthorized{}
}

// WithPayload adds the payload to the list external Jwt signers unauthorized response
func (o *ListExternalJWTSignersUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListExternalJWTSignersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list external Jwt signers unauthorized response
func (o *ListExternalJWTSignersUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExternalJWTSignersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListExternalJWTSignersTooManyRequestsCode is the HTTP code returned for type ListExternalJWTSignersTooManyRequests
const ListExternalJWTSignersTooManyRequestsCode int = 429

/*ListExternalJWTSignersTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listExternalJwtSignersTooManyRequests
*/
type ListExternalJWTSignersTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListExternalJWTSignersTooManyRequests creates ListExternalJWTSignersTooManyRequests with default headers values
func NewListExternalJWTSignersTooManyRequests() *ListExternalJWTSignersTooManyRequests {

	return &ListExternalJWTSignersTooManyRequests{}
}

// WithPayload adds the payload to the list external Jwt signers too many requests response
func (o *ListExternalJWTSignersTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListExternalJWTSignersTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list external Jwt signers too many requests response
func (o *ListExternalJWTSignersTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExternalJWTSignersTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
