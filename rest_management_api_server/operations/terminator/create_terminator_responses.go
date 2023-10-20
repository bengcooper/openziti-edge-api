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

package terminator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// CreateTerminatorCreatedCode is the HTTP code returned for type CreateTerminatorCreated
const CreateTerminatorCreatedCode int = 201

/*CreateTerminatorCreated The create request was successful and the resource has been added at the following location

swagger:response createTerminatorCreated
*/
type CreateTerminatorCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateTerminatorCreated creates CreateTerminatorCreated with default headers values
func NewCreateTerminatorCreated() *CreateTerminatorCreated {

	return &CreateTerminatorCreated{}
}

// WithPayload adds the payload to the create terminator created response
func (o *CreateTerminatorCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateTerminatorCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create terminator created response
func (o *CreateTerminatorCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTerminatorCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTerminatorBadRequestCode is the HTTP code returned for type CreateTerminatorBadRequest
const CreateTerminatorBadRequestCode int = 400

/*CreateTerminatorBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createTerminatorBadRequest
*/
type CreateTerminatorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateTerminatorBadRequest creates CreateTerminatorBadRequest with default headers values
func NewCreateTerminatorBadRequest() *CreateTerminatorBadRequest {

	return &CreateTerminatorBadRequest{}
}

// WithPayload adds the payload to the create terminator bad request response
func (o *CreateTerminatorBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateTerminatorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create terminator bad request response
func (o *CreateTerminatorBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTerminatorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTerminatorUnauthorizedCode is the HTTP code returned for type CreateTerminatorUnauthorized
const CreateTerminatorUnauthorizedCode int = 401

/*CreateTerminatorUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createTerminatorUnauthorized
*/
type CreateTerminatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateTerminatorUnauthorized creates CreateTerminatorUnauthorized with default headers values
func NewCreateTerminatorUnauthorized() *CreateTerminatorUnauthorized {

	return &CreateTerminatorUnauthorized{}
}

// WithPayload adds the payload to the create terminator unauthorized response
func (o *CreateTerminatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateTerminatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create terminator unauthorized response
func (o *CreateTerminatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTerminatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTerminatorTooManyRequestsCode is the HTTP code returned for type CreateTerminatorTooManyRequests
const CreateTerminatorTooManyRequestsCode int = 429

/*CreateTerminatorTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createTerminatorTooManyRequests
*/
type CreateTerminatorTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateTerminatorTooManyRequests creates CreateTerminatorTooManyRequests with default headers values
func NewCreateTerminatorTooManyRequests() *CreateTerminatorTooManyRequests {

	return &CreateTerminatorTooManyRequests{}
}

// WithPayload adds the payload to the create terminator too many requests response
func (o *CreateTerminatorTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateTerminatorTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create terminator too many requests response
func (o *CreateTerminatorTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTerminatorTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
