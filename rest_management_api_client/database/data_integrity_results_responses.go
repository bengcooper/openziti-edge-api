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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DataIntegrityResultsReader is a Reader for the DataIntegrityResults structure.
type DataIntegrityResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataIntegrityResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataIntegrityResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDataIntegrityResultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDataIntegrityResultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /database/data-integrity-results] dataIntegrityResults", response, response.Code())
	}
}

// NewDataIntegrityResultsOK creates a DataIntegrityResultsOK with default headers values
func NewDataIntegrityResultsOK() *DataIntegrityResultsOK {
	return &DataIntegrityResultsOK{}
}

/*
DataIntegrityResultsOK describes a response with status code 200, with default header values.

A list of data integrity issues found
*/
type DataIntegrityResultsOK struct {
	Payload *rest_model.DataIntegrityCheckResultEnvelope
}

// IsSuccess returns true when this data integrity results o k response has a 2xx status code
func (o *DataIntegrityResultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this data integrity results o k response has a 3xx status code
func (o *DataIntegrityResultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this data integrity results o k response has a 4xx status code
func (o *DataIntegrityResultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this data integrity results o k response has a 5xx status code
func (o *DataIntegrityResultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this data integrity results o k response a status code equal to that given
func (o *DataIntegrityResultsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the data integrity results o k response
func (o *DataIntegrityResultsOK) Code() int {
	return 200
}

func (o *DataIntegrityResultsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /database/data-integrity-results][%d] dataIntegrityResultsOK %s", 200, payload)
}

func (o *DataIntegrityResultsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /database/data-integrity-results][%d] dataIntegrityResultsOK %s", 200, payload)
}

func (o *DataIntegrityResultsOK) GetPayload() *rest_model.DataIntegrityCheckResultEnvelope {
	return o.Payload
}

func (o *DataIntegrityResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DataIntegrityCheckResultEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataIntegrityResultsUnauthorized creates a DataIntegrityResultsUnauthorized with default headers values
func NewDataIntegrityResultsUnauthorized() *DataIntegrityResultsUnauthorized {
	return &DataIntegrityResultsUnauthorized{}
}

/*
DataIntegrityResultsUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DataIntegrityResultsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this data integrity results unauthorized response has a 2xx status code
func (o *DataIntegrityResultsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this data integrity results unauthorized response has a 3xx status code
func (o *DataIntegrityResultsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this data integrity results unauthorized response has a 4xx status code
func (o *DataIntegrityResultsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this data integrity results unauthorized response has a 5xx status code
func (o *DataIntegrityResultsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this data integrity results unauthorized response a status code equal to that given
func (o *DataIntegrityResultsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the data integrity results unauthorized response
func (o *DataIntegrityResultsUnauthorized) Code() int {
	return 401
}

func (o *DataIntegrityResultsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /database/data-integrity-results][%d] dataIntegrityResultsUnauthorized %s", 401, payload)
}

func (o *DataIntegrityResultsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /database/data-integrity-results][%d] dataIntegrityResultsUnauthorized %s", 401, payload)
}

func (o *DataIntegrityResultsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DataIntegrityResultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataIntegrityResultsTooManyRequests creates a DataIntegrityResultsTooManyRequests with default headers values
func NewDataIntegrityResultsTooManyRequests() *DataIntegrityResultsTooManyRequests {
	return &DataIntegrityResultsTooManyRequests{}
}

/*
DataIntegrityResultsTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DataIntegrityResultsTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

// IsSuccess returns true when this data integrity results too many requests response has a 2xx status code
func (o *DataIntegrityResultsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this data integrity results too many requests response has a 3xx status code
func (o *DataIntegrityResultsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this data integrity results too many requests response has a 4xx status code
func (o *DataIntegrityResultsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this data integrity results too many requests response has a 5xx status code
func (o *DataIntegrityResultsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this data integrity results too many requests response a status code equal to that given
func (o *DataIntegrityResultsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the data integrity results too many requests response
func (o *DataIntegrityResultsTooManyRequests) Code() int {
	return 429
}

func (o *DataIntegrityResultsTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /database/data-integrity-results][%d] dataIntegrityResultsTooManyRequests %s", 429, payload)
}

func (o *DataIntegrityResultsTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /database/data-integrity-results][%d] dataIntegrityResultsTooManyRequests %s", 429, payload)
}

func (o *DataIntegrityResultsTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DataIntegrityResultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
