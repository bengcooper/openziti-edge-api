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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AuthenticateHandlerFunc turns a function with the right signature into a authenticate handler
type AuthenticateHandlerFunc func(AuthenticateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthenticateHandlerFunc) Handle(params AuthenticateParams) middleware.Responder {
	return fn(params)
}

// AuthenticateHandler interface for that can handle valid authenticate params
type AuthenticateHandler interface {
	Handle(AuthenticateParams) middleware.Responder
}

// NewAuthenticate creates a new http.Handler for the authenticate operation
func NewAuthenticate(ctx *middleware.Context, handler AuthenticateHandler) *Authenticate {
	return &Authenticate{Context: ctx, Handler: handler}
}

/*
	Authenticate swagger:route POST /authenticate Authentication authenticate

# Authenticate via a method supplied via a query string parameter

Allowed authentication methods include "password", "cert", and "ext-jwt"
*/
type Authenticate struct {
	Context *middleware.Context
	Handler AuthenticateHandler
}

func (o *Authenticate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthenticateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
