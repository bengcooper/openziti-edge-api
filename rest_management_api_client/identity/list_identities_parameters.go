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
	"github.com/go-openapi/swag"
)

// NewListIdentitiesParams creates a new ListIdentitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListIdentitiesParams() *ListIdentitiesParams {
	return &ListIdentitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListIdentitiesParamsWithTimeout creates a new ListIdentitiesParams object
// with the ability to set a timeout on a request.
func NewListIdentitiesParamsWithTimeout(timeout time.Duration) *ListIdentitiesParams {
	return &ListIdentitiesParams{
		timeout: timeout,
	}
}

// NewListIdentitiesParamsWithContext creates a new ListIdentitiesParams object
// with the ability to set a context for a request.
func NewListIdentitiesParamsWithContext(ctx context.Context) *ListIdentitiesParams {
	return &ListIdentitiesParams{
		Context: ctx,
	}
}

// NewListIdentitiesParamsWithHTTPClient creates a new ListIdentitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListIdentitiesParamsWithHTTPClient(client *http.Client) *ListIdentitiesParams {
	return &ListIdentitiesParams{
		HTTPClient: client,
	}
}

/*
ListIdentitiesParams contains all the parameters to send to the API endpoint

	for the list identities operation.

	Typically these are written to a http.Request.
*/
type ListIdentitiesParams struct {

	// Filter.
	Filter *string

	// Limit.
	Limit *int64

	// Offset.
	Offset *int64

	// RoleFilter.
	RoleFilter []string

	// RoleSemantic.
	RoleSemantic *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIdentitiesParams) WithDefaults() *ListIdentitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIdentitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list identities params
func (o *ListIdentitiesParams) WithTimeout(timeout time.Duration) *ListIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list identities params
func (o *ListIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list identities params
func (o *ListIdentitiesParams) WithContext(ctx context.Context) *ListIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list identities params
func (o *ListIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list identities params
func (o *ListIdentitiesParams) WithHTTPClient(client *http.Client) *ListIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list identities params
func (o *ListIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the list identities params
func (o *ListIdentitiesParams) WithFilter(filter *string) *ListIdentitiesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list identities params
func (o *ListIdentitiesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the list identities params
func (o *ListIdentitiesParams) WithLimit(limit *int64) *ListIdentitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list identities params
func (o *ListIdentitiesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list identities params
func (o *ListIdentitiesParams) WithOffset(offset *int64) *ListIdentitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list identities params
func (o *ListIdentitiesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRoleFilter adds the roleFilter to the list identities params
func (o *ListIdentitiesParams) WithRoleFilter(roleFilter []string) *ListIdentitiesParams {
	o.SetRoleFilter(roleFilter)
	return o
}

// SetRoleFilter adds the roleFilter to the list identities params
func (o *ListIdentitiesParams) SetRoleFilter(roleFilter []string) {
	o.RoleFilter = roleFilter
}

// WithRoleSemantic adds the roleSemantic to the list identities params
func (o *ListIdentitiesParams) WithRoleSemantic(roleSemantic *string) *ListIdentitiesParams {
	o.SetRoleSemantic(roleSemantic)
	return o
}

// SetRoleSemantic adds the roleSemantic to the list identities params
func (o *ListIdentitiesParams) SetRoleSemantic(roleSemantic *string) {
	o.RoleSemantic = roleSemantic
}

// WriteToRequest writes these params to a swagger request
func (o *ListIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.RoleFilter != nil {

		// binding items for roleFilter
		joinedRoleFilter := o.bindParamRoleFilter(reg)

		// query array param roleFilter
		if err := r.SetQueryParam("roleFilter", joinedRoleFilter...); err != nil {
			return err
		}
	}

	if o.RoleSemantic != nil {

		// query param roleSemantic
		var qrRoleSemantic string

		if o.RoleSemantic != nil {
			qrRoleSemantic = *o.RoleSemantic
		}
		qRoleSemantic := qrRoleSemantic
		if qRoleSemantic != "" {

			if err := r.SetQueryParam("roleSemantic", qRoleSemantic); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListIdentities binds the parameter roleFilter
func (o *ListIdentitiesParams) bindParamRoleFilter(formats strfmt.Registry) []string {
	roleFilterIR := o.RoleFilter

	var roleFilterIC []string
	for _, roleFilterIIR := range roleFilterIR { // explode []string

		roleFilterIIV := roleFilterIIR // string as string
		roleFilterIC = append(roleFilterIC, roleFilterIIV)
	}

	// items.CollectionFormat: "multi"
	roleFilterIS := swag.JoinByFormat(roleFilterIC, "multi")

	return roleFilterIS
}
