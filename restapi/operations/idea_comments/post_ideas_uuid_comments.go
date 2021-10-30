// Code generated by go-swagger; DO NOT EDIT.

package idea_comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"devsmake/models"
)

// PostIdeasUUIDCommentsHandlerFunc turns a function with the right signature into a post ideas UUID comments handler
type PostIdeasUUIDCommentsHandlerFunc func(PostIdeasUUIDCommentsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIdeasUUIDCommentsHandlerFunc) Handle(params PostIdeasUUIDCommentsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostIdeasUUIDCommentsHandler interface for that can handle valid post ideas UUID comments params
type PostIdeasUUIDCommentsHandler interface {
	Handle(PostIdeasUUIDCommentsParams, *models.Principal) middleware.Responder
}

// NewPostIdeasUUIDComments creates a new http.Handler for the post ideas UUID comments operation
func NewPostIdeasUUIDComments(ctx *middleware.Context, handler PostIdeasUUIDCommentsHandler) *PostIdeasUUIDComments {
	return &PostIdeasUUIDComments{Context: ctx, Handler: handler}
}

/* PostIdeasUUIDComments swagger:route POST /ideas/{uuid}/comments idea comments postIdeasUuidComments

creates a comment for the idea post

*/
type PostIdeasUUIDComments struct {
	Context *middleware.Context
	Handler PostIdeasUUIDCommentsHandler
}

func (o *PostIdeasUUIDComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostIdeasUUIDCommentsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}