// Code generated by go-swagger; DO NOT EDIT.

package idea_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPutIdeasUUIDRateParams creates a new PutIdeasUUIDRateParams object
//
// There are no default values defined in the spec.
func NewPutIdeasUUIDRateParams() PutIdeasUUIDRateParams {

	return PutIdeasUUIDRateParams{}
}

// PutIdeasUUIDRateParams contains all the bound params for the put ideas UUID rate operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutIdeasUUIDRate
type PutIdeasUUIDRateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Rating PutIdeasUUIDRateBody
	/*idea post uuid
	  Required: true
	  In: path
	*/
	UUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutIdeasUUIDRateParams() beforehand.
func (o *PutIdeasUUIDRateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PutIdeasUUIDRateBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("rating", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Rating = body
			}
		}
	}

	rUUID, rhkUUID, _ := route.Params.GetOK("uuid")
	if err := o.bindUUID(rUUID, rhkUUID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUUID binds and validates parameter UUID from path.
func (o *PutIdeasUUIDRateParams) bindUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UUID = raw

	return nil
}
