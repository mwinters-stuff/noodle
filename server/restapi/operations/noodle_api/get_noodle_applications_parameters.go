// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetNoodleApplicationsParams creates a new GetNoodleApplicationsParams object
//
// There are no default values defined in the spec.
func NewGetNoodleApplicationsParams() GetNoodleApplicationsParams {

	return GetNoodleApplicationsParams{}
}

// GetNoodleApplicationsParams contains all the bound params for the get noodle applications operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetNoodleApplications
type GetNoodleApplicationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	ApplicationID *int64
	/*
	  In: query
	*/
	ApplicationTemplate *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNoodleApplicationsParams() beforehand.
func (o *GetNoodleApplicationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qApplicationID, qhkApplicationID, _ := qs.GetOK("application_id")
	if err := o.bindApplicationID(qApplicationID, qhkApplicationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationTemplate, qhkApplicationTemplate, _ := qs.GetOK("application_template")
	if err := o.bindApplicationTemplate(qApplicationTemplate, qhkApplicationTemplate, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplicationID binds and validates parameter ApplicationID from query.
func (o *GetNoodleApplicationsParams) bindApplicationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("application_id", "query", "int64", raw)
	}
	o.ApplicationID = &value

	return nil
}

// bindApplicationTemplate binds and validates parameter ApplicationTemplate from query.
func (o *GetNoodleApplicationsParams) bindApplicationTemplate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationTemplate = &raw

	return nil
}
