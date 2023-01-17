// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleGroupApplicationsOKCode is the HTTP code returned for type GetNoodleGroupApplicationsOK
const GetNoodleGroupApplicationsOKCode int = 200

/*
GetNoodleGroupApplicationsOK OK

swagger:response getNoodleGroupApplicationsOK
*/
type GetNoodleGroupApplicationsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.GroupApplications `json:"body,omitempty"`
}

// NewGetNoodleGroupApplicationsOK creates GetNoodleGroupApplicationsOK with default headers values
func NewGetNoodleGroupApplicationsOK() *GetNoodleGroupApplicationsOK {

	return &GetNoodleGroupApplicationsOK{}
}

// WithPayload adds the payload to the get noodle group applications o k response
func (o *GetNoodleGroupApplicationsOK) WithPayload(payload []*models.GroupApplications) *GetNoodleGroupApplicationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle group applications o k response
func (o *GetNoodleGroupApplicationsOK) SetPayload(payload []*models.GroupApplications) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleGroupApplicationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.GroupApplications, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetNoodleGroupApplicationsUnauthorizedCode is the HTTP code returned for type GetNoodleGroupApplicationsUnauthorized
const GetNoodleGroupApplicationsUnauthorizedCode int = 401

/*
GetNoodleGroupApplicationsUnauthorized unauthorized

swagger:response getNoodleGroupApplicationsUnauthorized
*/
type GetNoodleGroupApplicationsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleGroupApplicationsUnauthorized creates GetNoodleGroupApplicationsUnauthorized with default headers values
func NewGetNoodleGroupApplicationsUnauthorized() *GetNoodleGroupApplicationsUnauthorized {

	return &GetNoodleGroupApplicationsUnauthorized{}
}

// WithPayload adds the payload to the get noodle group applications unauthorized response
func (o *GetNoodleGroupApplicationsUnauthorized) WithPayload(payload *models.Error) *GetNoodleGroupApplicationsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle group applications unauthorized response
func (o *GetNoodleGroupApplicationsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleGroupApplicationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoodleGroupApplicationsConflictCode is the HTTP code returned for type GetNoodleGroupApplicationsConflict
const GetNoodleGroupApplicationsConflictCode int = 409

/*
GetNoodleGroupApplicationsConflict Failed

swagger:response getNoodleGroupApplicationsConflict
*/
type GetNoodleGroupApplicationsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleGroupApplicationsConflict creates GetNoodleGroupApplicationsConflict with default headers values
func NewGetNoodleGroupApplicationsConflict() *GetNoodleGroupApplicationsConflict {

	return &GetNoodleGroupApplicationsConflict{}
}

// WithPayload adds the payload to the get noodle group applications conflict response
func (o *GetNoodleGroupApplicationsConflict) WithPayload(payload *models.Error) *GetNoodleGroupApplicationsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle group applications conflict response
func (o *GetNoodleGroupApplicationsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleGroupApplicationsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
