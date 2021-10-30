// Code generated by go-swagger; DO NOT EDIT.

package idea_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// GetIdeasUUIDOKCode is the HTTP code returned for type GetIdeasUUIDOK
const GetIdeasUUIDOKCode int = 200

/*GetIdeasUUIDOK the idea post fetched

swagger:response getIdeasUuidOK
*/
type GetIdeasUUIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Idea `json:"body,omitempty"`
}

// NewGetIdeasUUIDOK creates GetIdeasUUIDOK with default headers values
func NewGetIdeasUUIDOK() *GetIdeasUUIDOK {

	return &GetIdeasUUIDOK{}
}

// WithPayload adds the payload to the get ideas Uuid o k response
func (o *GetIdeasUUIDOK) WithPayload(payload *models.Idea) *GetIdeasUUIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ideas Uuid o k response
func (o *GetIdeasUUIDOK) SetPayload(payload *models.Idea) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdeasUUIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetIdeasUUIDDefault error

swagger:response getIdeasUuidDefault
*/
type GetIdeasUUIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIdeasUUIDDefault creates GetIdeasUUIDDefault with default headers values
func NewGetIdeasUUIDDefault(code int) *GetIdeasUUIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetIdeasUUIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ideas UUID default response
func (o *GetIdeasUUIDDefault) WithStatusCode(code int) *GetIdeasUUIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ideas UUID default response
func (o *GetIdeasUUIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ideas UUID default response
func (o *GetIdeasUUIDDefault) WithPayload(payload *models.Error) *GetIdeasUUIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ideas UUID default response
func (o *GetIdeasUUIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdeasUUIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
