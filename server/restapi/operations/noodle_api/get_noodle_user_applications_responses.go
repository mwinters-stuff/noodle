// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUserApplicationsOKCode is the HTTP code returned for type GetNoodleUserApplicationsOK
const GetNoodleUserApplicationsOKCode int = 200

/*
GetNoodleUserApplicationsOK OK

swagger:response getNoodleUserApplicationsOK
*/
type GetNoodleUserApplicationsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserApplications `json:"body,omitempty"`
}

// NewGetNoodleUserApplicationsOK creates GetNoodleUserApplicationsOK with default headers values
func NewGetNoodleUserApplicationsOK() *GetNoodleUserApplicationsOK {

	return &GetNoodleUserApplicationsOK{}
}

// WithPayload adds the payload to the get noodle user applications o k response
func (o *GetNoodleUserApplicationsOK) WithPayload(payload []*models.UserApplications) *GetNoodleUserApplicationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle user applications o k response
func (o *GetNoodleUserApplicationsOK) SetPayload(payload []*models.UserApplications) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUserApplicationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserApplications, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetNoodleUserApplicationsUnauthorizedCode is the HTTP code returned for type GetNoodleUserApplicationsUnauthorized
const GetNoodleUserApplicationsUnauthorizedCode int = 401

/*
GetNoodleUserApplicationsUnauthorized unauthorized

swagger:response getNoodleUserApplicationsUnauthorized
*/
type GetNoodleUserApplicationsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleUserApplicationsUnauthorized creates GetNoodleUserApplicationsUnauthorized with default headers values
func NewGetNoodleUserApplicationsUnauthorized() *GetNoodleUserApplicationsUnauthorized {

	return &GetNoodleUserApplicationsUnauthorized{}
}

// WithPayload adds the payload to the get noodle user applications unauthorized response
func (o *GetNoodleUserApplicationsUnauthorized) WithPayload(payload *models.Error) *GetNoodleUserApplicationsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle user applications unauthorized response
func (o *GetNoodleUserApplicationsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUserApplicationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoodleUserApplicationsConflictCode is the HTTP code returned for type GetNoodleUserApplicationsConflict
const GetNoodleUserApplicationsConflictCode int = 409

/*
GetNoodleUserApplicationsConflict Failed

swagger:response getNoodleUserApplicationsConflict
*/
type GetNoodleUserApplicationsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleUserApplicationsConflict creates GetNoodleUserApplicationsConflict with default headers values
func NewGetNoodleUserApplicationsConflict() *GetNoodleUserApplicationsConflict {

	return &GetNoodleUserApplicationsConflict{}
}

// WithPayload adds the payload to the get noodle user applications conflict response
func (o *GetNoodleUserApplicationsConflict) WithPayload(payload *models.Error) *GetNoodleUserApplicationsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle user applications conflict response
func (o *GetNoodleUserApplicationsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUserApplicationsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
