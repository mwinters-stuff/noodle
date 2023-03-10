// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUsersOKCode is the HTTP code returned for type GetNoodleUsersOK
const GetNoodleUsersOKCode int = 200

/*
GetNoodleUsersOK OK

swagger:response getNoodleUsersOK
*/
type GetNoodleUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewGetNoodleUsersOK creates GetNoodleUsersOK with default headers values
func NewGetNoodleUsersOK() *GetNoodleUsersOK {

	return &GetNoodleUsersOK{}
}

// WithPayload adds the payload to the get noodle users o k response
func (o *GetNoodleUsersOK) WithPayload(payload []*models.User) *GetNoodleUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle users o k response
func (o *GetNoodleUsersOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetNoodleUsersUnauthorizedCode is the HTTP code returned for type GetNoodleUsersUnauthorized
const GetNoodleUsersUnauthorizedCode int = 401

/*
GetNoodleUsersUnauthorized unauthorized

swagger:response getNoodleUsersUnauthorized
*/
type GetNoodleUsersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleUsersUnauthorized creates GetNoodleUsersUnauthorized with default headers values
func NewGetNoodleUsersUnauthorized() *GetNoodleUsersUnauthorized {

	return &GetNoodleUsersUnauthorized{}
}

// WithPayload adds the payload to the get noodle users unauthorized response
func (o *GetNoodleUsersUnauthorized) WithPayload(payload *models.Error) *GetNoodleUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle users unauthorized response
func (o *GetNoodleUsersUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoodleUsersConflictCode is the HTTP code returned for type GetNoodleUsersConflict
const GetNoodleUsersConflictCode int = 409

/*
GetNoodleUsersConflict Failed

swagger:response getNoodleUsersConflict
*/
type GetNoodleUsersConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleUsersConflict creates GetNoodleUsersConflict with default headers values
func NewGetNoodleUsersConflict() *GetNoodleUsersConflict {

	return &GetNoodleUsersConflict{}
}

// WithPayload adds the payload to the get noodle users conflict response
func (o *GetNoodleUsersConflict) WithPayload(payload *models.Error) *GetNoodleUsersConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle users conflict response
func (o *GetNoodleUsersConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUsersConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
