// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPolicySelectorsHandlerFunc turns a function with the right signature into a get policy selectors handler
type GetPolicySelectorsHandlerFunc func(GetPolicySelectorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPolicySelectorsHandlerFunc) Handle(params GetPolicySelectorsParams) middleware.Responder {
	return fn(params)
}

// GetPolicySelectorsHandler interface for that can handle valid get policy selectors params
type GetPolicySelectorsHandler interface {
	Handle(GetPolicySelectorsParams) middleware.Responder
}

// NewGetPolicySelectors creates a new http.Handler for the get policy selectors operation
func NewGetPolicySelectors(ctx *middleware.Context, handler GetPolicySelectorsHandler) *GetPolicySelectors {
	return &GetPolicySelectors{Context: ctx, Handler: handler}
}

/*GetPolicySelectors swagger:route GET /policy/selectors policy getPolicySelectors

See what selectors match which identities

*/
type GetPolicySelectors struct {
	Context *middleware.Context
	Handler GetPolicySelectorsHandler
}

func (o *GetPolicySelectors) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPolicySelectorsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
