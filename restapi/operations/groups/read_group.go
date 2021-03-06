// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReadGroupHandlerFunc turns a function with the right signature into a read group handler
type ReadGroupHandlerFunc func(ReadGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadGroupHandlerFunc) Handle(params ReadGroupParams) middleware.Responder {
	return fn(params)
}

// ReadGroupHandler interface for that can handle valid read group params
type ReadGroupHandler interface {
	Handle(ReadGroupParams) middleware.Responder
}

// NewReadGroup creates a new http.Handler for the read group operation
func NewReadGroup(ctx *middleware.Context, handler ReadGroupHandler) *ReadGroup {
	return &ReadGroup{Context: ctx, Handler: handler}
}

/*ReadGroup swagger:route GET /groups/{id} Groups readGroup

ReadGroup read group API

*/
type ReadGroup struct {
	Context *middleware.Context
	Handler ReadGroupHandler
}

func (o *ReadGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
