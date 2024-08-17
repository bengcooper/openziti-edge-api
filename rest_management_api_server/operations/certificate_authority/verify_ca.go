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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VerifyCaHandlerFunc turns a function with the right signature into a verify ca handler
type VerifyCaHandlerFunc func(VerifyCaParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyCaHandlerFunc) Handle(params VerifyCaParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// VerifyCaHandler interface for that can handle valid verify ca params
type VerifyCaHandler interface {
	Handle(VerifyCaParams, interface{}) middleware.Responder
}

// NewVerifyCa creates a new http.Handler for the verify ca operation
func NewVerifyCa(ctx *middleware.Context, handler VerifyCaHandler) *VerifyCa {
	return &VerifyCa{Context: ctx, Handler: handler}
}

/*
	VerifyCa swagger:route POST /cas/{id}/verify Certificate Authority verifyCa

# Verify a CA

Allows a CA to become verified by submitting a certificate in PEM format that has been signed by the target CA.
The common name on the certificate must match the verificationToken property of the CA. Unverfieid CAs can not
be used for enrollment/authentication. Requires admin access.
*/
type VerifyCa struct {
	Context *middleware.Context
	Handler VerifyCaHandler
}

func (o *VerifyCa) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVerifyCaParams()
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
