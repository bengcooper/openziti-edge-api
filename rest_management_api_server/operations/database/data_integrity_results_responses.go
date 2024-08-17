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

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge-api/rest_model"
)

// DataIntegrityResultsOKCode is the HTTP code returned for type DataIntegrityResultsOK
const DataIntegrityResultsOKCode int = 200

/*
DataIntegrityResultsOK A list of data integrity issues found

swagger:response dataIntegrityResultsOK
*/
type DataIntegrityResultsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DataIntegrityCheckResultEnvelope `json:"body,omitempty"`
}

// NewDataIntegrityResultsOK creates DataIntegrityResultsOK with default headers values
func NewDataIntegrityResultsOK() *DataIntegrityResultsOK {

	return &DataIntegrityResultsOK{}
}

// WithPayload adds the payload to the data integrity results o k response
func (o *DataIntegrityResultsOK) WithPayload(payload *rest_model.DataIntegrityCheckResultEnvelope) *DataIntegrityResultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the data integrity results o k response
func (o *DataIntegrityResultsOK) SetPayload(payload *rest_model.DataIntegrityCheckResultEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DataIntegrityResultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DataIntegrityResultsUnauthorizedCode is the HTTP code returned for type DataIntegrityResultsUnauthorized
const DataIntegrityResultsUnauthorizedCode int = 401

/*
DataIntegrityResultsUnauthorized The supplied session does not have the correct access rights to request this resource

swagger:response dataIntegrityResultsUnauthorized
*/
type DataIntegrityResultsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDataIntegrityResultsUnauthorized creates DataIntegrityResultsUnauthorized with default headers values
func NewDataIntegrityResultsUnauthorized() *DataIntegrityResultsUnauthorized {

	return &DataIntegrityResultsUnauthorized{}
}

// WithPayload adds the payload to the data integrity results unauthorized response
func (o *DataIntegrityResultsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DataIntegrityResultsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the data integrity results unauthorized response
func (o *DataIntegrityResultsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DataIntegrityResultsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DataIntegrityResultsTooManyRequestsCode is the HTTP code returned for type DataIntegrityResultsTooManyRequests
const DataIntegrityResultsTooManyRequestsCode int = 429

/*
DataIntegrityResultsTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response dataIntegrityResultsTooManyRequests
*/
type DataIntegrityResultsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDataIntegrityResultsTooManyRequests creates DataIntegrityResultsTooManyRequests with default headers values
func NewDataIntegrityResultsTooManyRequests() *DataIntegrityResultsTooManyRequests {

	return &DataIntegrityResultsTooManyRequests{}
}

// WithPayload adds the payload to the data integrity results too many requests response
func (o *DataIntegrityResultsTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DataIntegrityResultsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the data integrity results too many requests response
func (o *DataIntegrityResultsTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DataIntegrityResultsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
