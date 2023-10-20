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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// CreatePostureResponseCreatedCode is the HTTP code returned for type CreatePostureResponseCreated
const CreatePostureResponseCreatedCode int = 201

/*CreatePostureResponseCreated Contains a list of services that have had their timers altered

swagger:response createPostureResponseCreated
*/
type CreatePostureResponseCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.PostureResponseEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseCreated creates CreatePostureResponseCreated with default headers values
func NewCreatePostureResponseCreated() *CreatePostureResponseCreated {

	return &CreatePostureResponseCreated{}
}

// WithPayload adds the payload to the create posture response created response
func (o *CreatePostureResponseCreated) WithPayload(payload *rest_model.PostureResponseEnvelope) *CreatePostureResponseCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response created response
func (o *CreatePostureResponseCreated) SetPayload(payload *rest_model.PostureResponseEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePostureResponseBadRequestCode is the HTTP code returned for type CreatePostureResponseBadRequest
const CreatePostureResponseBadRequestCode int = 400

/*CreatePostureResponseBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createPostureResponseBadRequest
*/
type CreatePostureResponseBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseBadRequest creates CreatePostureResponseBadRequest with default headers values
func NewCreatePostureResponseBadRequest() *CreatePostureResponseBadRequest {

	return &CreatePostureResponseBadRequest{}
}

// WithPayload adds the payload to the create posture response bad request response
func (o *CreatePostureResponseBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreatePostureResponseBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response bad request response
func (o *CreatePostureResponseBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePostureResponseUnauthorizedCode is the HTTP code returned for type CreatePostureResponseUnauthorized
const CreatePostureResponseUnauthorizedCode int = 401

/*CreatePostureResponseUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createPostureResponseUnauthorized
*/
type CreatePostureResponseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseUnauthorized creates CreatePostureResponseUnauthorized with default headers values
func NewCreatePostureResponseUnauthorized() *CreatePostureResponseUnauthorized {

	return &CreatePostureResponseUnauthorized{}
}

// WithPayload adds the payload to the create posture response unauthorized response
func (o *CreatePostureResponseUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreatePostureResponseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response unauthorized response
func (o *CreatePostureResponseUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePostureResponseTooManyRequestsCode is the HTTP code returned for type CreatePostureResponseTooManyRequests
const CreatePostureResponseTooManyRequestsCode int = 429

/*CreatePostureResponseTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createPostureResponseTooManyRequests
*/
type CreatePostureResponseTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreatePostureResponseTooManyRequests creates CreatePostureResponseTooManyRequests with default headers values
func NewCreatePostureResponseTooManyRequests() *CreatePostureResponseTooManyRequests {

	return &CreatePostureResponseTooManyRequests{}
}

// WithPayload adds the payload to the create posture response too many requests response
func (o *CreatePostureResponseTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreatePostureResponseTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create posture response too many requests response
func (o *CreatePostureResponseTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePostureResponseTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
