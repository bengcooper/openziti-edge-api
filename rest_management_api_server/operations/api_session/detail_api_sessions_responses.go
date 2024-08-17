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

// DetailAPISessionsOKCode is the HTTP code returned for type DetailAPISessionsOK
const DetailAPISessionsOKCode int = 200

/*
DetailAPISessionsOK Retrieves a singular API Session by id

swagger:response detailApiSessionsOK
*/
type DetailAPISessionsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailAPISessionEnvelope `json:"body,omitempty"`
}

// NewDetailAPISessionsOK creates DetailAPISessionsOK with default headers values
func NewDetailAPISessionsOK() *DetailAPISessionsOK {

	return &DetailAPISessionsOK{}
}

// WithPayload adds the payload to the detail Api sessions o k response
func (o *DetailAPISessionsOK) WithPayload(payload *rest_model.DetailAPISessionEnvelope) *DetailAPISessionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail Api sessions o k response
func (o *DetailAPISessionsOK) SetPayload(payload *rest_model.DetailAPISessionEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAPISessionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailAPISessionsUnauthorizedCode is the HTTP code returned for type DetailAPISessionsUnauthorized
const DetailAPISessionsUnauthorizedCode int = 401

/*
DetailAPISessionsUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailApiSessionsUnauthorized
*/
type DetailAPISessionsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailAPISessionsUnauthorized creates DetailAPISessionsUnauthorized with default headers values
func NewDetailAPISessionsUnauthorized() *DetailAPISessionsUnauthorized {

	return &DetailAPISessionsUnauthorized{}
}

// WithPayload adds the payload to the detail Api sessions unauthorized response
func (o *DetailAPISessionsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailAPISessionsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail Api sessions unauthorized response
func (o *DetailAPISessionsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAPISessionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailAPISessionsNotFoundCode is the HTTP code returned for type DetailAPISessionsNotFound
const DetailAPISessionsNotFoundCode int = 404

/*
DetailAPISessionsNotFound The requested resource does not exist

swagger:response detailApiSessionsNotFound
*/
type DetailAPISessionsNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailAPISessionsNotFound creates DetailAPISessionsNotFound with default headers values
func NewDetailAPISessionsNotFound() *DetailAPISessionsNotFound {

	return &DetailAPISessionsNotFound{}
}

// WithPayload adds the payload to the detail Api sessions not found response
func (o *DetailAPISessionsNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailAPISessionsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail Api sessions not found response
func (o *DetailAPISessionsNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAPISessionsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailAPISessionsTooManyRequestsCode is the HTTP code returned for type DetailAPISessionsTooManyRequests
const DetailAPISessionsTooManyRequestsCode int = 429

/*
DetailAPISessionsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailApiSessionsTooManyRequests
*/
type DetailAPISessionsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailAPISessionsTooManyRequests creates DetailAPISessionsTooManyRequests with default headers values
func NewDetailAPISessionsTooManyRequests() *DetailAPISessionsTooManyRequests {

	return &DetailAPISessionsTooManyRequests{}
}

// WithPayload adds the payload to the detail Api sessions too many requests response
func (o *DetailAPISessionsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailAPISessionsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail Api sessions too many requests response
func (o *DetailAPISessionsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailAPISessionsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
