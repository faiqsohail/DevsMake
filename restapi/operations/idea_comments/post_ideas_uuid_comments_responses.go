// Code generated by go-swagger; DO NOT EDIT.

package idea_comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"devsmake/models"
)

// PostIdeasUUIDCommentsOKCode is the HTTP code returned for type PostIdeasUUIDCommentsOK
const PostIdeasUUIDCommentsOKCode int = 200

/*PostIdeasUUIDCommentsOK the comment post which was created

swagger:response postIdeasUuidCommentsOK
*/
type PostIdeasUUIDCommentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Comment `json:"body,omitempty"`
}

// NewPostIdeasUUIDCommentsOK creates PostIdeasUUIDCommentsOK with default headers values
func NewPostIdeasUUIDCommentsOK() *PostIdeasUUIDCommentsOK {

	return &PostIdeasUUIDCommentsOK{}
}

// WithPayload adds the payload to the post ideas Uuid comments o k response
func (o *PostIdeasUUIDCommentsOK) WithPayload(payload *models.Comment) *PostIdeasUUIDCommentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ideas Uuid comments o k response
func (o *PostIdeasUUIDCommentsOK) SetPayload(payload *models.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIdeasUUIDCommentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostIdeasUUIDCommentsDefault error

swagger:response postIdeasUuidCommentsDefault
*/
type PostIdeasUUIDCommentsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostIdeasUUIDCommentsDefault creates PostIdeasUUIDCommentsDefault with default headers values
func NewPostIdeasUUIDCommentsDefault(code int) *PostIdeasUUIDCommentsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostIdeasUUIDCommentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post ideas UUID comments default response
func (o *PostIdeasUUIDCommentsDefault) WithStatusCode(code int) *PostIdeasUUIDCommentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post ideas UUID comments default response
func (o *PostIdeasUUIDCommentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post ideas UUID comments default response
func (o *PostIdeasUUIDCommentsDefault) WithPayload(payload *models.Error) *PostIdeasUUIDCommentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ideas UUID comments default response
func (o *PostIdeasUUIDCommentsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIdeasUUIDCommentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}