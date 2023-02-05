// Code generated by go-swagger; DO NOT EDIT.

package noodle_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetAuthSessionOKCode is the HTTP code returned for type GetAuthSessionOK
const GetAuthSessionOKCode int = 200

/*
GetAuthSessionOK OK

swagger:response getAuthSessionOK
*/
type GetAuthSessionOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserSession `json:"body,omitempty"`
}

// NewGetAuthSessionOK creates GetAuthSessionOK with default headers values
func NewGetAuthSessionOK() *GetAuthSessionOK {

	return &GetAuthSessionOK{}
}

// WithPayload adds the payload to the get auth session o k response
func (o *GetAuthSessionOK) WithPayload(payload *models.UserSession) *GetAuthSessionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth session o k response
func (o *GetAuthSessionOK) SetPayload(payload *models.UserSession) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthSessionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAuthSessionUnauthorizedCode is the HTTP code returned for type GetAuthSessionUnauthorized
const GetAuthSessionUnauthorizedCode int = 401

/*
GetAuthSessionUnauthorized unauthorized

swagger:response getAuthSessionUnauthorized
*/
type GetAuthSessionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAuthSessionUnauthorized creates GetAuthSessionUnauthorized with default headers values
func NewGetAuthSessionUnauthorized() *GetAuthSessionUnauthorized {

	return &GetAuthSessionUnauthorized{}
}

// WithPayload adds the payload to the get auth session unauthorized response
func (o *GetAuthSessionUnauthorized) WithPayload(payload *models.Error) *GetAuthSessionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth session unauthorized response
func (o *GetAuthSessionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthSessionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAuthSessionConflictCode is the HTTP code returned for type GetAuthSessionConflict
const GetAuthSessionConflictCode int = 409

/*
GetAuthSessionConflict Failed

swagger:response getAuthSessionConflict
*/
type GetAuthSessionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAuthSessionConflict creates GetAuthSessionConflict with default headers values
func NewGetAuthSessionConflict() *GetAuthSessionConflict {

	return &GetAuthSessionConflict{}
}

// WithPayload adds the payload to the get auth session conflict response
func (o *GetAuthSessionConflict) WithPayload(payload *models.Error) *GetAuthSessionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth session conflict response
func (o *GetAuthSessionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthSessionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
