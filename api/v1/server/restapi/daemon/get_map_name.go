// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMapNameHandlerFunc turns a function with the right signature into a get map name handler
type GetMapNameHandlerFunc func(GetMapNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMapNameHandlerFunc) Handle(params GetMapNameParams) middleware.Responder {
	return fn(params)
}

// GetMapNameHandler interface for that can handle valid get map name params
type GetMapNameHandler interface {
	Handle(GetMapNameParams) middleware.Responder
}

// NewGetMapName creates a new http.Handler for the get map name operation
func NewGetMapName(ctx *middleware.Context, handler GetMapNameHandler) *GetMapName {
	return &GetMapName{Context: ctx, Handler: handler}
}

/*GetMapName swagger:route GET /map/{name} daemon getMapName

Retrieve contents of BPF map

*/
type GetMapName struct {
	Context *middleware.Context
	Handler GetMapNameHandler
}

func (o *GetMapName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMapNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
