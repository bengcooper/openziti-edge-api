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

package edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListEdgeRouterPolicyEdgeRoutersOKCode is the HTTP code returned for type ListEdgeRouterPolicyEdgeRoutersOK
const ListEdgeRouterPolicyEdgeRoutersOKCode int = 200

/*ListEdgeRouterPolicyEdgeRoutersOK A list of edge routers

swagger:response listEdgeRouterPolicyEdgeRoutersOK
*/
type ListEdgeRouterPolicyEdgeRoutersOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListEdgeRoutersEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterPolicyEdgeRoutersOK creates ListEdgeRouterPolicyEdgeRoutersOK with default headers values
func NewListEdgeRouterPolicyEdgeRoutersOK() *ListEdgeRouterPolicyEdgeRoutersOK {

	return &ListEdgeRouterPolicyEdgeRoutersOK{}
}

// WithPayload adds the payload to the list edge router policy edge routers o k response
func (o *ListEdgeRouterPolicyEdgeRoutersOK) WithPayload(payload *rest_model.ListEdgeRoutersEnvelope) *ListEdgeRouterPolicyEdgeRoutersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router policy edge routers o k response
func (o *ListEdgeRouterPolicyEdgeRoutersOK) SetPayload(payload *rest_model.ListEdgeRoutersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterPolicyEdgeRoutersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterPolicyEdgeRoutersUnauthorizedCode is the HTTP code returned for type ListEdgeRouterPolicyEdgeRoutersUnauthorized
const ListEdgeRouterPolicyEdgeRoutersUnauthorizedCode int = 401

/*ListEdgeRouterPolicyEdgeRoutersUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listEdgeRouterPolicyEdgeRoutersUnauthorized
*/
type ListEdgeRouterPolicyEdgeRoutersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterPolicyEdgeRoutersUnauthorized creates ListEdgeRouterPolicyEdgeRoutersUnauthorized with default headers values
func NewListEdgeRouterPolicyEdgeRoutersUnauthorized() *ListEdgeRouterPolicyEdgeRoutersUnauthorized {

	return &ListEdgeRouterPolicyEdgeRoutersUnauthorized{}
}

// WithPayload adds the payload to the list edge router policy edge routers unauthorized response
func (o *ListEdgeRouterPolicyEdgeRoutersUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterPolicyEdgeRoutersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router policy edge routers unauthorized response
func (o *ListEdgeRouterPolicyEdgeRoutersUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterPolicyEdgeRoutersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterPolicyEdgeRoutersNotFoundCode is the HTTP code returned for type ListEdgeRouterPolicyEdgeRoutersNotFound
const ListEdgeRouterPolicyEdgeRoutersNotFoundCode int = 404

/*ListEdgeRouterPolicyEdgeRoutersNotFound The requested resource does not exist

swagger:response listEdgeRouterPolicyEdgeRoutersNotFound
*/
type ListEdgeRouterPolicyEdgeRoutersNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterPolicyEdgeRoutersNotFound creates ListEdgeRouterPolicyEdgeRoutersNotFound with default headers values
func NewListEdgeRouterPolicyEdgeRoutersNotFound() *ListEdgeRouterPolicyEdgeRoutersNotFound {

	return &ListEdgeRouterPolicyEdgeRoutersNotFound{}
}

// WithPayload adds the payload to the list edge router policy edge routers not found response
func (o *ListEdgeRouterPolicyEdgeRoutersNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterPolicyEdgeRoutersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router policy edge routers not found response
func (o *ListEdgeRouterPolicyEdgeRoutersNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterPolicyEdgeRoutersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterPolicyEdgeRoutersTooManyRequestsCode is the HTTP code returned for type ListEdgeRouterPolicyEdgeRoutersTooManyRequests
const ListEdgeRouterPolicyEdgeRoutersTooManyRequestsCode int = 429

/*ListEdgeRouterPolicyEdgeRoutersTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listEdgeRouterPolicyEdgeRoutersTooManyRequests
*/
type ListEdgeRouterPolicyEdgeRoutersTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterPolicyEdgeRoutersTooManyRequests creates ListEdgeRouterPolicyEdgeRoutersTooManyRequests with default headers values
func NewListEdgeRouterPolicyEdgeRoutersTooManyRequests() *ListEdgeRouterPolicyEdgeRoutersTooManyRequests {

	return &ListEdgeRouterPolicyEdgeRoutersTooManyRequests{}
}

// WithPayload adds the payload to the list edge router policy edge routers too many requests response
func (o *ListEdgeRouterPolicyEdgeRoutersTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterPolicyEdgeRoutersTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router policy edge routers too many requests response
func (o *ListEdgeRouterPolicyEdgeRoutersTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterPolicyEdgeRoutersTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
