// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// PostNoodleApplicationsHandlerFunc turns a function with the right signature into a post noodle applications handler
type PostNoodleApplicationsHandlerFunc func(PostNoodleApplicationsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostNoodleApplicationsHandlerFunc) Handle(params PostNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostNoodleApplicationsHandler interface for that can handle valid post noodle applications params
type PostNoodleApplicationsHandler interface {
	Handle(PostNoodleApplicationsParams, *models.Principal) middleware.Responder
}

// NewPostNoodleApplications creates a new http.Handler for the post noodle applications operation
func NewPostNoodleApplications(ctx *middleware.Context, handler PostNoodleApplicationsHandler) *PostNoodleApplications {
	return &PostNoodleApplications{Context: ctx, Handler: handler}
}

/*
	PostNoodleApplications swagger:route POST /noodle/applications noodle-api postNoodleApplications

Adds a new application
*/
type PostNoodleApplications struct {
	Context *middleware.Context
	Handler PostNoodleApplicationsHandler
}

func (o *PostNoodleApplications) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostNoodleApplicationsParams()
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
