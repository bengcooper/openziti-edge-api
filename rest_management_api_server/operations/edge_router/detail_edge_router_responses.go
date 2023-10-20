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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DetailEdgeRouterOKCode is the HTTP code returned for type DetailEdgeRouterOK
const DetailEdgeRouterOKCode int = 200

/*DetailEdgeRouterOK A singular edge router resource

swagger:response detailEdgeRouterOK
*/
type DetailEdgeRouterOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailedEdgeRouterEnvelope `json:"body,omitempty"`
}

// NewDetailEdgeRouterOK creates DetailEdgeRouterOK with default headers values
func NewDetailEdgeRouterOK() *DetailEdgeRouterOK {

	return &DetailEdgeRouterOK{}
}

// WithPayload adds the payload to the detail edge router o k response
func (o *DetailEdgeRouterOK) WithPayload(payload *rest_model.DetailedEdgeRouterEnvelope) *DetailEdgeRouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail edge router o k response
func (o *DetailEdgeRouterOK) SetPayload(payload *rest_model.DetailedEdgeRouterEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailEdgeRouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailEdgeRouterUnauthorizedCode is the HTTP code returned for type DetailEdgeRouterUnauthorized
const DetailEdgeRouterUnauthorizedCode int = 401

/*DetailEdgeRouterUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailEdgeRouterUnauthorized
*/
type DetailEdgeRouterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailEdgeRouterUnauthorized creates DetailEdgeRouterUnauthorized with default headers values
func NewDetailEdgeRouterUnauthorized() *DetailEdgeRouterUnauthorized {

	return &DetailEdgeRouterUnauthorized{}
}

// WithPayload adds the payload to the detail edge router unauthorized response
func (o *DetailEdgeRouterUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailEdgeRouterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail edge router unauthorized response
func (o *DetailEdgeRouterUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailEdgeRouterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailEdgeRouterNotFoundCode is the HTTP code returned for type DetailEdgeRouterNotFound
const DetailEdgeRouterNotFoundCode int = 404

/*DetailEdgeRouterNotFound The requested resource does not exist

swagger:response detailEdgeRouterNotFound
*/
type DetailEdgeRouterNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailEdgeRouterNotFound creates DetailEdgeRouterNotFound with default headers values
func NewDetailEdgeRouterNotFound() *DetailEdgeRouterNotFound {

	return &DetailEdgeRouterNotFound{}
}

// WithPayload adds the payload to the detail edge router not found response
func (o *DetailEdgeRouterNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailEdgeRouterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail edge router not found response
func (o *DetailEdgeRouterNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailEdgeRouterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailEdgeRouterTooManyRequestsCode is the HTTP code returned for type DetailEdgeRouterTooManyRequests
const DetailEdgeRouterTooManyRequestsCode int = 429

/*DetailEdgeRouterTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailEdgeRouterTooManyRequests
*/
type DetailEdgeRouterTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailEdgeRouterTooManyRequests creates DetailEdgeRouterTooManyRequests with default headers values
func NewDetailEdgeRouterTooManyRequests() *DetailEdgeRouterTooManyRequests {

	return &DetailEdgeRouterTooManyRequests{}
}

// WithPayload adds the payload to the detail edge router too many requests response
func (o *DetailEdgeRouterTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailEdgeRouterTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail edge router too many requests response
func (o *DetailEdgeRouterTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailEdgeRouterTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
