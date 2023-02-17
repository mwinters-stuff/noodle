// Code generated by go-swagger; DO NOT EDIT.

package noodle_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetAuthLogoutOKCode is the HTTP code returned for type GetAuthLogoutOK
const GetAuthLogoutOKCode int = 200

/*
GetAuthLogoutOK OK

swagger:response getAuthLogoutOK
*/
type GetAuthLogoutOK struct {
}

// NewGetAuthLogoutOK creates GetAuthLogoutOK with default headers values
func NewGetAuthLogoutOK() *GetAuthLogoutOK {

	return &GetAuthLogoutOK{}
}

// WriteResponse to the client
func (o *GetAuthLogoutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetAuthLogoutUnauthorizedCode is the HTTP code returned for type GetAuthLogoutUnauthorized
const GetAuthLogoutUnauthorizedCode int = 401

/*
GetAuthLogoutUnauthorized unauthorized

swagger:response getAuthLogoutUnauthorized
*/
type GetAuthLogoutUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAuthLogoutUnauthorized creates GetAuthLogoutUnauthorized with default headers values
func NewGetAuthLogoutUnauthorized() *GetAuthLogoutUnauthorized {

	return &GetAuthLogoutUnauthorized{}
}

// WithPayload adds the payload to the get auth logout unauthorized response
func (o *GetAuthLogoutUnauthorized) WithPayload(payload *models.Error) *GetAuthLogoutUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth logout unauthorized response
func (o *GetAuthLogoutUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthLogoutUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}