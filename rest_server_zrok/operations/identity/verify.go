// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VerifyHandlerFunc turns a function with the right signature into a verify handler
type VerifyHandlerFunc func(VerifyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyHandlerFunc) Handle(params VerifyParams) middleware.Responder {
	return fn(params)
}

// VerifyHandler interface for that can handle valid verify params
type VerifyHandler interface {
	Handle(VerifyParams) middleware.Responder
}

// NewVerify creates a new http.Handler for the verify operation
func NewVerify(ctx *middleware.Context, handler VerifyHandler) *Verify {
	return &Verify{Context: ctx, Handler: handler}
}

/*
	Verify swagger:route GET /verify identity verify

Verify verify API
*/
type Verify struct {
	Context *middleware.Context
	Handler VerifyHandler
}

func (o *Verify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVerifyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
