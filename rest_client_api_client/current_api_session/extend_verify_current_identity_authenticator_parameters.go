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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// NewExtendVerifyCurrentIdentityAuthenticatorParams creates a new ExtendVerifyCurrentIdentityAuthenticatorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtendVerifyCurrentIdentityAuthenticatorParams() *ExtendVerifyCurrentIdentityAuthenticatorParams {
	return &ExtendVerifyCurrentIdentityAuthenticatorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtendVerifyCurrentIdentityAuthenticatorParamsWithTimeout creates a new ExtendVerifyCurrentIdentityAuthenticatorParams object
// with the ability to set a timeout on a request.
func NewExtendVerifyCurrentIdentityAuthenticatorParamsWithTimeout(timeout time.Duration) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	return &ExtendVerifyCurrentIdentityAuthenticatorParams{
		timeout: timeout,
	}
}

// NewExtendVerifyCurrentIdentityAuthenticatorParamsWithContext creates a new ExtendVerifyCurrentIdentityAuthenticatorParams object
// with the ability to set a context for a request.
func NewExtendVerifyCurrentIdentityAuthenticatorParamsWithContext(ctx context.Context) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	return &ExtendVerifyCurrentIdentityAuthenticatorParams{
		Context: ctx,
	}
}

// NewExtendVerifyCurrentIdentityAuthenticatorParamsWithHTTPClient creates a new ExtendVerifyCurrentIdentityAuthenticatorParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtendVerifyCurrentIdentityAuthenticatorParamsWithHTTPClient(client *http.Client) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	return &ExtendVerifyCurrentIdentityAuthenticatorParams{
		HTTPClient: client,
	}
}

/*
ExtendVerifyCurrentIdentityAuthenticatorParams contains all the parameters to send to the API endpoint

	for the extend verify current identity authenticator operation.

	Typically these are written to a http.Request.
*/
type ExtendVerifyCurrentIdentityAuthenticatorParams struct {

	// Extend.
	Extend *rest_model.IdentityExtendValidateEnrollmentRequest

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extend verify current identity authenticator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WithDefaults() *ExtendVerifyCurrentIdentityAuthenticatorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extend verify current identity authenticator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WithTimeout(timeout time.Duration) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WithContext(ctx context.Context) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WithHTTPClient(client *http.Client) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtend adds the extend to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WithExtend(extend *rest_model.IdentityExtendValidateEnrollmentRequest) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	o.SetExtend(extend)
	return o
}

// SetExtend adds the extend to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) SetExtend(extend *rest_model.IdentityExtendValidateEnrollmentRequest) {
	o.Extend = extend
}

// WithID adds the id to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WithID(id string) *ExtendVerifyCurrentIdentityAuthenticatorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extend verify current identity authenticator params
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtendVerifyCurrentIdentityAuthenticatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Extend != nil {
		if err := r.SetBodyParam(o.Extend); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
