// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUserApplicationsHandlerFunc turns a function with the right signature into a get noodle user applications handler
type GetNoodleUserApplicationsHandlerFunc func(GetNoodleUserApplicationsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNoodleUserApplicationsHandlerFunc) Handle(params GetNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetNoodleUserApplicationsHandler interface for that can handle valid get noodle user applications params
type GetNoodleUserApplicationsHandler interface {
	Handle(GetNoodleUserApplicationsParams, *models.Principal) middleware.Responder
}

// NewGetNoodleUserApplications creates a new http.Handler for the get noodle user applications operation
func NewGetNoodleUserApplications(ctx *middleware.Context, handler GetNoodleUserApplicationsHandler) *GetNoodleUserApplications {
	return &GetNoodleUserApplications{Context: ctx, Handler: handler}
}

/*
	GetNoodleUserApplications swagger:route GET /noodle/user-applications noodle-api getNoodleUserApplications

Gets the list of user applications
*/
type GetNoodleUserApplications struct {
	Context *middleware.Context
	Handler GetNoodleUserApplicationsHandler
}

func (o *GetNoodleUserApplications) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetNoodleUserApplicationsParams()
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
