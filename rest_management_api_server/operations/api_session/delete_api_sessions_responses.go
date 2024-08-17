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

package api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteAPISessionsOKCode is the HTTP code returned for type DeleteAPISessionsOK
const DeleteAPISessionsOKCode int = 200

/*
DeleteAPISessionsOK The delete request was successful and the resource has been removed

swagger:response deleteApiSessionsOK
*/
type DeleteAPISessionsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteAPISessionsOK creates DeleteAPISessionsOK with default headers values
func NewDeleteAPISessionsOK() *DeleteAPISessionsOK {

	return &DeleteAPISessionsOK{}
}

// WithPayload adds the payload to the delete Api sessions o k response
func (o *DeleteAPISessionsOK) WithPayload(payload *rest_model.Empty) *DeleteAPISessionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Api sessions o k response
func (o *DeleteAPISessionsOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAPISessionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAPISessionsUnauthorizedCode is the HTTP code returned for type DeleteAPISessionsUnauthorized
const DeleteAPISessionsUnauthorizedCode int = 401

/*
DeleteAPISessionsUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteApiSessionsUnauthorized
*/
type DeleteAPISessionsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteAPISessionsUnauthorized creates DeleteAPISessionsUnauthorized with default headers values
func NewDeleteAPISessionsUnauthorized() *DeleteAPISessionsUnauthorized {

	return &DeleteAPISessionsUnauthorized{}
}

// WithPayload adds the payload to the delete Api sessions unauthorized response
func (o *DeleteAPISessionsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteAPISessionsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Api sessions unauthorized response
func (o *DeleteAPISessionsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAPISessionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAPISessionsForbiddenCode is the HTTP code returned for type DeleteAPISessionsForbidden
const DeleteAPISessionsForbiddenCode int = 403

/*
DeleteAPISessionsForbidden The supplied session does not have the correct access rights to request this resource

swagger:response deleteApiSessionsForbidden
*/
type DeleteAPISessionsForbidden struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteAPISessionsForbidden creates DeleteAPISessionsForbidden with default headers values
func NewDeleteAPISessionsForbidden() *DeleteAPISessionsForbidden {

	return &DeleteAPISessionsForbidden{}
}

// WithPayload adds the payload to the delete Api sessions forbidden response
func (o *DeleteAPISessionsForbidden) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteAPISessionsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Api sessions forbidden response
func (o *DeleteAPISessionsForbidden) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAPISessionsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAPISessionsNotFoundCode is the HTTP code returned for type DeleteAPISessionsNotFound
const DeleteAPISessionsNotFoundCode int = 404

/*
DeleteAPISessionsNotFound The requested resource does not exist

swagger:response deleteApiSessionsNotFound
*/
type DeleteAPISessionsNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteAPISessionsNotFound creates DeleteAPISessionsNotFound with default headers values
func NewDeleteAPISessionsNotFound() *DeleteAPISessionsNotFound {

	return &DeleteAPISessionsNotFound{}
}

// WithPayload adds the payload to the delete Api sessions not found response
func (o *DeleteAPISessionsNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteAPISessionsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Api sessions not found response
func (o *DeleteAPISessionsNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAPISessionsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAPISessionsTooManyRequestsCode is the HTTP code returned for type DeleteAPISessionsTooManyRequests
const DeleteAPISessionsTooManyRequestsCode int = 429

/*
DeleteAPISessionsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteApiSessionsTooManyRequests
*/
type DeleteAPISessionsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteAPISessionsTooManyRequests creates DeleteAPISessionsTooManyRequests with default headers values
func NewDeleteAPISessionsTooManyRequests() *DeleteAPISessionsTooManyRequests {

	return &DeleteAPISessionsTooManyRequests{}
}

// WithPayload adds the payload to the delete Api sessions too many requests response
func (o *DeleteAPISessionsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteAPISessionsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Api sessions too many requests response
func (o *DeleteAPISessionsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAPISessionsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
