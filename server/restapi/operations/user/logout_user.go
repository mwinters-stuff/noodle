// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LogoutUserHandlerFunc turns a function with the right signature into a logout user handler
type LogoutUserHandlerFunc func(LogoutUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LogoutUserHandlerFunc) Handle(params LogoutUserParams) middleware.Responder {
	return fn(params)
}

// LogoutUserHandler interface for that can handle valid logout user params
type LogoutUserHandler interface {
	Handle(LogoutUserParams) middleware.Responder
}

// NewLogoutUser creates a new http.Handler for the logout user operation
func NewLogoutUser(ctx *middleware.Context, handler LogoutUserHandler) *LogoutUser {
	return &LogoutUser{Context: ctx, Handler: handler}
}

/*
	LogoutUser swagger:route GET /noodle/logout user logoutUser

Logs out current logged in user session
*/
type LogoutUser struct {
	Context *middleware.Context
	Handler LogoutUserHandler
}

func (o *LogoutUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLogoutUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
