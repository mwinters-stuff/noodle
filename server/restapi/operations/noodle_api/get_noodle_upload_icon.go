// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUploadIconHandlerFunc turns a function with the right signature into a get noodle upload icon handler
type GetNoodleUploadIconHandlerFunc func(GetNoodleUploadIconParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNoodleUploadIconHandlerFunc) Handle(params GetNoodleUploadIconParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetNoodleUploadIconHandler interface for that can handle valid get noodle upload icon params
type GetNoodleUploadIconHandler interface {
	Handle(GetNoodleUploadIconParams, *models.Principal) middleware.Responder
}

// NewGetNoodleUploadIcon creates a new http.Handler for the get noodle upload icon operation
func NewGetNoodleUploadIcon(ctx *middleware.Context, handler GetNoodleUploadIconHandler) *GetNoodleUploadIcon {
	return &GetNoodleUploadIcon{Context: ctx, Handler: handler}
}

/*
	GetNoodleUploadIcon swagger:route GET /noodle/upload-icon noodle-api getNoodleUploadIcon

Gets list of upload icon filenames
*/
type GetNoodleUploadIcon struct {
	Context *middleware.Context
	Handler GetNoodleUploadIconHandler
}

func (o *GetNoodleUploadIcon) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetNoodleUploadIconParams()
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
