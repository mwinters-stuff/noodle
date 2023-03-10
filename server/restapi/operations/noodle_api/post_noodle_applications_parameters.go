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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/mwinters-stuff/noodle/server/models"
)

// NewPostNoodleApplicationsParams creates a new PostNoodleApplicationsParams object
//
// There are no default values defined in the spec.
func NewPostNoodleApplicationsParams() PostNoodleApplicationsParams {

	return PostNoodleApplicationsParams{}
}

// PostNoodleApplicationsParams contains all the bound params for the post noodle applications operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostNoodleApplications
type PostNoodleApplicationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Action string
	/*
	  Required: true
	  In: body
	*/
	Application *models.Application
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostNoodleApplicationsParams() beforehand.
func (o *PostNoodleApplicationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAction, qhkAction, _ := qs.GetOK("action")
	if err := o.bindAction(qAction, qhkAction, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Application
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("application", "body", ""))
			} else {
				res = append(res, errors.NewParseError("application", "body", "", err))
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
				o.Application = &body
			}
		}
	} else {
		res = append(res, errors.Required("application", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAction binds and validates parameter Action from query.
func (o *PostNoodleApplicationsParams) bindAction(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("action", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("action", "query", raw); err != nil {
		return err
	}
	o.Action = raw

	if err := o.validateAction(formats); err != nil {
		return err
	}

	return nil
}

// validateAction carries on validations for parameter Action
func (o *PostNoodleApplicationsParams) validateAction(formats strfmt.Registry) error {

	if err := validate.EnumCase("action", "query", o.Action, []interface{}{"insert", "update"}, true); err != nil {
		return err
	}

	return nil
}
