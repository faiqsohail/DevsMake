// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProfilesHandlerFunc turns a function with the right signature into a get profiles handler
type GetProfilesHandlerFunc func(GetProfilesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProfilesHandlerFunc) Handle(params GetProfilesParams) middleware.Responder {
	return fn(params)
}

// GetProfilesHandler interface for that can handle valid get profiles params
type GetProfilesHandler interface {
	Handle(GetProfilesParams) middleware.Responder
}

// NewGetProfiles creates a new http.Handler for the get profiles operation
func NewGetProfiles(ctx *middleware.Context, handler GetProfilesHandler) *GetProfiles {
	return &GetProfiles{Context: ctx, Handler: handler}
}

/* GetProfiles swagger:route GET /profiles profile getProfiles

fetches all profiles sorting by criteria

*/
type GetProfiles struct {
	Context *middleware.Context
	Handler GetProfilesHandler
}

func (o *GetProfiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProfilesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
