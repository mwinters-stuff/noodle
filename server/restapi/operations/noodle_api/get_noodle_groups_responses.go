// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleGroupsOKCode is the HTTP code returned for type GetNoodleGroupsOK
const GetNoodleGroupsOKCode int = 200

/*
GetNoodleGroupsOK OK

swagger:response getNoodleGroupsOK
*/
type GetNoodleGroupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Group `json:"body,omitempty"`
}

// NewGetNoodleGroupsOK creates GetNoodleGroupsOK with default headers values
func NewGetNoodleGroupsOK() *GetNoodleGroupsOK {

	return &GetNoodleGroupsOK{}
}

// WithPayload adds the payload to the get noodle groups o k response
func (o *GetNoodleGroupsOK) WithPayload(payload []*models.Group) *GetNoodleGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle groups o k response
func (o *GetNoodleGroupsOK) SetPayload(payload []*models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Group, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetNoodleGroupsUnauthorizedCode is the HTTP code returned for type GetNoodleGroupsUnauthorized
const GetNoodleGroupsUnauthorizedCode int = 401

/*
GetNoodleGroupsUnauthorized unauthorized

swagger:response getNoodleGroupsUnauthorized
*/
type GetNoodleGroupsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleGroupsUnauthorized creates GetNoodleGroupsUnauthorized with default headers values
func NewGetNoodleGroupsUnauthorized() *GetNoodleGroupsUnauthorized {

	return &GetNoodleGroupsUnauthorized{}
}

// WithPayload adds the payload to the get noodle groups unauthorized response
func (o *GetNoodleGroupsUnauthorized) WithPayload(payload *models.Error) *GetNoodleGroupsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle groups unauthorized response
func (o *GetNoodleGroupsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleGroupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoodleGroupsConflictCode is the HTTP code returned for type GetNoodleGroupsConflict
const GetNoodleGroupsConflictCode int = 409

/*
GetNoodleGroupsConflict Failed

swagger:response getNoodleGroupsConflict
*/
type GetNoodleGroupsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleGroupsConflict creates GetNoodleGroupsConflict with default headers values
func NewGetNoodleGroupsConflict() *GetNoodleGroupsConflict {

	return &GetNoodleGroupsConflict{}
}

// WithPayload adds the payload to the get noodle groups conflict response
func (o *GetNoodleGroupsConflict) WithPayload(payload *models.Error) *GetNoodleGroupsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle groups conflict response
func (o *GetNoodleGroupsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleGroupsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
