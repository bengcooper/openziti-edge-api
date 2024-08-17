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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListIdentityEdgeRoutersHandlerFunc turns a function with the right signature into a list identity edge routers handler
type ListIdentityEdgeRoutersHandlerFunc func(ListIdentityEdgeRoutersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListIdentityEdgeRoutersHandlerFunc) Handle(params ListIdentityEdgeRoutersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListIdentityEdgeRoutersHandler interface for that can handle valid list identity edge routers params
type ListIdentityEdgeRoutersHandler interface {
	Handle(ListIdentityEdgeRoutersParams, interface{}) middleware.Responder
}

// NewListIdentityEdgeRouters creates a new http.Handler for the list identity edge routers operation
func NewListIdentityEdgeRouters(ctx *middleware.Context, handler ListIdentityEdgeRoutersHandler) *ListIdentityEdgeRouters {
	return &ListIdentityEdgeRouters{Context: ctx, Handler: handler}
}

/*
	ListIdentityEdgeRouters swagger:route GET /identities/{id}/edge-routers Identity listIdentityEdgeRouters

# List accessible edge-routers

Retrieves a list of edge-routers that the given identity may use to access services. Supports filtering, sorting, and pagination. Requires admin access.
*/
type ListIdentityEdgeRouters struct {
	Context *middleware.Context
	Handler ListIdentityEdgeRoutersHandler
}

func (o *ListIdentityEdgeRouters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListIdentityEdgeRoutersParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
