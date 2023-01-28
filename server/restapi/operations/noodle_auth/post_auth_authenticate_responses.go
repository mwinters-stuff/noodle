// Code generated by go-swagger; DO NOT EDIT.

package noodle_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// PostAuthAuthenticateOKCode is the HTTP code returned for type PostAuthAuthenticateOK
const PostAuthAuthenticateOKCode int = 200

/*
PostAuthAuthenticateOK OK

swagger:response postAuthAuthenticateOK
*/
type PostAuthAuthenticateOK struct {

	/*
	  In: Body
	*/
	Payload *PostAuthAuthenticateOKBody `json:"body,omitempty"`
}

// NewPostAuthAuthenticateOK creates PostAuthAuthenticateOK with default headers values
func NewPostAuthAuthenticateOK() *PostAuthAuthenticateOK {

	return &PostAuthAuthenticateOK{}
}

// WithPayload adds the payload to the post auth authenticate o k response
func (o *PostAuthAuthenticateOK) WithPayload(payload *PostAuthAuthenticateOKBody) *PostAuthAuthenticateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth authenticate o k response
func (o *PostAuthAuthenticateOK) SetPayload(payload *PostAuthAuthenticateOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthAuthenticateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthAuthenticateUnauthorizedCode is the HTTP code returned for type PostAuthAuthenticateUnauthorized
const PostAuthAuthenticateUnauthorizedCode int = 401

/*
PostAuthAuthenticateUnauthorized unauthorized

swagger:response postAuthAuthenticateUnauthorized
*/
type PostAuthAuthenticateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthAuthenticateUnauthorized creates PostAuthAuthenticateUnauthorized with default headers values
func NewPostAuthAuthenticateUnauthorized() *PostAuthAuthenticateUnauthorized {

	return &PostAuthAuthenticateUnauthorized{}
}

// WithPayload adds the payload to the post auth authenticate unauthorized response
func (o *PostAuthAuthenticateUnauthorized) WithPayload(payload *models.Error) *PostAuthAuthenticateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth authenticate unauthorized response
func (o *PostAuthAuthenticateUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthAuthenticateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthAuthenticateConflictCode is the HTTP code returned for type PostAuthAuthenticateConflict
const PostAuthAuthenticateConflictCode int = 409

/*
PostAuthAuthenticateConflict Failed

swagger:response postAuthAuthenticateConflict
*/
type PostAuthAuthenticateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthAuthenticateConflict creates PostAuthAuthenticateConflict with default headers values
func NewPostAuthAuthenticateConflict() *PostAuthAuthenticateConflict {

	return &PostAuthAuthenticateConflict{}
}

// WithPayload adds the payload to the post auth authenticate conflict response
func (o *PostAuthAuthenticateConflict) WithPayload(payload *models.Error) *PostAuthAuthenticateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth authenticate conflict response
func (o *PostAuthAuthenticateConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthAuthenticateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
