// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
)

// FindUsersOKCode is the HTTP code returned for type FindUsersOK
const FindUsersOKCode int = 200

/*FindUsersOK list the User operations

swagger:response findUsersOK
*/
type FindUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewFindUsersOK creates FindUsersOK with default headers values
func NewFindUsersOK() *FindUsersOK {

	return &FindUsersOK{}
}

// WithPayload adds the payload to the find users o k response
func (o *FindUsersOK) WithPayload(payload []*models.User) *FindUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find users o k response
func (o *FindUsersOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*FindUsersDefault generic error response

swagger:response findUsersDefault
*/
type FindUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindUsersDefault creates FindUsersDefault with default headers values
func NewFindUsersDefault(code int) *FindUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &FindUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find users default response
func (o *FindUsersDefault) WithStatusCode(code int) *FindUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find users default response
func (o *FindUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find users default response
func (o *FindUsersDefault) WithPayload(payload *models.Error) *FindUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find users default response
func (o *FindUsersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
