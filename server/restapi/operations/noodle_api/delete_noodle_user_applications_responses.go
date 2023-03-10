// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// DeleteNoodleUserApplicationsOKCode is the HTTP code returned for type DeleteNoodleUserApplicationsOK
const DeleteNoodleUserApplicationsOKCode int = 200

/*
DeleteNoodleUserApplicationsOK User Application Deleted.

swagger:response deleteNoodleUserApplicationsOK
*/
type DeleteNoodleUserApplicationsOK struct {
}

// NewDeleteNoodleUserApplicationsOK creates DeleteNoodleUserApplicationsOK with default headers values
func NewDeleteNoodleUserApplicationsOK() *DeleteNoodleUserApplicationsOK {

	return &DeleteNoodleUserApplicationsOK{}
}

// WriteResponse to the client
func (o *DeleteNoodleUserApplicationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteNoodleUserApplicationsUnauthorizedCode is the HTTP code returned for type DeleteNoodleUserApplicationsUnauthorized
const DeleteNoodleUserApplicationsUnauthorizedCode int = 401

/*
DeleteNoodleUserApplicationsUnauthorized unauthorized

swagger:response deleteNoodleUserApplicationsUnauthorized
*/
type DeleteNoodleUserApplicationsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNoodleUserApplicationsUnauthorized creates DeleteNoodleUserApplicationsUnauthorized with default headers values
func NewDeleteNoodleUserApplicationsUnauthorized() *DeleteNoodleUserApplicationsUnauthorized {

	return &DeleteNoodleUserApplicationsUnauthorized{}
}

// WithPayload adds the payload to the delete noodle user applications unauthorized response
func (o *DeleteNoodleUserApplicationsUnauthorized) WithPayload(payload *models.Error) *DeleteNoodleUserApplicationsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete noodle user applications unauthorized response
func (o *DeleteNoodleUserApplicationsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNoodleUserApplicationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNoodleUserApplicationsConflictCode is the HTTP code returned for type DeleteNoodleUserApplicationsConflict
const DeleteNoodleUserApplicationsConflictCode int = 409

/*
DeleteNoodleUserApplicationsConflict Invalid Input

swagger:response deleteNoodleUserApplicationsConflict
*/
type DeleteNoodleUserApplicationsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNoodleUserApplicationsConflict creates DeleteNoodleUserApplicationsConflict with default headers values
func NewDeleteNoodleUserApplicationsConflict() *DeleteNoodleUserApplicationsConflict {

	return &DeleteNoodleUserApplicationsConflict{}
}

// WithPayload adds the payload to the delete noodle user applications conflict response
func (o *DeleteNoodleUserApplicationsConflict) WithPayload(payload *models.Error) *DeleteNoodleUserApplicationsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete noodle user applications conflict response
func (o *DeleteNoodleUserApplicationsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNoodleUserApplicationsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
