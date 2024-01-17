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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// ListCurrentIdentityAuthenticatorsOKCode is the HTTP code returned for type ListCurrentIdentityAuthenticatorsOK
const ListCurrentIdentityAuthenticatorsOKCode int = 200

/*ListCurrentIdentityAuthenticatorsOK A list of authenticators

swagger:response listCurrentIdentityAuthenticatorsOK
*/
type ListCurrentIdentityAuthenticatorsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListAuthenticatorsEnvelope `json:"body,omitempty"`
}

// NewListCurrentIdentityAuthenticatorsOK creates ListCurrentIdentityAuthenticatorsOK with default headers values
func NewListCurrentIdentityAuthenticatorsOK() *ListCurrentIdentityAuthenticatorsOK {

	return &ListCurrentIdentityAuthenticatorsOK{}
}

// WithPayload adds the payload to the list current identity authenticators o k response
func (o *ListCurrentIdentityAuthenticatorsOK) WithPayload(payload *rest_model.ListAuthenticatorsEnvelope) *ListCurrentIdentityAuthenticatorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list current identity authenticators o k response
func (o *ListCurrentIdentityAuthenticatorsOK) SetPayload(payload *rest_model.ListAuthenticatorsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCurrentIdentityAuthenticatorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCurrentIdentityAuthenticatorsBadRequestCode is the HTTP code returned for type ListCurrentIdentityAuthenticatorsBadRequest
const ListCurrentIdentityAuthenticatorsBadRequestCode int = 400

/*ListCurrentIdentityAuthenticatorsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listCurrentIdentityAuthenticatorsBadRequest
*/
type ListCurrentIdentityAuthenticatorsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListCurrentIdentityAuthenticatorsBadRequest creates ListCurrentIdentityAuthenticatorsBadRequest with default headers values
func NewListCurrentIdentityAuthenticatorsBadRequest() *ListCurrentIdentityAuthenticatorsBadRequest {

	return &ListCurrentIdentityAuthenticatorsBadRequest{}
}

// WithPayload adds the payload to the list current identity authenticators bad request response
func (o *ListCurrentIdentityAuthenticatorsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListCurrentIdentityAuthenticatorsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list current identity authenticators bad request response
func (o *ListCurrentIdentityAuthenticatorsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCurrentIdentityAuthenticatorsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCurrentIdentityAuthenticatorsUnauthorizedCode is the HTTP code returned for type ListCurrentIdentityAuthenticatorsUnauthorized
const ListCurrentIdentityAuthenticatorsUnauthorizedCode int = 401

/*ListCurrentIdentityAuthenticatorsUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response listCurrentIdentityAuthenticatorsUnauthorized
*/
type ListCurrentIdentityAuthenticatorsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListCurrentIdentityAuthenticatorsUnauthorized creates ListCurrentIdentityAuthenticatorsUnauthorized with default headers values
func NewListCurrentIdentityAuthenticatorsUnauthorized() *ListCurrentIdentityAuthenticatorsUnauthorized {

	return &ListCurrentIdentityAuthenticatorsUnauthorized{}
}

// WithPayload adds the payload to the list current identity authenticators unauthorized response
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListCurrentIdentityAuthenticatorsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list current identity authenticators unauthorized response
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCurrentIdentityAuthenticatorsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
