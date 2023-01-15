// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleLdapReloadOKCode is the HTTP code returned for type GetNoodleLdapReloadOK
const GetNoodleLdapReloadOKCode int = 200

/*
GetNoodleLdapReloadOK Success

swagger:response getNoodleLdapReloadOK
*/
type GetNoodleLdapReloadOK struct {
}

// NewGetNoodleLdapReloadOK creates GetNoodleLdapReloadOK with default headers values
func NewGetNoodleLdapReloadOK() *GetNoodleLdapReloadOK {

	return &GetNoodleLdapReloadOK{}
}

// WriteResponse to the client
func (o *GetNoodleLdapReloadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetNoodleLdapReloadUnauthorizedCode is the HTTP code returned for type GetNoodleLdapReloadUnauthorized
const GetNoodleLdapReloadUnauthorizedCode int = 401

/*
GetNoodleLdapReloadUnauthorized unauthorized

swagger:response getNoodleLdapReloadUnauthorized
*/
type GetNoodleLdapReloadUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleLdapReloadUnauthorized creates GetNoodleLdapReloadUnauthorized with default headers values
func NewGetNoodleLdapReloadUnauthorized() *GetNoodleLdapReloadUnauthorized {

	return &GetNoodleLdapReloadUnauthorized{}
}

// WithPayload adds the payload to the get noodle ldap reload unauthorized response
func (o *GetNoodleLdapReloadUnauthorized) WithPayload(payload *models.Error) *GetNoodleLdapReloadUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle ldap reload unauthorized response
func (o *GetNoodleLdapReloadUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleLdapReloadUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoodleLdapReloadConflictCode is the HTTP code returned for type GetNoodleLdapReloadConflict
const GetNoodleLdapReloadConflictCode int = 409

/*
GetNoodleLdapReloadConflict Failed

swagger:response getNoodleLdapReloadConflict
*/
type GetNoodleLdapReloadConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleLdapReloadConflict creates GetNoodleLdapReloadConflict with default headers values
func NewGetNoodleLdapReloadConflict() *GetNoodleLdapReloadConflict {

	return &GetNoodleLdapReloadConflict{}
}

// WithPayload adds the payload to the get noodle ldap reload conflict response
func (o *GetNoodleLdapReloadConflict) WithPayload(payload *models.Error) *GetNoodleLdapReloadConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle ldap reload conflict response
func (o *GetNoodleLdapReloadConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleLdapReloadConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
