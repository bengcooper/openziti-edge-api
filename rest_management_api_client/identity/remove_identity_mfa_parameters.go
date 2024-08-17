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

// NewRemoveIdentityMfaParams creates a new RemoveIdentityMfaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveIdentityMfaParams() *RemoveIdentityMfaParams {
	return &RemoveIdentityMfaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveIdentityMfaParamsWithTimeout creates a new RemoveIdentityMfaParams object
// with the ability to set a timeout on a request.
func NewRemoveIdentityMfaParamsWithTimeout(timeout time.Duration) *RemoveIdentityMfaParams {
	return &RemoveIdentityMfaParams{
		timeout: timeout,
	}
}

// NewRemoveIdentityMfaParamsWithContext creates a new RemoveIdentityMfaParams object
// with the ability to set a context for a request.
func NewRemoveIdentityMfaParamsWithContext(ctx context.Context) *RemoveIdentityMfaParams {
	return &RemoveIdentityMfaParams{
		Context: ctx,
	}
}

// NewRemoveIdentityMfaParamsWithHTTPClient creates a new RemoveIdentityMfaParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveIdentityMfaParamsWithHTTPClient(client *http.Client) *RemoveIdentityMfaParams {
	return &RemoveIdentityMfaParams{
		HTTPClient: client,
	}
}

/*
RemoveIdentityMfaParams contains all the parameters to send to the API endpoint

	for the remove identity mfa operation.

	Typically these are written to a http.Request.
*/
type RemoveIdentityMfaParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove identity mfa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveIdentityMfaParams) WithDefaults() *RemoveIdentityMfaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove identity mfa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveIdentityMfaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove identity mfa params
func (o *RemoveIdentityMfaParams) WithTimeout(timeout time.Duration) *RemoveIdentityMfaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove identity mfa params
func (o *RemoveIdentityMfaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove identity mfa params
func (o *RemoveIdentityMfaParams) WithContext(ctx context.Context) *RemoveIdentityMfaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove identity mfa params
func (o *RemoveIdentityMfaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove identity mfa params
func (o *RemoveIdentityMfaParams) WithHTTPClient(client *http.Client) *RemoveIdentityMfaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove identity mfa params
func (o *RemoveIdentityMfaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove identity mfa params
func (o *RemoveIdentityMfaParams) WithID(id string) *RemoveIdentityMfaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove identity mfa params
func (o *RemoveIdentityMfaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveIdentityMfaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
