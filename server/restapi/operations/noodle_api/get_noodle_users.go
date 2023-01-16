// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUsersHandlerFunc turns a function with the right signature into a get noodle users handler
type GetNoodleUsersHandlerFunc func(GetNoodleUsersParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNoodleUsersHandlerFunc) Handle(params GetNoodleUsersParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetNoodleUsersHandler interface for that can handle valid get noodle users params
type GetNoodleUsersHandler interface {
	Handle(GetNoodleUsersParams, *models.Principal) middleware.Responder
}

// NewGetNoodleUsers creates a new http.Handler for the get noodle users operation
func NewGetNoodleUsers(ctx *middleware.Context, handler GetNoodleUsersHandler) *GetNoodleUsers {
	return &GetNoodleUsers{Context: ctx, Handler: handler}
}

/*
	GetNoodleUsers swagger:route GET /noodle/users noodle-api getNoodleUsers

Gets the list of users or a single user
*/
type GetNoodleUsers struct {
	Context *middleware.Context
	Handler GetNoodleUsersHandler
}

func (o *GetNoodleUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetNoodleUsersParams()
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