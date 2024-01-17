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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// GetCurrentIdentityEdgeRoutersOKCode is the HTTP code returned for type GetCurrentIdentityEdgeRoutersOK
const GetCurrentIdentityEdgeRoutersOKCode int = 200

/*GetCurrentIdentityEdgeRoutersOK A list of edge routers

swagger:response getCurrentIdentityEdgeRoutersOK
*/
type GetCurrentIdentityEdgeRoutersOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListCurrentIdentityEdgeRoutersEnvelope `json:"body,omitempty"`
}

// NewGetCurrentIdentityEdgeRoutersOK creates GetCurrentIdentityEdgeRoutersOK with default headers values
func NewGetCurrentIdentityEdgeRoutersOK() *GetCurrentIdentityEdgeRoutersOK {

	return &GetCurrentIdentityEdgeRoutersOK{}
}

// WithPayload adds the payload to the get current identity edge routers o k response
func (o *GetCurrentIdentityEdgeRoutersOK) WithPayload(payload *rest_model.ListCurrentIdentityEdgeRoutersEnvelope) *GetCurrentIdentityEdgeRoutersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current identity edge routers o k response
func (o *GetCurrentIdentityEdgeRoutersOK) SetPayload(payload *rest_model.ListCurrentIdentityEdgeRoutersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentIdentityEdgeRoutersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentIdentityEdgeRoutersUnauthorizedCode is the HTTP code returned for type GetCurrentIdentityEdgeRoutersUnauthorized
const GetCurrentIdentityEdgeRoutersUnauthorizedCode int = 401

/*GetCurrentIdentityEdgeRoutersUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response getCurrentIdentityEdgeRoutersUnauthorized
*/
type GetCurrentIdentityEdgeRoutersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCurrentIdentityEdgeRoutersUnauthorized creates GetCurrentIdentityEdgeRoutersUnauthorized with default headers values
func NewGetCurrentIdentityEdgeRoutersUnauthorized() *GetCurrentIdentityEdgeRoutersUnauthorized {

	return &GetCurrentIdentityEdgeRoutersUnauthorized{}
}

// WithPayload adds the payload to the get current identity edge routers unauthorized response
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCurrentIdentityEdgeRoutersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current identity edge routers unauthorized response
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentIdentityEdgeRoutersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentIdentityEdgeRoutersTooManyRequestsCode is the HTTP code returned for type GetCurrentIdentityEdgeRoutersTooManyRequests
const GetCurrentIdentityEdgeRoutersTooManyRequestsCode int = 429

/*GetCurrentIdentityEdgeRoutersTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response getCurrentIdentityEdgeRoutersTooManyRequests
*/
type GetCurrentIdentityEdgeRoutersTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCurrentIdentityEdgeRoutersTooManyRequests creates GetCurrentIdentityEdgeRoutersTooManyRequests with default headers values
func NewGetCurrentIdentityEdgeRoutersTooManyRequests() *GetCurrentIdentityEdgeRoutersTooManyRequests {

	return &GetCurrentIdentityEdgeRoutersTooManyRequests{}
}

// WithPayload adds the payload to the get current identity edge routers too many requests response
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCurrentIdentityEdgeRoutersTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current identity edge routers too many requests response
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentIdentityEdgeRoutersTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
