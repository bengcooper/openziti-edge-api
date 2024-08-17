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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListConfigServicesOKCode is the HTTP code returned for type ListConfigServicesOK
const ListConfigServicesOKCode int = 200

/*
ListConfigServicesOK A list of services

swagger:response listConfigServicesOK
*/
type ListConfigServicesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServicesEnvelope `json:"body,omitempty"`
}

// NewListConfigServicesOK creates ListConfigServicesOK with default headers values
func NewListConfigServicesOK() *ListConfigServicesOK {

	return &ListConfigServicesOK{}
}

// WithPayload adds the payload to the list config services o k response
func (o *ListConfigServicesOK) WithPayload(payload *rest_model.ListServicesEnvelope) *ListConfigServicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config services o k response
func (o *ListConfigServicesOK) SetPayload(payload *rest_model.ListServicesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigServicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigServicesBadRequestCode is the HTTP code returned for type ListConfigServicesBadRequest
const ListConfigServicesBadRequestCode int = 400

/*
ListConfigServicesBadRequest The requested resource does not exist

swagger:response listConfigServicesBadRequest
*/
type ListConfigServicesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigServicesBadRequest creates ListConfigServicesBadRequest with default headers values
func NewListConfigServicesBadRequest() *ListConfigServicesBadRequest {

	return &ListConfigServicesBadRequest{}
}

// WithPayload adds the payload to the list config services bad request response
func (o *ListConfigServicesBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigServicesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config services bad request response
func (o *ListConfigServicesBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigServicesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigServicesUnauthorizedCode is the HTTP code returned for type ListConfigServicesUnauthorized
const ListConfigServicesUnauthorizedCode int = 401

/*
ListConfigServicesUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listConfigServicesUnauthorized
*/
type ListConfigServicesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigServicesUnauthorized creates ListConfigServicesUnauthorized with default headers values
func NewListConfigServicesUnauthorized() *ListConfigServicesUnauthorized {

	return &ListConfigServicesUnauthorized{}
}

// WithPayload adds the payload to the list config services unauthorized response
func (o *ListConfigServicesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigServicesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config services unauthorized response
func (o *ListConfigServicesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigServicesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListConfigServicesTooManyRequestsCode is the HTTP code returned for type ListConfigServicesTooManyRequests
const ListConfigServicesTooManyRequestsCode int = 429

/*
ListConfigServicesTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listConfigServicesTooManyRequests
*/
type ListConfigServicesTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListConfigServicesTooManyRequests creates ListConfigServicesTooManyRequests with default headers values
func NewListConfigServicesTooManyRequests() *ListConfigServicesTooManyRequests {

	return &ListConfigServicesTooManyRequests{}
}

// WithPayload adds the payload to the list config services too many requests response
func (o *ListConfigServicesTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListConfigServicesTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list config services too many requests response
func (o *ListConfigServicesTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListConfigServicesTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
