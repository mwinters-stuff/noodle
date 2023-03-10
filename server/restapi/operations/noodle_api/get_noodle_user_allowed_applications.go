// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUserAllowedApplicationsHandlerFunc turns a function with the right signature into a get noodle user allowed applications handler
type GetNoodleUserAllowedApplicationsHandlerFunc func(GetNoodleUserAllowedApplicationsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNoodleUserAllowedApplicationsHandlerFunc) Handle(params GetNoodleUserAllowedApplicationsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetNoodleUserAllowedApplicationsHandler interface for that can handle valid get noodle user allowed applications params
type GetNoodleUserAllowedApplicationsHandler interface {
	Handle(GetNoodleUserAllowedApplicationsParams, *models.Principal) middleware.Responder
}

// NewGetNoodleUserAllowedApplications creates a new http.Handler for the get noodle user allowed applications operation
func NewGetNoodleUserAllowedApplications(ctx *middleware.Context, handler GetNoodleUserAllowedApplicationsHandler) *GetNoodleUserAllowedApplications {
	return &GetNoodleUserAllowedApplications{Context: ctx, Handler: handler}
}

/*
	GetNoodleUserAllowedApplications swagger:route GET /noodle/user-allowed-applications noodle-api getNoodleUserAllowedApplications

Gets the list of the applications the user can see
*/
type GetNoodleUserAllowedApplications struct {
	Context *middleware.Context
	Handler GetNoodleUserAllowedApplicationsHandler
}

func (o *GetNoodleUserAllowedApplications) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetNoodleUserAllowedApplicationsParams()
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
