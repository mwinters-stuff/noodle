// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/mwinters-stuff/noodle/server/models"
)

// NewPostNoodleUserApplicationsParams creates a new PostNoodleUserApplicationsParams object
//
// There are no default values defined in the spec.
func NewPostNoodleUserApplicationsParams() PostNoodleUserApplicationsParams {

	return PostNoodleUserApplicationsParams{}
}

// PostNoodleUserApplicationsParams contains all the bound params for the post noodle user applications operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostNoodleUserApplications
type PostNoodleUserApplicationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	UserApplication *models.UserApplications
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostNoodleUserApplicationsParams() beforehand.
func (o *PostNoodleUserApplicationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserApplications
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("userApplication", "body", ""))
			} else {
				res = append(res, errors.NewParseError("userApplication", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UserApplication = &body
			}
		}
	} else {
		res = append(res, errors.Required("userApplication", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
