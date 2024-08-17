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

package enroll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// EnrollErOttOKCode is the HTTP code returned for type EnrollErOttOK
const EnrollErOttOKCode int = 200

/*
EnrollErOttOK A response containing the edge routers signed certificates (server chain, server cert, CAs).

swagger:response enrollErOttOK
*/
type EnrollErOttOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.EnrollmentCertsEnvelope `json:"body,omitempty"`
}

// NewEnrollErOttOK creates EnrollErOttOK with default headers values
func NewEnrollErOttOK() *EnrollErOttOK {

	return &EnrollErOttOK{}
}

// WithPayload adds the payload to the enroll er ott o k response
func (o *EnrollErOttOK) WithPayload(payload *rest_model.EnrollmentCertsEnvelope) *EnrollErOttOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll er ott o k response
func (o *EnrollErOttOK) SetPayload(payload *rest_model.EnrollmentCertsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollErOttOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnrollErOttTooManyRequestsCode is the HTTP code returned for type EnrollErOttTooManyRequests
const EnrollErOttTooManyRequestsCode int = 429

/*
EnrollErOttTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response enrollErOttTooManyRequests
*/
type EnrollErOttTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewEnrollErOttTooManyRequests creates EnrollErOttTooManyRequests with default headers values
func NewEnrollErOttTooManyRequests() *EnrollErOttTooManyRequests {

	return &EnrollErOttTooManyRequests{}
}

// WithPayload adds the payload to the enroll er ott too many requests response
func (o *EnrollErOttTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *EnrollErOttTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enroll er ott too many requests response
func (o *EnrollErOttTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnrollErOttTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
