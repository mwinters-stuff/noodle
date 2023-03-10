// Code generated by go-swagger; DO NOT EDIT.

package noodle_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/noodle/server/models"
)

// GetNoodleUploadIconOKCode is the HTTP code returned for type GetNoodleUploadIconOK
const GetNoodleUploadIconOKCode int = 200

/*
GetNoodleUploadIconOK OK

swagger:response getNoodleUploadIconOK
*/
type GetNoodleUploadIconOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetNoodleUploadIconOK creates GetNoodleUploadIconOK with default headers values
func NewGetNoodleUploadIconOK() *GetNoodleUploadIconOK {

	return &GetNoodleUploadIconOK{}
}

// WithPayload adds the payload to the get noodle upload icon o k response
func (o *GetNoodleUploadIconOK) WithPayload(payload []string) *GetNoodleUploadIconOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle upload icon o k response
func (o *GetNoodleUploadIconOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUploadIconOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetNoodleUploadIconUnauthorizedCode is the HTTP code returned for type GetNoodleUploadIconUnauthorized
const GetNoodleUploadIconUnauthorizedCode int = 401

/*
GetNoodleUploadIconUnauthorized unauthorized

swagger:response getNoodleUploadIconUnauthorized
*/
type GetNoodleUploadIconUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleUploadIconUnauthorized creates GetNoodleUploadIconUnauthorized with default headers values
func NewGetNoodleUploadIconUnauthorized() *GetNoodleUploadIconUnauthorized {

	return &GetNoodleUploadIconUnauthorized{}
}

// WithPayload adds the payload to the get noodle upload icon unauthorized response
func (o *GetNoodleUploadIconUnauthorized) WithPayload(payload *models.Error) *GetNoodleUploadIconUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle upload icon unauthorized response
func (o *GetNoodleUploadIconUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUploadIconUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoodleUploadIconConflictCode is the HTTP code returned for type GetNoodleUploadIconConflict
const GetNoodleUploadIconConflictCode int = 409

/*
GetNoodleUploadIconConflict Failed

swagger:response getNoodleUploadIconConflict
*/
type GetNoodleUploadIconConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNoodleUploadIconConflict creates GetNoodleUploadIconConflict with default headers values
func NewGetNoodleUploadIconConflict() *GetNoodleUploadIconConflict {

	return &GetNoodleUploadIconConflict{}
}

// WithPayload adds the payload to the get noodle upload icon conflict response
func (o *GetNoodleUploadIconConflict) WithPayload(payload *models.Error) *GetNoodleUploadIconConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get noodle upload icon conflict response
func (o *GetNoodleUploadIconConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoodleUploadIconConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
