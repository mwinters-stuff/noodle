// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// PostNoodleUserApplicationsHandlerFunc turns a function with the right signature into a post noodle user applications handler
type PostNoodleUserApplicationsHandlerFunc func(PostNoodleUserApplicationsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostNoodleUserApplicationsHandlerFunc) Handle(params PostNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostNoodleUserApplicationsHandler interface for that can handle valid post noodle user applications params
type PostNoodleUserApplicationsHandler interface {
	Handle(PostNoodleUserApplicationsParams, *models.Principal) middleware.Responder
}

// NewPostNoodleUserApplications creates a new http.Handler for the post noodle user applications operation
func NewPostNoodleUserApplications(ctx *middleware.Context, handler PostNoodleUserApplicationsHandler) *PostNoodleUserApplications {
	return &PostNoodleUserApplications{Context: ctx, Handler: handler}
}

/*
	PostNoodleUserApplications swagger:route POST /noodle/user-applications noodle-api postNoodleUserApplications

Adds a new user application
*/
type PostNoodleUserApplications struct {
	Context *middleware.Context
	Handler PostNoodleUserApplicationsHandler
}

func (o *PostNoodleUserApplications) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostNoodleUserApplicationsParams()
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