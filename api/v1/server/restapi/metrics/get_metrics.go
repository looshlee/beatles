// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMetricsHandlerFunc turns a function with the right signature into a get metrics handler
type GetMetricsHandlerFunc func(GetMetricsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMetricsHandlerFunc) Handle(params GetMetricsParams) middleware.Responder {
	return fn(params)
}

// GetMetricsHandler interface for that can handle valid get metrics params
type GetMetricsHandler interface {
	Handle(GetMetricsParams) middleware.Responder
}

// NewGetMetrics creates a new http.Handler for the get metrics operation
func NewGetMetrics(ctx *middleware.Context, handler GetMetricsHandler) *GetMetrics {
	return &GetMetrics{Context: ctx, Handler: handler}
}

/*GetMetrics swagger:route GET /metrics/ metrics getMetrics

Retrieve cilium metrics

*/
type GetMetrics struct {
	Context *middleware.Context
	Handler GetMetricsHandler
}

func (o *GetMetrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMetricsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
