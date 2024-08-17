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

package service

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

// NewListServiceIdentitiesParams creates a new ListServiceIdentitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServiceIdentitiesParams() *ListServiceIdentitiesParams {
	return &ListServiceIdentitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceIdentitiesParamsWithTimeout creates a new ListServiceIdentitiesParams object
// with the ability to set a timeout on a request.
func NewListServiceIdentitiesParamsWithTimeout(timeout time.Duration) *ListServiceIdentitiesParams {
	return &ListServiceIdentitiesParams{
		timeout: timeout,
	}
}

// NewListServiceIdentitiesParamsWithContext creates a new ListServiceIdentitiesParams object
// with the ability to set a context for a request.
func NewListServiceIdentitiesParamsWithContext(ctx context.Context) *ListServiceIdentitiesParams {
	return &ListServiceIdentitiesParams{
		Context: ctx,
	}
}

// NewListServiceIdentitiesParamsWithHTTPClient creates a new ListServiceIdentitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServiceIdentitiesParamsWithHTTPClient(client *http.Client) *ListServiceIdentitiesParams {
	return &ListServiceIdentitiesParams{
		HTTPClient: client,
	}
}

/*
ListServiceIdentitiesParams contains all the parameters to send to the API endpoint

	for the list service identities operation.

	Typically these are written to a http.Request.
*/
type ListServiceIdentitiesParams struct {

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

	// PolicyType.
	PolicyType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list service identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceIdentitiesParams) WithDefaults() *ListServiceIdentitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list service identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceIdentitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list service identities params
func (o *ListServiceIdentitiesParams) WithTimeout(timeout time.Duration) *ListServiceIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service identities params
func (o *ListServiceIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service identities params
func (o *ListServiceIdentitiesParams) WithContext(ctx context.Context) *ListServiceIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service identities params
func (o *ListServiceIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service identities params
func (o *ListServiceIdentitiesParams) WithHTTPClient(client *http.Client) *ListServiceIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service identities params
func (o *ListServiceIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the list service identities params
func (o *ListServiceIdentitiesParams) WithFilter(filter *string) *ListServiceIdentitiesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list service identities params
func (o *ListServiceIdentitiesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the list service identities params
func (o *ListServiceIdentitiesParams) WithID(id string) *ListServiceIdentitiesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list service identities params
func (o *ListServiceIdentitiesParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the list service identities params
func (o *ListServiceIdentitiesParams) WithLimit(limit *int64) *ListServiceIdentitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list service identities params
func (o *ListServiceIdentitiesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list service identities params
func (o *ListServiceIdentitiesParams) WithOffset(offset *int64) *ListServiceIdentitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list service identities params
func (o *ListServiceIdentitiesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPolicyType adds the policyType to the list service identities params
func (o *ListServiceIdentitiesParams) WithPolicyType(policyType *string) *ListServiceIdentitiesParams {
	o.SetPolicyType(policyType)
	return o
}

// SetPolicyType adds the policyType to the list service identities params
func (o *ListServiceIdentitiesParams) SetPolicyType(policyType *string) {
	o.PolicyType = policyType
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PolicyType != nil {

		// query param policyType
		var qrPolicyType string

		if o.PolicyType != nil {
			qrPolicyType = *o.PolicyType
		}
		qPolicyType := qrPolicyType
		if qPolicyType != "" {

			if err := r.SetQueryParam("policyType", qPolicyType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
