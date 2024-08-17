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

// DetailMfaRecoveryCodesOKCode is the HTTP code returned for type DetailMfaRecoveryCodesOK
const DetailMfaRecoveryCodesOKCode int = 200

/*
DetailMfaRecoveryCodesOK The recovery codes of an MFA enrollment

swagger:response detailMfaRecoveryCodesOK
*/
type DetailMfaRecoveryCodesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailMfaRecoveryCodesEnvelope `json:"body,omitempty"`
}

// NewDetailMfaRecoveryCodesOK creates DetailMfaRecoveryCodesOK with default headers values
func NewDetailMfaRecoveryCodesOK() *DetailMfaRecoveryCodesOK {

	return &DetailMfaRecoveryCodesOK{}
}

// WithPayload adds the payload to the detail mfa recovery codes o k response
func (o *DetailMfaRecoveryCodesOK) WithPayload(payload *rest_model.DetailMfaRecoveryCodesEnvelope) *DetailMfaRecoveryCodesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail mfa recovery codes o k response
func (o *DetailMfaRecoveryCodesOK) SetPayload(payload *rest_model.DetailMfaRecoveryCodesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailMfaRecoveryCodesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailMfaRecoveryCodesUnauthorizedCode is the HTTP code returned for type DetailMfaRecoveryCodesUnauthorized
const DetailMfaRecoveryCodesUnauthorizedCode int = 401

/*
DetailMfaRecoveryCodesUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response detailMfaRecoveryCodesUnauthorized
*/
type DetailMfaRecoveryCodesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailMfaRecoveryCodesUnauthorized creates DetailMfaRecoveryCodesUnauthorized with default headers values
func NewDetailMfaRecoveryCodesUnauthorized() *DetailMfaRecoveryCodesUnauthorized {

	return &DetailMfaRecoveryCodesUnauthorized{}
}

// WithPayload adds the payload to the detail mfa recovery codes unauthorized response
func (o *DetailMfaRecoveryCodesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailMfaRecoveryCodesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail mfa recovery codes unauthorized response
func (o *DetailMfaRecoveryCodesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailMfaRecoveryCodesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailMfaRecoveryCodesNotFoundCode is the HTTP code returned for type DetailMfaRecoveryCodesNotFound
const DetailMfaRecoveryCodesNotFoundCode int = 404

/*
DetailMfaRecoveryCodesNotFound The requested resource does not exist

swagger:response detailMfaRecoveryCodesNotFound
*/
type DetailMfaRecoveryCodesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailMfaRecoveryCodesNotFound creates DetailMfaRecoveryCodesNotFound with default headers values
func NewDetailMfaRecoveryCodesNotFound() *DetailMfaRecoveryCodesNotFound {

	return &DetailMfaRecoveryCodesNotFound{}
}

// WithPayload adds the payload to the detail mfa recovery codes not found response
func (o *DetailMfaRecoveryCodesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailMfaRecoveryCodesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail mfa recovery codes not found response
func (o *DetailMfaRecoveryCodesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailMfaRecoveryCodesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
