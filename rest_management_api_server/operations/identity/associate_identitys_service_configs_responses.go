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

// AssociateIdentitysServiceConfigsOKCode is the HTTP code returned for type AssociateIdentitysServiceConfigsOK
const AssociateIdentitysServiceConfigsOKCode int = 200

/*AssociateIdentitysServiceConfigsOK Base empty response

swagger:response associateIdentitysServiceConfigsOK
*/
type AssociateIdentitysServiceConfigsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewAssociateIdentitysServiceConfigsOK creates AssociateIdentitysServiceConfigsOK with default headers values
func NewAssociateIdentitysServiceConfigsOK() *AssociateIdentitysServiceConfigsOK {

	return &AssociateIdentitysServiceConfigsOK{}
}

// WithPayload adds the payload to the associate identitys service configs o k response
func (o *AssociateIdentitysServiceConfigsOK) WithPayload(payload *rest_model.Empty) *AssociateIdentitysServiceConfigsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate identitys service configs o k response
func (o *AssociateIdentitysServiceConfigsOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateIdentitysServiceConfigsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AssociateIdentitysServiceConfigsBadRequestCode is the HTTP code returned for type AssociateIdentitysServiceConfigsBadRequest
const AssociateIdentitysServiceConfigsBadRequestCode int = 400

/*AssociateIdentitysServiceConfigsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response associateIdentitysServiceConfigsBadRequest
*/
type AssociateIdentitysServiceConfigsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewAssociateIdentitysServiceConfigsBadRequest creates AssociateIdentitysServiceConfigsBadRequest with default headers values
func NewAssociateIdentitysServiceConfigsBadRequest() *AssociateIdentitysServiceConfigsBadRequest {

	return &AssociateIdentitysServiceConfigsBadRequest{}
}

// WithPayload adds the payload to the associate identitys service configs bad request response
func (o *AssociateIdentitysServiceConfigsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *AssociateIdentitysServiceConfigsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate identitys service configs bad request response
func (o *AssociateIdentitysServiceConfigsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateIdentitysServiceConfigsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AssociateIdentitysServiceConfigsUnauthorizedCode is the HTTP code returned for type AssociateIdentitysServiceConfigsUnauthorized
const AssociateIdentitysServiceConfigsUnauthorizedCode int = 401

/*AssociateIdentitysServiceConfigsUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response associateIdentitysServiceConfigsUnauthorized
*/
type AssociateIdentitysServiceConfigsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewAssociateIdentitysServiceConfigsUnauthorized creates AssociateIdentitysServiceConfigsUnauthorized with default headers values
func NewAssociateIdentitysServiceConfigsUnauthorized() *AssociateIdentitysServiceConfigsUnauthorized {

	return &AssociateIdentitysServiceConfigsUnauthorized{}
}

// WithPayload adds the payload to the associate identitys service configs unauthorized response
func (o *AssociateIdentitysServiceConfigsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *AssociateIdentitysServiceConfigsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate identitys service configs unauthorized response
func (o *AssociateIdentitysServiceConfigsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateIdentitysServiceConfigsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AssociateIdentitysServiceConfigsNotFoundCode is the HTTP code returned for type AssociateIdentitysServiceConfigsNotFound
const AssociateIdentitysServiceConfigsNotFoundCode int = 404

/*AssociateIdentitysServiceConfigsNotFound The requested resource does not exist

swagger:response associateIdentitysServiceConfigsNotFound
*/
type AssociateIdentitysServiceConfigsNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewAssociateIdentitysServiceConfigsNotFound creates AssociateIdentitysServiceConfigsNotFound with default headers values
func NewAssociateIdentitysServiceConfigsNotFound() *AssociateIdentitysServiceConfigsNotFound {

	return &AssociateIdentitysServiceConfigsNotFound{}
}

// WithPayload adds the payload to the associate identitys service configs not found response
func (o *AssociateIdentitysServiceConfigsNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *AssociateIdentitysServiceConfigsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate identitys service configs not found response
func (o *AssociateIdentitysServiceConfigsNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateIdentitysServiceConfigsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AssociateIdentitysServiceConfigsTooManyRequestsCode is the HTTP code returned for type AssociateIdentitysServiceConfigsTooManyRequests
const AssociateIdentitysServiceConfigsTooManyRequestsCode int = 429

/*AssociateIdentitysServiceConfigsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response associateIdentitysServiceConfigsTooManyRequests
*/
type AssociateIdentitysServiceConfigsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewAssociateIdentitysServiceConfigsTooManyRequests creates AssociateIdentitysServiceConfigsTooManyRequests with default headers values
func NewAssociateIdentitysServiceConfigsTooManyRequests() *AssociateIdentitysServiceConfigsTooManyRequests {

	return &AssociateIdentitysServiceConfigsTooManyRequests{}
}

// WithPayload adds the payload to the associate identitys service configs too many requests response
func (o *AssociateIdentitysServiceConfigsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *AssociateIdentitysServiceConfigsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate identitys service configs too many requests response
func (o *AssociateIdentitysServiceConfigsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateIdentitysServiceConfigsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
