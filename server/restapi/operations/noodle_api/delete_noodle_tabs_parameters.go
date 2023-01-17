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
	"github.com/go-openapi/validate"
)

// NewDeleteNoodleTabsParams creates a new DeleteNoodleTabsParams object
//
// There are no default values defined in the spec.
func NewDeleteNoodleTabsParams() DeleteNoodleTabsParams {

	return DeleteNoodleTabsParams{}
}

// DeleteNoodleTabsParams contains all the bound params for the delete noodle tabs operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteNoodleTabs
type DeleteNoodleTabsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Tabid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteNoodleTabsParams() beforehand.
func (o *DeleteNoodleTabsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTabid, qhkTabid, _ := qs.GetOK("tabid")
	if err := o.bindTabid(qTabid, qhkTabid, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTabid binds and validates parameter Tabid from query.
func (o *DeleteNoodleTabsParams) bindTabid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("tabid", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("tabid", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("tabid", "query", "int64", raw)
	}
	o.Tabid = value

	return nil
}
