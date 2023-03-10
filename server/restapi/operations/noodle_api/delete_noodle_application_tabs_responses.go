// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// DeleteNoodleApplicationTabsOKCode is the HTTP code returned for type DeleteNoodleApplicationTabsOK
const DeleteNoodleApplicationTabsOKCode int = 200

/*
DeleteNoodleApplicationTabsOK Application Tab Deleted.

swagger:response deleteNoodleApplicationTabsOK
*/
type DeleteNoodleApplicationTabsOK struct {
}

// NewDeleteNoodleApplicationTabsOK creates DeleteNoodleApplicationTabsOK with default headers values
func NewDeleteNoodleApplicationTabsOK() *DeleteNoodleApplicationTabsOK {

	return &DeleteNoodleApplicationTabsOK{}
}

// WriteResponse to the client
func (o *DeleteNoodleApplicationTabsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteNoodleApplicationTabsUnauthorizedCode is the HTTP code returned for type DeleteNoodleApplicationTabsUnauthorized
const DeleteNoodleApplicationTabsUnauthorizedCode int = 401

/*
DeleteNoodleApplicationTabsUnauthorized unauthorized

swagger:response deleteNoodleApplicationTabsUnauthorized
*/
type DeleteNoodleApplicationTabsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNoodleApplicationTabsUnauthorized creates DeleteNoodleApplicationTabsUnauthorized with default headers values
func NewDeleteNoodleApplicationTabsUnauthorized() *DeleteNoodleApplicationTabsUnauthorized {

	return &DeleteNoodleApplicationTabsUnauthorized{}
}

// WithPayload adds the payload to the delete noodle application tabs unauthorized response
func (o *DeleteNoodleApplicationTabsUnauthorized) WithPayload(payload *models.Error) *DeleteNoodleApplicationTabsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete noodle application tabs unauthorized response
func (o *DeleteNoodleApplicationTabsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNoodleApplicationTabsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNoodleApplicationTabsConflictCode is the HTTP code returned for type DeleteNoodleApplicationTabsConflict
const DeleteNoodleApplicationTabsConflictCode int = 409

/*
DeleteNoodleApplicationTabsConflict Invalid Input

swagger:response deleteNoodleApplicationTabsConflict
*/
type DeleteNoodleApplicationTabsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNoodleApplicationTabsConflict creates DeleteNoodleApplicationTabsConflict with default headers values
func NewDeleteNoodleApplicationTabsConflict() *DeleteNoodleApplicationTabsConflict {

	return &DeleteNoodleApplicationTabsConflict{}
}

// WithPayload adds the payload to the delete noodle application tabs conflict response
func (o *DeleteNoodleApplicationTabsConflict) WithPayload(payload *models.Error) *DeleteNoodleApplicationTabsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete noodle application tabs conflict response
func (o *DeleteNoodleApplicationTabsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNoodleApplicationTabsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
