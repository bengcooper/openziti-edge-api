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
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new posture checks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new posture checks API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new posture checks API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for posture checks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationJSONCharsetUTF8 sets the Accept header to "application/json; charset=utf-8".
func WithAcceptApplicationJSONCharsetUTF8(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json; charset=utf-8"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePostureCheck(params *CreatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePostureCheckCreated, error)

	DeletePostureCheck(params *DeletePostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePostureCheckOK, error)

	DetailPostureCheck(params *DetailPostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailPostureCheckOK, error)

	DetailPostureCheckType(params *DetailPostureCheckTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailPostureCheckTypeOK, error)

	ListPostureCheckTypes(params *ListPostureCheckTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostureCheckTypesOK, error)

	ListPostureChecks(params *ListPostureChecksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostureChecksOK, error)

	PatchPostureCheck(params *PatchPostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchPostureCheckOK, error)

	UpdatePostureCheck(params *UpdatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePostureCheckOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePostureCheck creates a posture checks

Creates a Posture Checks
*/
func (a *Client) CreatePostureCheck(params *CreatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePostureCheckCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostureCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPostureCheck",
		Method:             "POST",
		PathPattern:        "/posture-checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePostureCheckCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePostureCheck deletes an posture checks

Deletes and Posture Checks by id
*/
func (a *Client) DeletePostureCheck(params *DeletePostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePostureCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePostureCheck",
		Method:             "DELETE",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailPostureCheck retrieves a single posture checks

Retrieves a single Posture Checks by id
*/
func (a *Client) DetailPostureCheck(params *DetailPostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailPostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailPostureCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailPostureCheck",
		Method:             "GET",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailPostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetailPostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailPostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DetailPostureCheckType retrieves a single posture check type

Retrieves a single posture check type by id
*/
func (a *Client) DetailPostureCheckType(params *DetailPostureCheckTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetailPostureCheckTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailPostureCheckTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detailPostureCheckType",
		Method:             "GET",
		PathPattern:        "/posture-check-types/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailPostureCheckTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetailPostureCheckTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailPostureCheckType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListPostureCheckTypes lists a subset of posture check types

Retrieves a list of posture check types
*/
func (a *Client) ListPostureCheckTypes(params *ListPostureCheckTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostureCheckTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPostureCheckTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPostureCheckTypes",
		Method:             "GET",
		PathPattern:        "/posture-check-types",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPostureCheckTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPostureCheckTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPostureCheckTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListPostureChecks lists a subset of posture checks

Retrieves a list of posture checks
*/
func (a *Client) ListPostureChecks(params *ListPostureChecksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostureChecksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPostureChecksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPostureChecks",
		Method:             "GET",
		PathPattern:        "/posture-checks",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPostureChecksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPostureChecksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPostureChecks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchPostureCheck updates the supplied fields on a posture checks

Update only the supplied fields on a Posture Checks by id
*/
func (a *Client) PatchPostureCheck(params *PatchPostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchPostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPostureCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchPostureCheck",
		Method:             "PATCH",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchPostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchPostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePostureCheck updates all fields on a posture checks

Update all fields on a Posture Checks by id
*/
func (a *Client) UpdatePostureCheck(params *UpdatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePostureCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePostureCheck",
		Method:             "PUT",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
