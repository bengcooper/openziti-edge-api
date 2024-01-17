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

package service_edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DeleteServiceEdgeRouterPolicyOKCode is the HTTP code returned for type DeleteServiceEdgeRouterPolicyOK
const DeleteServiceEdgeRouterPolicyOKCode int = 200

/*DeleteServiceEdgeRouterPolicyOK The delete request was successful and the resource has been removed

swagger:response deleteServiceEdgeRouterPolicyOK
*/
type DeleteServiceEdgeRouterPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteServiceEdgeRouterPolicyOK creates DeleteServiceEdgeRouterPolicyOK with default headers values
func NewDeleteServiceEdgeRouterPolicyOK() *DeleteServiceEdgeRouterPolicyOK {

	return &DeleteServiceEdgeRouterPolicyOK{}
}

// WithPayload adds the payload to the delete service edge router policy o k response
func (o *DeleteServiceEdgeRouterPolicyOK) WithPayload(payload *rest_model.Empty) *DeleteServiceEdgeRouterPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service edge router policy o k response
func (o *DeleteServiceEdgeRouterPolicyOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceEdgeRouterPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceEdgeRouterPolicyBadRequestCode is the HTTP code returned for type DeleteServiceEdgeRouterPolicyBadRequest
const DeleteServiceEdgeRouterPolicyBadRequestCode int = 400

/*DeleteServiceEdgeRouterPolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteServiceEdgeRouterPolicyBadRequest
*/
type DeleteServiceEdgeRouterPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceEdgeRouterPolicyBadRequest creates DeleteServiceEdgeRouterPolicyBadRequest with default headers values
func NewDeleteServiceEdgeRouterPolicyBadRequest() *DeleteServiceEdgeRouterPolicyBadRequest {

	return &DeleteServiceEdgeRouterPolicyBadRequest{}
}

// WithPayload adds the payload to the delete service edge router policy bad request response
func (o *DeleteServiceEdgeRouterPolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceEdgeRouterPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service edge router policy bad request response
func (o *DeleteServiceEdgeRouterPolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceEdgeRouterPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceEdgeRouterPolicyUnauthorizedCode is the HTTP code returned for type DeleteServiceEdgeRouterPolicyUnauthorized
const DeleteServiceEdgeRouterPolicyUnauthorizedCode int = 401

/*DeleteServiceEdgeRouterPolicyUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response deleteServiceEdgeRouterPolicyUnauthorized
*/
type DeleteServiceEdgeRouterPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceEdgeRouterPolicyUnauthorized creates DeleteServiceEdgeRouterPolicyUnauthorized with default headers values
func NewDeleteServiceEdgeRouterPolicyUnauthorized() *DeleteServiceEdgeRouterPolicyUnauthorized {

	return &DeleteServiceEdgeRouterPolicyUnauthorized{}
}

// WithPayload adds the payload to the delete service edge router policy unauthorized response
func (o *DeleteServiceEdgeRouterPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceEdgeRouterPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service edge router policy unauthorized response
func (o *DeleteServiceEdgeRouterPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceEdgeRouterPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceEdgeRouterPolicyConflictCode is the HTTP code returned for type DeleteServiceEdgeRouterPolicyConflict
const DeleteServiceEdgeRouterPolicyConflictCode int = 409

/*DeleteServiceEdgeRouterPolicyConflict The resource requested to be removed/altered cannot be as it is referenced by another object.

swagger:response deleteServiceEdgeRouterPolicyConflict
*/
type DeleteServiceEdgeRouterPolicyConflict struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceEdgeRouterPolicyConflict creates DeleteServiceEdgeRouterPolicyConflict with default headers values
func NewDeleteServiceEdgeRouterPolicyConflict() *DeleteServiceEdgeRouterPolicyConflict {

	return &DeleteServiceEdgeRouterPolicyConflict{}
}

// WithPayload adds the payload to the delete service edge router policy conflict response
func (o *DeleteServiceEdgeRouterPolicyConflict) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceEdgeRouterPolicyConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service edge router policy conflict response
func (o *DeleteServiceEdgeRouterPolicyConflict) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceEdgeRouterPolicyConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceEdgeRouterPolicyTooManyRequestsCode is the HTTP code returned for type DeleteServiceEdgeRouterPolicyTooManyRequests
const DeleteServiceEdgeRouterPolicyTooManyRequestsCode int = 429

/*DeleteServiceEdgeRouterPolicyTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteServiceEdgeRouterPolicyTooManyRequests
*/
type DeleteServiceEdgeRouterPolicyTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteServiceEdgeRouterPolicyTooManyRequests creates DeleteServiceEdgeRouterPolicyTooManyRequests with default headers values
func NewDeleteServiceEdgeRouterPolicyTooManyRequests() *DeleteServiceEdgeRouterPolicyTooManyRequests {

	return &DeleteServiceEdgeRouterPolicyTooManyRequests{}
}

// WithPayload adds the payload to the delete service edge router policy too many requests response
func (o *DeleteServiceEdgeRouterPolicyTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteServiceEdgeRouterPolicyTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service edge router policy too many requests response
func (o *DeleteServiceEdgeRouterPolicyTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceEdgeRouterPolicyTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
