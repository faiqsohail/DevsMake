// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// GetProfileOKCode is the HTTP code returned for type GetProfileOK
const GetProfileOKCode int = 200

/*GetProfileOK success

swagger:response getProfileOK
*/
type GetProfileOK struct {

	/*
	  In: Body
	*/
	Payload *GetProfileOKBody `json:"body,omitempty"`
}

// NewGetProfileOK creates GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {

	return &GetProfileOK{}
}

// WithPayload adds the payload to the get profile o k response
func (o *GetProfileOK) WithPayload(payload *GetProfileOKBody) *GetProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile o k response
func (o *GetProfileOK) SetPayload(payload *GetProfileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProfileDefault error

swagger:response getProfileDefault
*/
type GetProfileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProfileDefault creates GetProfileDefault with default headers values
func NewGetProfileDefault(code int) *GetProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get profile default response
func (o *GetProfileDefault) WithStatusCode(code int) *GetProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get profile default response
func (o *GetProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get profile default response
func (o *GetProfileDefault) WithPayload(payload *models.Error) *GetProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile default response
func (o *GetProfileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
