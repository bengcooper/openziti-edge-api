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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetIdentityPolicyAdviceParams creates a new GetIdentityPolicyAdviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentityPolicyAdviceParams() *GetIdentityPolicyAdviceParams {
	return &GetIdentityPolicyAdviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityPolicyAdviceParamsWithTimeout creates a new GetIdentityPolicyAdviceParams object
// with the ability to set a timeout on a request.
func NewGetIdentityPolicyAdviceParamsWithTimeout(timeout time.Duration) *GetIdentityPolicyAdviceParams {
	return &GetIdentityPolicyAdviceParams{
		timeout: timeout,
	}
}

// NewGetIdentityPolicyAdviceParamsWithContext creates a new GetIdentityPolicyAdviceParams object
// with the ability to set a context for a request.
func NewGetIdentityPolicyAdviceParamsWithContext(ctx context.Context) *GetIdentityPolicyAdviceParams {
	return &GetIdentityPolicyAdviceParams{
		Context: ctx,
	}
}

// NewGetIdentityPolicyAdviceParamsWithHTTPClient creates a new GetIdentityPolicyAdviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentityPolicyAdviceParamsWithHTTPClient(client *http.Client) *GetIdentityPolicyAdviceParams {
	return &GetIdentityPolicyAdviceParams{
		HTTPClient: client,
	}
}

/*
GetIdentityPolicyAdviceParams contains all the parameters to send to the API endpoint

	for the get identity policy advice operation.

	Typically these are written to a http.Request.
*/
type GetIdentityPolicyAdviceParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* ServiceID.

	   The id of a service
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identity policy advice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityPolicyAdviceParams) WithDefaults() *GetIdentityPolicyAdviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identity policy advice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityPolicyAdviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) WithTimeout(timeout time.Duration) *GetIdentityPolicyAdviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) WithContext(ctx context.Context) *GetIdentityPolicyAdviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) WithHTTPClient(client *http.Client) *GetIdentityPolicyAdviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) WithID(id string) *GetIdentityPolicyAdviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) SetID(id string) {
	o.ID = id
}

// WithServiceID adds the serviceID to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) WithServiceID(serviceID string) *GetIdentityPolicyAdviceParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the get identity policy advice params
func (o *GetIdentityPolicyAdviceParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityPolicyAdviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
