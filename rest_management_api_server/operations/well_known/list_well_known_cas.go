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

package well_known

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListWellKnownCasHandlerFunc turns a function with the right signature into a list well known cas handler
type ListWellKnownCasHandlerFunc func(ListWellKnownCasParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListWellKnownCasHandlerFunc) Handle(params ListWellKnownCasParams) middleware.Responder {
	return fn(params)
}

// ListWellKnownCasHandler interface for that can handle valid list well known cas params
type ListWellKnownCasHandler interface {
	Handle(ListWellKnownCasParams) middleware.Responder
}

// NewListWellKnownCas creates a new http.Handler for the list well known cas operation
func NewListWellKnownCas(ctx *middleware.Context, handler ListWellKnownCasHandler) *ListWellKnownCas {
	return &ListWellKnownCas{Context: ctx, Handler: handler}
}

/*
	ListWellKnownCas swagger:route GET /.well-known/est/cacerts Well Known listWellKnownCas

# Get CA Cert Store

This endpoint is used during enrollments to bootstrap trust between enrolling clients and the Ziti Edge API.
This endpoint returns a base64 encoded PKCS7 store. The content can be base64 decoded and parsed by any library
that supports parsing PKCS7 stores.
*/
type ListWellKnownCas struct {
	Context *middleware.Context
	Handler ListWellKnownCasHandler
}

func (o *ListWellKnownCas) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListWellKnownCasParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
