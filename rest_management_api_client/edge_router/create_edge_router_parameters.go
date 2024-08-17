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

package edge_router

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

// NewCreateEdgeRouterParams creates a new CreateEdgeRouterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEdgeRouterParams() *CreateEdgeRouterParams {
	return &CreateEdgeRouterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEdgeRouterParamsWithTimeout creates a new CreateEdgeRouterParams object
// with the ability to set a timeout on a request.
func NewCreateEdgeRouterParamsWithTimeout(timeout time.Duration) *CreateEdgeRouterParams {
	return &CreateEdgeRouterParams{
		timeout: timeout,
	}
}

// NewCreateEdgeRouterParamsWithContext creates a new CreateEdgeRouterParams object
// with the ability to set a context for a request.
func NewCreateEdgeRouterParamsWithContext(ctx context.Context) *CreateEdgeRouterParams {
	return &CreateEdgeRouterParams{
		Context: ctx,
	}
}

// NewCreateEdgeRouterParamsWithHTTPClient creates a new CreateEdgeRouterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEdgeRouterParamsWithHTTPClient(client *http.Client) *CreateEdgeRouterParams {
	return &CreateEdgeRouterParams{
		HTTPClient: client,
	}
}

/*
CreateEdgeRouterParams contains all the parameters to send to the API endpoint

	for the create edge router operation.

	Typically these are written to a http.Request.
*/
type CreateEdgeRouterParams struct {

	/* EdgeRouter.

	   A edge router to create
	*/
	EdgeRouter *rest_model.EdgeRouterCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create edge router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEdgeRouterParams) WithDefaults() *CreateEdgeRouterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create edge router params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEdgeRouterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create edge router params
func (o *CreateEdgeRouterParams) WithTimeout(timeout time.Duration) *CreateEdgeRouterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create edge router params
func (o *CreateEdgeRouterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create edge router params
func (o *CreateEdgeRouterParams) WithContext(ctx context.Context) *CreateEdgeRouterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create edge router params
func (o *CreateEdgeRouterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create edge router params
func (o *CreateEdgeRouterParams) WithHTTPClient(client *http.Client) *CreateEdgeRouterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create edge router params
func (o *CreateEdgeRouterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeRouter adds the edgeRouter to the create edge router params
func (o *CreateEdgeRouterParams) WithEdgeRouter(edgeRouter *rest_model.EdgeRouterCreate) *CreateEdgeRouterParams {
	o.SetEdgeRouter(edgeRouter)
	return o
}

// SetEdgeRouter adds the edgeRouter to the create edge router params
func (o *CreateEdgeRouterParams) SetEdgeRouter(edgeRouter *rest_model.EdgeRouterCreate) {
	o.EdgeRouter = edgeRouter
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEdgeRouterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.EdgeRouter != nil {
		if err := r.SetBodyParam(o.EdgeRouter); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
