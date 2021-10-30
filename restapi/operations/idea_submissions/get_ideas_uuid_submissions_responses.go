// Code generated by go-swagger; DO NOT EDIT.

package idea_submissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// GetIdeasUUIDSubmissionsOKCode is the HTTP code returned for type GetIdeasUUIDSubmissionsOK
const GetIdeasUUIDSubmissionsOKCode int = 200

/*GetIdeasUUIDSubmissionsOK the submissions for the idea post fetched

swagger:response getIdeasUuidSubmissionsOK
*/
type GetIdeasUUIDSubmissionsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Submission `json:"body,omitempty"`
}

// NewGetIdeasUUIDSubmissionsOK creates GetIdeasUUIDSubmissionsOK with default headers values
func NewGetIdeasUUIDSubmissionsOK() *GetIdeasUUIDSubmissionsOK {

	return &GetIdeasUUIDSubmissionsOK{}
}

// WithPayload adds the payload to the get ideas Uuid submissions o k response
func (o *GetIdeasUUIDSubmissionsOK) WithPayload(payload []*models.Submission) *GetIdeasUUIDSubmissionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ideas Uuid submissions o k response
func (o *GetIdeasUUIDSubmissionsOK) SetPayload(payload []*models.Submission) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdeasUUIDSubmissionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Submission, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetIdeasUUIDSubmissionsDefault error

swagger:response getIdeasUuidSubmissionsDefault
*/
type GetIdeasUUIDSubmissionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIdeasUUIDSubmissionsDefault creates GetIdeasUUIDSubmissionsDefault with default headers values
func NewGetIdeasUUIDSubmissionsDefault(code int) *GetIdeasUUIDSubmissionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetIdeasUUIDSubmissionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ideas UUID submissions default response
func (o *GetIdeasUUIDSubmissionsDefault) WithStatusCode(code int) *GetIdeasUUIDSubmissionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ideas UUID submissions default response
func (o *GetIdeasUUIDSubmissionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ideas UUID submissions default response
func (o *GetIdeasUUIDSubmissionsDefault) WithPayload(payload *models.Error) *GetIdeasUUIDSubmissionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ideas UUID submissions default response
func (o *GetIdeasUUIDSubmissionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdeasUUIDSubmissionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
