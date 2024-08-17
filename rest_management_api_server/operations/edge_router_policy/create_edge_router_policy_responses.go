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

// CreateEdgeRouterPolicyCreatedCode is the HTTP code returned for type CreateEdgeRouterPolicyCreated
const CreateEdgeRouterPolicyCreatedCode int = 201

/*
CreateEdgeRouterPolicyCreated The create request was successful and the resource has been added at the following location

swagger:response createEdgeRouterPolicyCreated
*/
type CreateEdgeRouterPolicyCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterPolicyCreated creates CreateEdgeRouterPolicyCreated with default headers values
func NewCreateEdgeRouterPolicyCreated() *CreateEdgeRouterPolicyCreated {

	return &CreateEdgeRouterPolicyCreated{}
}

// WithPayload adds the payload to the create edge router policy created response
func (o *CreateEdgeRouterPolicyCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateEdgeRouterPolicyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router policy created response
func (o *CreateEdgeRouterPolicyCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterPolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEdgeRouterPolicyBadRequestCode is the HTTP code returned for type CreateEdgeRouterPolicyBadRequest
const CreateEdgeRouterPolicyBadRequestCode int = 400

/*
CreateEdgeRouterPolicyBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createEdgeRouterPolicyBadRequest
*/
type CreateEdgeRouterPolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterPolicyBadRequest creates CreateEdgeRouterPolicyBadRequest with default headers values
func NewCreateEdgeRouterPolicyBadRequest() *CreateEdgeRouterPolicyBadRequest {

	return &CreateEdgeRouterPolicyBadRequest{}
}

// WithPayload adds the payload to the create edge router policy bad request response
func (o *CreateEdgeRouterPolicyBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateEdgeRouterPolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router policy bad request response
func (o *CreateEdgeRouterPolicyBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterPolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEdgeRouterPolicyUnauthorizedCode is the HTTP code returned for type CreateEdgeRouterPolicyUnauthorized
const CreateEdgeRouterPolicyUnauthorizedCode int = 401

/*
CreateEdgeRouterPolicyUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response createEdgeRouterPolicyUnauthorized
*/
type CreateEdgeRouterPolicyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterPolicyUnauthorized creates CreateEdgeRouterPolicyUnauthorized with default headers values
func NewCreateEdgeRouterPolicyUnauthorized() *CreateEdgeRouterPolicyUnauthorized {

	return &CreateEdgeRouterPolicyUnauthorized{}
}

// WithPayload adds the payload to the create edge router policy unauthorized response
func (o *CreateEdgeRouterPolicyUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateEdgeRouterPolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router policy unauthorized response
func (o *CreateEdgeRouterPolicyUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEdgeRouterPolicyTooManyRequestsCode is the HTTP code returned for type CreateEdgeRouterPolicyTooManyRequests
const CreateEdgeRouterPolicyTooManyRequestsCode int = 429

/*
CreateEdgeRouterPolicyTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createEdgeRouterPolicyTooManyRequests
*/
type CreateEdgeRouterPolicyTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateEdgeRouterPolicyTooManyRequests creates CreateEdgeRouterPolicyTooManyRequests with default headers values
func NewCreateEdgeRouterPolicyTooManyRequests() *CreateEdgeRouterPolicyTooManyRequests {

	return &CreateEdgeRouterPolicyTooManyRequests{}
}

// WithPayload adds the payload to the create edge router policy too many requests response
func (o *CreateEdgeRouterPolicyTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateEdgeRouterPolicyTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create edge router policy too many requests response
func (o *CreateEdgeRouterPolicyTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEdgeRouterPolicyTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
