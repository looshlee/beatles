// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFqdnNamesHandlerFunc turns a function with the right signature into a get fqdn names handler
type GetFqdnNamesHandlerFunc func(GetFqdnNamesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFqdnNamesHandlerFunc) Handle(params GetFqdnNamesParams) middleware.Responder {
	return fn(params)
}

// GetFqdnNamesHandler interface for that can handle valid get fqdn names params
type GetFqdnNamesHandler interface {
	Handle(GetFqdnNamesParams) middleware.Responder
}

// NewGetFqdnNames creates a new http.Handler for the get fqdn names operation
func NewGetFqdnNames(ctx *middleware.Context, handler GetFqdnNamesHandler) *GetFqdnNames {
	return &GetFqdnNames{Context: ctx, Handler: handler}
}

/*GetFqdnNames swagger:route GET /fqdn/names policy getFqdnNames

List internal DNS selector representations

Retrieves the list of DNS-related fields (names to poll, selectors and
their corresponding regexes).


*/
type GetFqdnNames struct {
	Context *middleware.Context
	Handler GetFqdnNamesHandler
}

func (o *GetFqdnNames) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFqdnNamesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
