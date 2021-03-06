// Code generated by go-swagger; DO NOT EDIT.

package idea_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// PostIdeasOKCode is the HTTP code returned for type PostIdeasOK
const PostIdeasOKCode int = 200

/*PostIdeasOK the idea post which was created

swagger:response postIdeasOK
*/
type PostIdeasOK struct {

	/*
	  In: Body
	*/
	Payload *models.Idea `json:"body,omitempty"`
}

// NewPostIdeasOK creates PostIdeasOK with default headers values
func NewPostIdeasOK() *PostIdeasOK {

	return &PostIdeasOK{}
}

// WithPayload adds the payload to the post ideas o k response
func (o *PostIdeasOK) WithPayload(payload *models.Idea) *PostIdeasOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ideas o k response
func (o *PostIdeasOK) SetPayload(payload *models.Idea) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIdeasOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostIdeasDefault error

swagger:response postIdeasDefault
*/
type PostIdeasDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostIdeasDefault creates PostIdeasDefault with default headers values
func NewPostIdeasDefault(code int) *PostIdeasDefault {
	if code <= 0 {
		code = 500
	}

	return &PostIdeasDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post ideas default response
func (o *PostIdeasDefault) WithStatusCode(code int) *PostIdeasDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post ideas default response
func (o *PostIdeasDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post ideas default response
func (o *PostIdeasDefault) WithPayload(payload *models.Error) *PostIdeasDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ideas default response
func (o *PostIdeasDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIdeasDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
