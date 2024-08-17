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

package service_policy

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
	"github.com/go-openapi/swag"
)

// NewListServicePolicyPostureChecksParams creates a new ListServicePolicyPostureChecksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServicePolicyPostureChecksParams() *ListServicePolicyPostureChecksParams {
	return &ListServicePolicyPostureChecksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServicePolicyPostureChecksParamsWithTimeout creates a new ListServicePolicyPostureChecksParams object
// with the ability to set a timeout on a request.
func NewListServicePolicyPostureChecksParamsWithTimeout(timeout time.Duration) *ListServicePolicyPostureChecksParams {
	return &ListServicePolicyPostureChecksParams{
		timeout: timeout,
	}
}

// NewListServicePolicyPostureChecksParamsWithContext creates a new ListServicePolicyPostureChecksParams object
// with the ability to set a context for a request.
func NewListServicePolicyPostureChecksParamsWithContext(ctx context.Context) *ListServicePolicyPostureChecksParams {
	return &ListServicePolicyPostureChecksParams{
		Context: ctx,
	}
}

// NewListServicePolicyPostureChecksParamsWithHTTPClient creates a new ListServicePolicyPostureChecksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServicePolicyPostureChecksParamsWithHTTPClient(client *http.Client) *ListServicePolicyPostureChecksParams {
	return &ListServicePolicyPostureChecksParams{
		HTTPClient: client,
	}
}

/*
ListServicePolicyPostureChecksParams contains all the parameters to send to the API endpoint

	for the list service policy posture checks operation.

	Typically these are written to a http.Request.
*/
type ListServicePolicyPostureChecksParams struct {

	// Filter.
	Filter *string

	/* ID.

	   The id of the requested resource
	*/
	ID string

	// Limit.
	Limit *int64

	// Offset.
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list service policy posture checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServicePolicyPostureChecksParams) WithDefaults() *ListServicePolicyPostureChecksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list service policy posture checks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServicePolicyPostureChecksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithTimeout(timeout time.Duration) *ListServicePolicyPostureChecksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithContext(ctx context.Context) *ListServicePolicyPostureChecksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithHTTPClient(client *http.Client) *ListServicePolicyPostureChecksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithFilter(filter *string) *ListServicePolicyPostureChecksParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithID(id string) *ListServicePolicyPostureChecksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithLimit(limit *int64) *ListServicePolicyPostureChecksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) WithOffset(offset *int64) *ListServicePolicyPostureChecksParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list service policy posture checks params
func (o *ListServicePolicyPostureChecksParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicePolicyPostureChecksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
