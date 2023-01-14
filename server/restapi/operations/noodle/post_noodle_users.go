// Code generated by go-swagger; DO NOT EDIT.

package noodle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// PostNoodleUsersHandlerFunc turns a function with the right signature into a post noodle users handler
type PostNoodleUsersHandlerFunc func(PostNoodleUsersParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostNoodleUsersHandlerFunc) Handle(params PostNoodleUsersParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostNoodleUsersHandler interface for that can handle valid post noodle users params
type PostNoodleUsersHandler interface {
	Handle(PostNoodleUsersParams, *models.Principal) middleware.Responder
}

// NewPostNoodleUsers creates a new http.Handler for the post noodle users operation
func NewPostNoodleUsers(ctx *middleware.Context, handler PostNoodleUsersHandler) *PostNoodleUsers {
	return &PostNoodleUsers{Context: ctx, Handler: handler}
}

/*
	PostNoodleUsers swagger:route POST /noodle/users noodle postNoodleUsers

Adds a new user
*/
type PostNoodleUsers struct {
	Context *middleware.Context
	Handler PostNoodleUsersHandler
}

func (o *PostNoodleUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostNoodleUsersParams()
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
