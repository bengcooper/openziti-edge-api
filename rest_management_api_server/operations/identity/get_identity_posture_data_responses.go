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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// GetIdentityPostureDataOKCode is the HTTP code returned for type GetIdentityPostureDataOK
const GetIdentityPostureDataOKCode int = 200

/*GetIdentityPostureDataOK Returns the document that represents posture data

swagger:response getIdentityPostureDataOK
*/
type GetIdentityPostureDataOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.PostureDataEnvelope `json:"body,omitempty"`
}

// NewGetIdentityPostureDataOK creates GetIdentityPostureDataOK with default headers values
func NewGetIdentityPostureDataOK() *GetIdentityPostureDataOK {

	return &GetIdentityPostureDataOK{}
}

// WithPayload adds the payload to the get identity posture data o k response
func (o *GetIdentityPostureDataOK) WithPayload(payload *rest_model.PostureDataEnvelope) *GetIdentityPostureDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity posture data o k response
func (o *GetIdentityPostureDataOK) SetPayload(payload *rest_model.PostureDataEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityPostureDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIdentityPostureDataUnauthorizedCode is the HTTP code returned for type GetIdentityPostureDataUnauthorized
const GetIdentityPostureDataUnauthorizedCode int = 401

/*GetIdentityPostureDataUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response getIdentityPostureDataUnauthorized
*/
type GetIdentityPostureDataUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetIdentityPostureDataUnauthorized creates GetIdentityPostureDataUnauthorized with default headers values
func NewGetIdentityPostureDataUnauthorized() *GetIdentityPostureDataUnauthorized {

	return &GetIdentityPostureDataUnauthorized{}
}

// WithPayload adds the payload to the get identity posture data unauthorized response
func (o *GetIdentityPostureDataUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *GetIdentityPostureDataUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity posture data unauthorized response
func (o *GetIdentityPostureDataUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityPostureDataUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIdentityPostureDataNotFoundCode is the HTTP code returned for type GetIdentityPostureDataNotFound
const GetIdentityPostureDataNotFoundCode int = 404

/*GetIdentityPostureDataNotFound The requested resource does not exist

swagger:response getIdentityPostureDataNotFound
*/
type GetIdentityPostureDataNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetIdentityPostureDataNotFound creates GetIdentityPostureDataNotFound with default headers values
func NewGetIdentityPostureDataNotFound() *GetIdentityPostureDataNotFound {

	return &GetIdentityPostureDataNotFound{}
}

// WithPayload adds the payload to the get identity posture data not found response
func (o *GetIdentityPostureDataNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *GetIdentityPostureDataNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity posture data not found response
func (o *GetIdentityPostureDataNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityPostureDataNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIdentityPostureDataTooManyRequestsCode is the HTTP code returned for type GetIdentityPostureDataTooManyRequests
const GetIdentityPostureDataTooManyRequestsCode int = 429

/*GetIdentityPostureDataTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response getIdentityPostureDataTooManyRequests
*/
type GetIdentityPostureDataTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetIdentityPostureDataTooManyRequests creates GetIdentityPostureDataTooManyRequests with default headers values
func NewGetIdentityPostureDataTooManyRequests() *GetIdentityPostureDataTooManyRequests {

	return &GetIdentityPostureDataTooManyRequests{}
}

// WithPayload adds the payload to the get identity posture data too many requests response
func (o *GetIdentityPostureDataTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *GetIdentityPostureDataTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity posture data too many requests response
func (o *GetIdentityPostureDataTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityPostureDataTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
