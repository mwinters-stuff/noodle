// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// PostNoodleTabsOKCode is the HTTP code returned for type PostNoodleTabsOK
const PostNoodleTabsOKCode int = 200

/*
PostNoodleTabsOK OK

swagger:response postNoodleTabsOK
*/
type PostNoodleTabsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Tab `json:"body,omitempty"`
}

// NewPostNoodleTabsOK creates PostNoodleTabsOK with default headers values
func NewPostNoodleTabsOK() *PostNoodleTabsOK {

	return &PostNoodleTabsOK{}
}

// WithPayload adds the payload to the post noodle tabs o k response
func (o *PostNoodleTabsOK) WithPayload(payload *models.Tab) *PostNoodleTabsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post noodle tabs o k response
func (o *PostNoodleTabsOK) SetPayload(payload *models.Tab) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostNoodleTabsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostNoodleTabsUnauthorizedCode is the HTTP code returned for type PostNoodleTabsUnauthorized
const PostNoodleTabsUnauthorizedCode int = 401

/*
PostNoodleTabsUnauthorized unauthorized

swagger:response postNoodleTabsUnauthorized
*/
type PostNoodleTabsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostNoodleTabsUnauthorized creates PostNoodleTabsUnauthorized with default headers values
func NewPostNoodleTabsUnauthorized() *PostNoodleTabsUnauthorized {

	return &PostNoodleTabsUnauthorized{}
}

// WithPayload adds the payload to the post noodle tabs unauthorized response
func (o *PostNoodleTabsUnauthorized) WithPayload(payload *models.Error) *PostNoodleTabsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post noodle tabs unauthorized response
func (o *PostNoodleTabsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostNoodleTabsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostNoodleTabsConflictCode is the HTTP code returned for type PostNoodleTabsConflict
const PostNoodleTabsConflictCode int = 409

/*
PostNoodleTabsConflict Failed

swagger:response postNoodleTabsConflict
*/
type PostNoodleTabsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostNoodleTabsConflict creates PostNoodleTabsConflict with default headers values
func NewPostNoodleTabsConflict() *PostNoodleTabsConflict {

	return &PostNoodleTabsConflict{}
}

// WithPayload adds the payload to the post noodle tabs conflict response
func (o *PostNoodleTabsConflict) WithPayload(payload *models.Error) *PostNoodleTabsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post noodle tabs conflict response
func (o *PostNoodleTabsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostNoodleTabsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
