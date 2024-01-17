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

// GetCurrentAPISessionOKCode is the HTTP code returned for type GetCurrentAPISessionOK
const GetCurrentAPISessionOKCode int = 200

/*GetCurrentAPISessionOK The API session associated with the session used to issue the request

swagger:response getCurrentApiSessionOK
*/
type GetCurrentAPISessionOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CurrentAPISessionDetailEnvelope `json:"body,omitempty"`
}

// NewGetCurrentAPISessionOK creates GetCurrentAPISessionOK with default headers values
func NewGetCurrentAPISessionOK() *GetCurrentAPISessionOK {

	return &GetCurrentAPISessionOK{}
}

// WithPayload adds the payload to the get current Api session o k response
func (o *GetCurrentAPISessionOK) WithPayload(payload *rest_model.CurrentAPISessionDetailEnvelope) *GetCurrentAPISessionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current Api session o k response
func (o *GetCurrentAPISessionOK) SetPayload(payload *rest_model.CurrentAPISessionDetailEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentAPISessionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentAPISessionUnauthorizedCode is the HTTP code returned for type GetCurrentAPISessionUnauthorized
const GetCurrentAPISessionUnauthorizedCode int = 401

/*GetCurrentAPISessionUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response getCurrentApiSessionUnauthorized
*/
type GetCurrentAPISessionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCurrentAPISessionUnauthorized creates GetCurrentAPISessionUnauthorized with default headers values
func NewGetCurrentAPISessionUnauthorized() *GetCurrentAPISessionUnauthorized {

	return &GetCurrentAPISessionUnauthorized{}
}

// WithPayload adds the payload to the get current Api session unauthorized response
func (o *GetCurrentAPISessionUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCurrentAPISessionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current Api session unauthorized response
func (o *GetCurrentAPISessionUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentAPISessionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
