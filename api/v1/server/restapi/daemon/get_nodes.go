// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNodesHandlerFunc turns a function with the right signature into a get nodes handler
type GetNodesHandlerFunc func(GetNodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNodesHandlerFunc) Handle(params GetNodesParams) middleware.Responder {
	return fn(params)
}

// GetNodesHandler interface for that can handle valid get nodes params
type GetNodesHandler interface {
	Handle(GetNodesParams) middleware.Responder
}

// NewGetNodes creates a new http.Handler for the get nodes operation
func NewGetNodes(ctx *middleware.Context, handler GetNodesHandler) *GetNodes {
	return &GetNodes{Context: ctx, Handler: handler}
}

/*GetNodes swagger:route GET /nodes daemon getNodes

Get nodes information stored in the cilium-agent

*/
type GetNodes struct {
	Context *middleware.Context
	Handler GetNodesHandler
}

func (o *GetNodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNodesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
