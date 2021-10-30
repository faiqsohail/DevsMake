// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// GetProfilesOKCode is the HTTP code returned for type GetProfilesOK
const GetProfilesOKCode int = 200

/*GetProfilesOK success

swagger:response getProfilesOK
*/
type GetProfilesOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetProfilesOK creates GetProfilesOK with default headers values
func NewGetProfilesOK() *GetProfilesOK {

	return &GetProfilesOK{}
}

// WithPayload adds the payload to the get profiles o k response
func (o *GetProfilesOK) WithPayload(payload *models.Profile) *GetProfilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profiles o k response
func (o *GetProfilesOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProfilesDefault error

swagger:response getProfilesDefault
*/
type GetProfilesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProfilesDefault creates GetProfilesDefault with default headers values
func NewGetProfilesDefault(code int) *GetProfilesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProfilesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get profiles default response
func (o *GetProfilesDefault) WithStatusCode(code int) *GetProfilesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get profiles default response
func (o *GetProfilesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get profiles default response
func (o *GetProfilesDefault) WithPayload(payload *models.Error) *GetProfilesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profiles default response
func (o *GetProfilesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfilesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
