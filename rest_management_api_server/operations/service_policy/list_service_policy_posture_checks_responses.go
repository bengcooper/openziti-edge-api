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

package service_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListServicePolicyPostureChecksOKCode is the HTTP code returned for type ListServicePolicyPostureChecksOK
const ListServicePolicyPostureChecksOKCode int = 200

/*ListServicePolicyPostureChecksOK A list of posture checks

swagger:response listServicePolicyPostureChecksOK
*/
type ListServicePolicyPostureChecksOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListPostureCheckEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyPostureChecksOK creates ListServicePolicyPostureChecksOK with default headers values
func NewListServicePolicyPostureChecksOK() *ListServicePolicyPostureChecksOK {

	return &ListServicePolicyPostureChecksOK{}
}

// WithPayload adds the payload to the list service policy posture checks o k response
func (o *ListServicePolicyPostureChecksOK) WithPayload(payload *rest_model.ListPostureCheckEnvelope) *ListServicePolicyPostureChecksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy posture checks o k response
func (o *ListServicePolicyPostureChecksOK) SetPayload(payload *rest_model.ListPostureCheckEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyPostureChecksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicePolicyPostureChecksBadRequestCode is the HTTP code returned for type ListServicePolicyPostureChecksBadRequest
const ListServicePolicyPostureChecksBadRequestCode int = 400

/*ListServicePolicyPostureChecksBadRequest The requested resource does not exist

swagger:response listServicePolicyPostureChecksBadRequest
*/
type ListServicePolicyPostureChecksBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyPostureChecksBadRequest creates ListServicePolicyPostureChecksBadRequest with default headers values
func NewListServicePolicyPostureChecksBadRequest() *ListServicePolicyPostureChecksBadRequest {

	return &ListServicePolicyPostureChecksBadRequest{}
}

// WithPayload adds the payload to the list service policy posture checks bad request response
func (o *ListServicePolicyPostureChecksBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicePolicyPostureChecksBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy posture checks bad request response
func (o *ListServicePolicyPostureChecksBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyPostureChecksBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicePolicyPostureChecksUnauthorizedCode is the HTTP code returned for type ListServicePolicyPostureChecksUnauthorized
const ListServicePolicyPostureChecksUnauthorizedCode int = 401

/*ListServicePolicyPostureChecksUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServicePolicyPostureChecksUnauthorized
*/
type ListServicePolicyPostureChecksUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyPostureChecksUnauthorized creates ListServicePolicyPostureChecksUnauthorized with default headers values
func NewListServicePolicyPostureChecksUnauthorized() *ListServicePolicyPostureChecksUnauthorized {

	return &ListServicePolicyPostureChecksUnauthorized{}
}

// WithPayload adds the payload to the list service policy posture checks unauthorized response
func (o *ListServicePolicyPostureChecksUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicePolicyPostureChecksUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy posture checks unauthorized response
func (o *ListServicePolicyPostureChecksUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyPostureChecksUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServicePolicyPostureChecksTooManyRequestsCode is the HTTP code returned for type ListServicePolicyPostureChecksTooManyRequests
const ListServicePolicyPostureChecksTooManyRequestsCode int = 429

/*ListServicePolicyPostureChecksTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listServicePolicyPostureChecksTooManyRequests
*/
type ListServicePolicyPostureChecksTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServicePolicyPostureChecksTooManyRequests creates ListServicePolicyPostureChecksTooManyRequests with default headers values
func NewListServicePolicyPostureChecksTooManyRequests() *ListServicePolicyPostureChecksTooManyRequests {

	return &ListServicePolicyPostureChecksTooManyRequests{}
}

// WithPayload adds the payload to the list service policy posture checks too many requests response
func (o *ListServicePolicyPostureChecksTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServicePolicyPostureChecksTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service policy posture checks too many requests response
func (o *ListServicePolicyPostureChecksTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicePolicyPostureChecksTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
