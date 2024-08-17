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

package posture_checks

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

// NewPatchPostureCheckParams creates a new PatchPostureCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchPostureCheckParams() *PatchPostureCheckParams {
	return &PatchPostureCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPostureCheckParamsWithTimeout creates a new PatchPostureCheckParams object
// with the ability to set a timeout on a request.
func NewPatchPostureCheckParamsWithTimeout(timeout time.Duration) *PatchPostureCheckParams {
	return &PatchPostureCheckParams{
		timeout: timeout,
	}
}

// NewPatchPostureCheckParamsWithContext creates a new PatchPostureCheckParams object
// with the ability to set a context for a request.
func NewPatchPostureCheckParamsWithContext(ctx context.Context) *PatchPostureCheckParams {
	return &PatchPostureCheckParams{
		Context: ctx,
	}
}

// NewPatchPostureCheckParamsWithHTTPClient creates a new PatchPostureCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchPostureCheckParamsWithHTTPClient(client *http.Client) *PatchPostureCheckParams {
	return &PatchPostureCheckParams{
		HTTPClient: client,
	}
}

/*
PatchPostureCheckParams contains all the parameters to send to the API endpoint

	for the patch posture check operation.

	Typically these are written to a http.Request.
*/
type PatchPostureCheckParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* PostureCheck.

	   A Posture Check patch object
	*/
	PostureCheck rest_model.PostureCheckPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch posture check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPostureCheckParams) WithDefaults() *PatchPostureCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch posture check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPostureCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch posture check params
func (o *PatchPostureCheckParams) WithTimeout(timeout time.Duration) *PatchPostureCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch posture check params
func (o *PatchPostureCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch posture check params
func (o *PatchPostureCheckParams) WithContext(ctx context.Context) *PatchPostureCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch posture check params
func (o *PatchPostureCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch posture check params
func (o *PatchPostureCheckParams) WithHTTPClient(client *http.Client) *PatchPostureCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch posture check params
func (o *PatchPostureCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch posture check params
func (o *PatchPostureCheckParams) WithID(id string) *PatchPostureCheckParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch posture check params
func (o *PatchPostureCheckParams) SetID(id string) {
	o.ID = id
}

// WithPostureCheck adds the postureCheck to the patch posture check params
func (o *PatchPostureCheckParams) WithPostureCheck(postureCheck rest_model.PostureCheckPatch) *PatchPostureCheckParams {
	o.SetPostureCheck(postureCheck)
	return o
}

// SetPostureCheck adds the postureCheck to the patch posture check params
func (o *PatchPostureCheckParams) SetPostureCheck(postureCheck rest_model.PostureCheckPatch) {
	o.PostureCheck = postureCheck
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPostureCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.PostureCheck); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
