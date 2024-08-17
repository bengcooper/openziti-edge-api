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

// GetEnrollmentJwksOKCode is the HTTP code returned for type GetEnrollmentJwksOK
const GetEnrollmentJwksOKCode int = 200

/*
GetEnrollmentJwksOK A JWKS response for enrollment.

swagger:response getEnrollmentJwksOK
*/
type GetEnrollmentJwksOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Jwks `json:"body,omitempty"`
}

// NewGetEnrollmentJwksOK creates GetEnrollmentJwksOK with default headers values
func NewGetEnrollmentJwksOK() *GetEnrollmentJwksOK {

	return &GetEnrollmentJwksOK{}
}

// WithPayload adds the payload to the get enrollment jwks o k response
func (o *GetEnrollmentJwksOK) WithPayload(payload *rest_model.Jwks) *GetEnrollmentJwksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get enrollment jwks o k response
func (o *GetEnrollmentJwksOK) SetPayload(payload *rest_model.Jwks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEnrollmentJwksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
