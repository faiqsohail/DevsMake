// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// GetProfileIDOKCode is the HTTP code returned for type GetProfileIDOK
const GetProfileIDOKCode int = 200

/*GetProfileIDOK success

swagger:response getProfileIdOK
*/
type GetProfileIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetProfileIDOK creates GetProfileIDOK with default headers values
func NewGetProfileIDOK() *GetProfileIDOK {

	return &GetProfileIDOK{}
}

// WithPayload adds the payload to the get profile Id o k response
func (o *GetProfileIDOK) WithPayload(payload *models.Profile) *GetProfileIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile Id o k response
func (o *GetProfileIDOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProfileIDDefault error

swagger:response getProfileIdDefault
*/
type GetProfileIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProfileIDDefault creates GetProfileIDDefault with default headers values
func NewGetProfileIDDefault(code int) *GetProfileIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProfileIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get profile ID default response
func (o *GetProfileIDDefault) WithStatusCode(code int) *GetProfileIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get profile ID default response
func (o *GetProfileIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get profile ID default response
func (o *GetProfileIDDefault) WithPayload(payload *models.Error) *GetProfileIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile ID default response
func (o *GetProfileIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}